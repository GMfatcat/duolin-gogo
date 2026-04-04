package cards

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Card struct {
	ID               string   `json:"id"`
	SourcePath       string   `json:"source_path,omitempty"`
	SourceModifiedAt string   `json:"source_modified_at,omitempty"`
	SourceHash       string   `json:"source_hash,omitempty"`
	Enabled          bool     `json:"enabled"`
	Title            string   `json:"title"`
	QuestionType     string   `json:"type"`
	BodyFormat       string   `json:"body_format"`
	Tags             []string `json:"tags"`
	Difficulty       int      `json:"difficulty"`
	QuestionText     string   `json:"question"`
	Choices          []string `json:"choices,omitempty"`
	AnswerValue      any      `json:"answer"`
	Clickbait        string   `json:"clickbait,omitempty"`
	ReviewHint       string   `json:"review_hint,omitempty"`
	BodyMarkdownZH   string   `json:"body_markdown_zh"`
	BodyMarkdownEN   string   `json:"body_markdown_en"`
	BodyPlaintextZH  string   `json:"body_plaintext_zh,omitempty"`
	BodyPlaintextEN  string   `json:"body_plaintext_en,omitempty"`
}

type ImportError struct {
	SourcePath string `json:"source_path"`
	Code       string `json:"code"`
	Field      string `json:"field,omitempty"`
	Message    string `json:"message"`
}

type ImportResult struct {
	Cards  []Card
	Errors []ImportError
}

type CacheFile struct {
	Version     int    `json:"version"`
	GeneratedAt string `json:"generated_at"`
	Cards       []Card `json:"cards"`
}

type ImportErrorsFile struct {
	Version     int           `json:"version"`
	GeneratedAt string        `json:"generated_at"`
	Errors      []ImportError `json:"errors"`
}

type frontmatter struct {
	ID         string   `yaml:"id"`
	Title      string   `yaml:"title"`
	Type       string   `yaml:"type"`
	BodyFormat string   `yaml:"body_format"`
	Tags       []string `yaml:"tags"`
	Difficulty int      `yaml:"difficulty"`
	Question   string   `yaml:"question"`
	Choices    []string `yaml:"choices"`
	Answer     any      `yaml:"answer"`
	Clickbait  string   `yaml:"clickbait"`
	ReviewHint string   `yaml:"review_hint"`
	Enabled    *bool    `yaml:"enabled"`
}

func ScanDirectories(paths []string) (ImportResult, error) {
	result := ImportResult{}
	seenIDs := map[string]string{}

	for _, root := range paths {
		if root == "" {
			continue
		}

		if err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || !strings.EqualFold(filepath.Ext(path), ".md") {
				return nil
			}

			card, importErr := parseFile(path)
			if importErr != nil {
				result.Errors = append(result.Errors, *importErr)
				return nil
			}

			if firstPath, exists := seenIDs[card.ID]; exists {
				result.Errors = append(result.Errors, ImportError{
					SourcePath: path,
					Code:       "duplicate_id",
					Field:      "id",
					Message:    fmt.Sprintf("Duplicate card id %q already used by %s.", card.ID, firstPath),
				})
				return nil
			}

			seenIDs[card.ID] = path
			result.Cards = append(result.Cards, card)
			return nil
		}); err != nil {
			return ImportResult{}, err
		}
	}

	slices.SortFunc(result.Cards, func(a, b Card) int {
		return strings.Compare(a.ID, b.ID)
	})

	return result, nil
}

func WriteCache(path string, cards []Card) error {
	cache := CacheFile{
		Version:     1,
		GeneratedAt: time.Now().Format(time.RFC3339),
		Cards:       cards,
	}

	return writeJSON(path, cache)
}

func WriteImportErrors(path string, errs []ImportError) error {
	if errs == nil {
		errs = []ImportError{}
	}

	file := ImportErrorsFile{
		Version:     1,
		GeneratedAt: time.Now().Format(time.RFC3339),
		Errors:      errs,
	}

	return writeJSON(path, file)
}

func RefreshKnowledge(knowledgeDir, dataDir string) (ImportResult, error) {
	result, err := ScanDirectories([]string{knowledgeDir})
	if err != nil {
		return ImportResult{}, err
	}

	if err := WriteCache(filepath.Join(dataDir, "cards-cache.json"), result.Cards); err != nil {
		return ImportResult{}, err
	}

	if err := WriteImportErrors(filepath.Join(dataDir, "import-errors.json"), result.Errors); err != nil {
		return ImportResult{}, err
	}

	return result, nil
}

func parseFile(path string) (Card, *ImportError) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Card{}, &ImportError{
			SourcePath: path,
			Code:       "read_failed",
			Message:    err.Error(),
		}
	}

	fmContent, body, err := splitFrontmatter(string(raw))
	if err != nil {
		return Card{}, &ImportError{
			SourcePath: path,
			Code:       "frontmatter_parse_failed",
			Field:      "frontmatter",
			Message:    err.Error(),
		}
	}

	var fm frontmatter
	if err := yaml.Unmarshal([]byte(fmContent), &fm); err != nil {
		return Card{}, &ImportError{
			SourcePath: path,
			Code:       "frontmatter_parse_failed",
			Field:      "frontmatter",
			Message:    err.Error(),
		}
	}

	card, importErr := buildCard(path, fm, body)
	if importErr != nil {
		return Card{}, importErr
	}

	return card, nil
}

func buildCard(path string, fm frontmatter, body string) (Card, *ImportError) {
	if strings.TrimSpace(fm.ID) == "" {
		return Card{}, validationError(path, "missing_required_field", "id", "Required field 'id' is missing.")
	}
	if strings.TrimSpace(fm.Title) == "" {
		return Card{}, validationError(path, "missing_required_field", "title", "Required field 'title' is missing.")
	}
	if strings.TrimSpace(fm.Type) == "" {
		return Card{}, validationError(path, "missing_required_field", "type", "Required field 'type' is missing.")
	}
	if strings.TrimSpace(fm.Question) == "" {
		return Card{}, validationError(path, "missing_required_field", "question", "Required field 'question' is missing.")
	}

	questionType := strings.ToLower(strings.TrimSpace(fm.Type))
	if questionType != "single-choice" && questionType != "true-false" {
		return Card{}, validationError(path, "unsupported_type", "type", fmt.Sprintf("Unsupported type %q.", fm.Type))
	}

	if fm.BodyFormat == "" {
		fm.BodyFormat = "bilingual-section"
	}
	if fm.Difficulty == 0 {
		fm.Difficulty = 2
	}

	enabled := true
	if fm.Enabled != nil {
		enabled = *fm.Enabled
	}

	zh, en, err := extractLanguageSections(body)
	if err != nil {
		return Card{}, validationError(path, "missing_language_section", "body", err.Error())
	}

	answer, importErr := normalizeAnswer(path, questionType, fm.Answer, fm.Choices)
	if importErr != nil {
		return Card{}, importErr
	}

	info, statErr := os.Stat(path)
	modifiedAt := ""
	if statErr == nil {
		modifiedAt = info.ModTime().Format(time.RFC3339)
	}

	return Card{
		ID:               strings.TrimSpace(fm.ID),
		SourcePath:       path,
		SourceModifiedAt: modifiedAt,
		Enabled:          enabled,
		Title:            strings.TrimSpace(fm.Title),
		QuestionType:     questionType,
		BodyFormat:       fm.BodyFormat,
		Tags:             normalizeTags(fm.Tags),
		Difficulty:       fm.Difficulty,
		QuestionText:     strings.TrimSpace(fm.Question),
		Choices:          fm.Choices,
		AnswerValue:      answer,
		Clickbait:        strings.TrimSpace(fm.Clickbait),
		ReviewHint:       strings.TrimSpace(fm.ReviewHint),
		BodyMarkdownZH:   zh,
		BodyMarkdownEN:   en,
		BodyPlaintextZH:  toPlaintext(zh),
		BodyPlaintextEN:  toPlaintext(en),
	}, nil
}

func normalizeAnswer(path, questionType string, raw any, choices []string) (any, *ImportError) {
	switch questionType {
	case "single-choice":
		if len(choices) < 2 {
			return nil, validationError(path, "missing_choices", "choices", "Single-choice cards require at least 2 choices.")
		}

		answer, ok := raw.(int)
		if !ok {
			return nil, validationError(path, "invalid_answer_type", "answer", "Single-choice answer must be an integer index.")
		}

		if answer < 0 || answer >= len(choices) {
			return nil, validationError(path, "answer_out_of_range", "answer", fmt.Sprintf("Answer index %d exceeds available choices.", answer))
		}

		return answer, nil
	case "true-false":
		answer, ok := raw.(bool)
		if !ok {
			return nil, validationError(path, "invalid_answer_type", "answer", "True-false answer must be boolean.")
		}

		return answer, nil
	default:
		return nil, validationError(path, "unsupported_type", "type", fmt.Sprintf("Unsupported type %q.", questionType))
	}
}

func splitFrontmatter(raw string) (string, string, error) {
	raw = strings.ReplaceAll(raw, "\r\n", "\n")
	if !strings.HasPrefix(raw, "---\n") {
		return "", "", errors.New("missing opening frontmatter delimiter")
	}

	parts := strings.SplitN(raw, "\n---\n", 2)
	if len(parts) != 2 {
		return "", "", errors.New("missing closing frontmatter delimiter")
	}

	return strings.TrimPrefix(parts[0], "---\n"), parts[1], nil
}

func extractLanguageSections(body string) (string, string, error) {
	body = strings.ReplaceAll(body, "\r\n", "\n")
	re := regexp.MustCompile(`(?s)## zh-TW\s*(.*?)\s*## en\s*(.*)$`)
	matches := re.FindStringSubmatch(body)
	if len(matches) != 3 {
		return "", "", errors.New("Body must contain both ## zh-TW and ## en sections.")
	}

	zh := strings.TrimSpace(matches[1])
	en := strings.TrimSpace(matches[2])
	if zh == "" || en == "" {
		return "", "", errors.New("Body must contain both ## zh-TW and ## en sections.")
	}

	return zh, en, nil
}

func validationError(path, code, field, message string) *ImportError {
	return &ImportError{
		SourcePath: path,
		Code:       code,
		Field:      field,
		Message:    message,
	}
}

func normalizeTags(tags []string) []string {
	out := make([]string, 0, len(tags))
	for _, tag := range tags {
		tag = strings.ToLower(strings.TrimSpace(tag))
		if tag != "" {
			out = append(out, tag)
		}
	}
	return out
}

func toPlaintext(markdown string) string {
	replacer := strings.NewReplacer("`", "", "*", "", "_", "")
	plain := replacer.Replace(markdown)
	plain = strings.ReplaceAll(plain, "\n", " ")
	return strings.Join(strings.Fields(plain), " ")
}

func writeJSON(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, append(bytes, '\n'), 0o644)
}
