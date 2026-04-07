package cards

import (
	"crypto/sha256"
	"encoding/gob"
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
	TitleZH          string   `json:"title_zh"`
	TitleEN          string   `json:"title_en"`
	QuestionType     string   `json:"type"`
	BodyFormat       string   `json:"body_format"`
	Tags             []string `json:"tags"`
	Difficulty       int      `json:"difficulty"`
	QuestionText     string   `json:"question"`
	QuestionTextZH   string   `json:"question_zh"`
	QuestionTextEN   string   `json:"question_en"`
	Choices          []string `json:"choices,omitempty"`
	ChoicesZH        []string `json:"choices_zh,omitempty"`
	ChoicesEN        []string `json:"choices_en,omitempty"`
	AnswerValue      any      `json:"answer"`
	Clickbait        string   `json:"clickbait,omitempty"`
	ClickbaitZH      string   `json:"clickbait_zh,omitempty"`
	ClickbaitEN      string   `json:"clickbait_en,omitempty"`
	ReviewHint       string   `json:"review_hint,omitempty"`
	ReviewHintZH     string   `json:"review_hint_zh,omitempty"`
	ReviewHintEN     string   `json:"review_hint_en,omitempty"`
	ConfusionWith    []string `json:"confusion_with,omitempty"`
	MetaphorSeed     []string `json:"metaphor_seed,omitempty"`
	HookStyleTags    []string `json:"hook_style_tags,omitempty"`
	BodyMarkdownZH   string   `json:"body_markdown_zh"`
	BodyMarkdownEN   string   `json:"body_markdown_en"`
	BodyPlaintextZH  string   `json:"body_plaintext_zh,omitempty"`
	BodyPlaintextEN  string   `json:"body_plaintext_en,omitempty"`
}

type ImportError struct {
	SourcePath string `json:"source_path"`
	Severity   string `json:"severity,omitempty"`
	Code       string `json:"code"`
	Field      string `json:"field,omitempty"`
	Message    string `json:"message"`
}

type ImportResult struct {
	Cards  []Card
	Errors []ImportError
}

type PreviewResult struct {
	Card   *Card         `json:"card,omitempty"`
	Errors []ImportError `json:"errors"`
}

type CacheFile struct {
	Version             int            `json:"version"`
	GeneratedAt         string         `json:"generated_at"`
	KnowledgeFingerprint string        `json:"knowledge_fingerprint,omitempty"`
	SourceFiles         []CacheSource  `json:"source_files,omitempty"`
	Cards               []Card         `json:"cards"`
}

type CacheSource struct {
	Path            string `json:"path"`
	ModifiedUnixNano int64 `json:"modified_unix_nano"`
	Size            int64  `json:"size"`
}

type ImportErrorsFile struct {
	Version     int           `json:"version"`
	GeneratedAt string        `json:"generated_at"`
	Errors      []ImportError `json:"errors"`
}

type frontmatter struct {
	ID            string   `yaml:"id"`
	Title         string   `yaml:"title"`
	TitleZH       string   `yaml:"title_zh"`
	TitleEN       string   `yaml:"title_en"`
	Type          string   `yaml:"type"`
	BodyFormat    string   `yaml:"body_format"`
	Tags          []string `yaml:"tags"`
	Difficulty    int      `yaml:"difficulty"`
	Question      string   `yaml:"question"`
	QuestionZH    string   `yaml:"question_zh"`
	QuestionEN    string   `yaml:"question_en"`
	Choices       []string `yaml:"choices"`
	ChoicesZH     []string `yaml:"choices_zh"`
	ChoicesEN     []string `yaml:"choices_en"`
	Answer        any      `yaml:"answer"`
	Clickbait     string   `yaml:"clickbait"`
	ClickbaitZH   string   `yaml:"clickbait_zh"`
	ClickbaitEN   string   `yaml:"clickbait_en"`
	ReviewHint    string   `yaml:"review_hint"`
	ReviewHintZH  string   `yaml:"review_hint_zh"`
	ReviewHintEN  string   `yaml:"review_hint_en"`
	ConfusionWith []string `yaml:"confusion_with"`
	MetaphorSeed  []string `yaml:"metaphor_seed"`
	HookStyleTags []string `yaml:"hook_style_tags"`
	Enabled       *bool    `yaml:"enabled"`
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

			card, diagnostics, importErr := parseFile(path)
			result.Errors = append(result.Errors, diagnostics...)
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

func WriteCache(path string, cache CacheFile) error {
	cache.Version = 2
	if cache.GeneratedAt == "" {
		cache.GeneratedAt = time.Now().Format(time.RFC3339)
	}

	switch strings.ToLower(filepath.Ext(path)) {
	case ".gob":
		return writeGOB(path, cache)
	default:
		return writeJSON(path, cache)
	}
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

	sources, err := SnapshotKnowledgeFiles([]string{knowledgeDir})
	if err != nil {
		return ImportResult{}, err
	}

	cache := CacheFile{
		KnowledgeFingerprint: fingerprintKnowledgeFiles(sources),
		SourceFiles:          sources,
		Cards:                result.Cards,
	}

	if err := WriteCache(filepath.Join(dataDir, "cards-cache.gob"), cache); err != nil {
		return ImportResult{}, err
	}

	if err := WriteImportErrors(filepath.Join(dataDir, "import-errors.json"), result.Errors); err != nil {
		return ImportResult{}, err
	}

	return result, nil
}

func EnsureKnowledgeCache(knowledgeDir, dataDir string) (CacheFile, bool, error) {
	sources, err := SnapshotKnowledgeFiles([]string{knowledgeDir})
	if err != nil {
		return CacheFile{}, false, err
	}

	cachePath := filepath.Join(dataDir, "cards-cache.gob")
	fingerprint := fingerprintKnowledgeFiles(sources)
	if cache, err := LoadCache(cachePath); err == nil {
		if cache.Version >= 2 && cache.KnowledgeFingerprint == fingerprint {
			return cache, false, nil
		}
	}

	result, err := RefreshKnowledge(knowledgeDir, dataDir)
	if err != nil {
		return CacheFile{}, false, err
	}

	cache, err := LoadCache(cachePath)
	if err != nil {
		return CacheFile{}, false, err
	}
	cache.Cards = result.Cards
	return cache, true, nil
}

func ListMarkdownFiles(paths []string) ([]string, error) {
	files := []string{}

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

			files = append(files, path)
			return nil
		}); err != nil {
			return nil, err
		}
	}

	slices.Sort(files)
	return files, nil
}

func PreviewFile(path string) (PreviewResult, error) {
	card, diagnostics, importErr := parseFile(path)
	result := PreviewResult{
		Errors: diagnostics,
	}

	if importErr != nil {
		result.Errors = append(result.Errors, *importErr)
		return result, nil
	}

	result.Card = &card
	return result, nil
}

func PreviewDraft(sourcePath string, raw string) (PreviewResult, error) {
	fmContent, body, err := splitFrontmatter(raw)
	if err != nil {
		return PreviewResult{
			Errors: []ImportError{{
				SourcePath: sourcePath,
				Severity:   "error",
				Code:       "frontmatter_parse_failed",
				Field:      "frontmatter",
				Message:    err.Error(),
			}},
		}, nil
	}

	var fm frontmatter
	if err := yaml.Unmarshal([]byte(fmContent), &fm); err != nil {
		return PreviewResult{
			Errors: []ImportError{{
				SourcePath: sourcePath,
				Severity:   "error",
				Code:       "frontmatter_parse_failed",
				Field:      "frontmatter",
				Message:    err.Error(),
			}},
		}, nil
	}

	card, diagnostics, importErr := buildCard(sourcePath, fm, body)
	result := PreviewResult{
		Errors: diagnostics,
	}
	if importErr != nil {
		result.Errors = append(result.Errors, *importErr)
		return result, nil
	}

	result.Card = &card
	return result, nil
}

func parseFile(path string) (Card, []ImportError, *ImportError) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Card{}, nil, &ImportError{
			SourcePath: path,
			Severity:   "error",
			Code:       "read_failed",
			Message:    err.Error(),
		}
	}

	fmContent, body, err := splitFrontmatter(string(raw))
	if err != nil {
		return Card{}, nil, &ImportError{
			SourcePath: path,
			Severity:   "error",
			Code:       "frontmatter_parse_failed",
			Field:      "frontmatter",
			Message:    err.Error(),
		}
	}

	var fm frontmatter
	if err := yaml.Unmarshal([]byte(fmContent), &fm); err != nil {
		return Card{}, nil, &ImportError{
			SourcePath: path,
			Severity:   "error",
			Code:       "frontmatter_parse_failed",
			Field:      "frontmatter",
			Message:    err.Error(),
		}
	}

	card, diagnostics, importErr := buildCard(path, fm, body)
	if importErr != nil {
		return Card{}, diagnostics, importErr
	}

	return card, diagnostics, nil
}

func buildCard(path string, fm frontmatter, body string) (Card, []ImportError, *ImportError) {
	diagnostics := []ImportError{}
	if strings.TrimSpace(fm.ID) == "" {
		return Card{}, diagnostics, validationError(path, "missing_required_field", "id", "Required field 'id' is missing.")
	}
	if localizeText(fm.TitleZH, fm.Title, fm.TitleEN) == "" {
		return Card{}, diagnostics, validationError(path, "missing_required_field", "title", "Required field 'title' is missing.")
	}
	if strings.TrimSpace(fm.Type) == "" {
		return Card{}, diagnostics, validationError(path, "missing_required_field", "type", "Required field 'type' is missing.")
	}
	if localizeText(fm.QuestionZH, fm.Question, fm.QuestionEN) == "" {
		return Card{}, diagnostics, validationError(path, "missing_required_field", "question", "Required field 'question' is missing.")
	}

	questionType := strings.ToLower(strings.TrimSpace(fm.Type))
	if questionType != "single-choice" && questionType != "true-false" {
		return Card{}, diagnostics, validationError(path, "unsupported_type", "type", fmt.Sprintf("Unsupported type %q.", fm.Type))
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
		return Card{}, diagnostics, validationError(path, "missing_language_section", "body", err.Error())
	}

	diagnostics = append(diagnostics, localizedFieldWarnings(path, fm)...)

	choicesEN := localizeChoices(fm.ChoicesEN, fm.Choices, fm.ChoicesZH)
	choicesZH := localizeChoices(fm.ChoicesZH, fm.Choices, fm.ChoicesEN)
	diagnostics = append(diagnostics, localizedChoiceWarnings(path, fm, choicesZH, choicesEN)...)

	answer, importErr := normalizeAnswer(path, questionType, fm.Answer, choicesEN)
	if importErr != nil {
		return Card{}, diagnostics, importErr
	}

	info, statErr := os.Stat(path)
	modifiedAt := ""
	if statErr == nil {
		modifiedAt = info.ModTime().Format(time.RFC3339)
	}

	sourceHash := ""
	if bodyHash, err := os.ReadFile(path); err == nil {
		sourceHash = fmt.Sprintf("sha256:%x", sha256.Sum256(bodyHash))
	}

	return Card{
		ID:               strings.TrimSpace(fm.ID),
		SourcePath:       path,
		SourceModifiedAt: modifiedAt,
		SourceHash:       sourceHash,
		Enabled:          enabled,
		Title:            localizeText(fm.TitleEN, fm.Title, fm.TitleZH),
		TitleZH:          localizeText(fm.TitleZH, fm.Title, fm.TitleEN),
		TitleEN:          localizeText(fm.TitleEN, fm.Title, fm.TitleZH),
		QuestionType:     questionType,
		BodyFormat:       fm.BodyFormat,
		Tags:             normalizeTags(fm.Tags),
		Difficulty:       fm.Difficulty,
		QuestionText:     localizeText(fm.QuestionEN, fm.Question, fm.QuestionZH),
		QuestionTextZH:   localizeText(fm.QuestionZH, fm.Question, fm.QuestionEN),
		QuestionTextEN:   localizeText(fm.QuestionEN, fm.Question, fm.QuestionZH),
		Choices:          choicesEN,
		ChoicesZH:        choicesZH,
		ChoicesEN:        choicesEN,
		AnswerValue:      answer,
		Clickbait:        localizeText(fm.ClickbaitEN, fm.Clickbait, fm.ClickbaitZH),
		ClickbaitZH:      localizeText(fm.ClickbaitZH, fm.Clickbait, fm.ClickbaitEN),
		ClickbaitEN:      localizeText(fm.ClickbaitEN, fm.Clickbait, fm.ClickbaitZH),
		ReviewHint:       localizeText(fm.ReviewHintEN, fm.ReviewHint, fm.ReviewHintZH),
		ReviewHintZH:     localizeText(fm.ReviewHintZH, fm.ReviewHint, fm.ReviewHintEN),
		ReviewHintEN:     localizeText(fm.ReviewHintEN, fm.ReviewHint, fm.ReviewHintZH),
		ConfusionWith:    normalizeTags(fm.ConfusionWith),
		MetaphorSeed:     normalizeList(fm.MetaphorSeed),
		HookStyleTags:    normalizeTags(fm.HookStyleTags),
		BodyMarkdownZH:   zh,
		BodyMarkdownEN:   en,
		BodyPlaintextZH:  toPlaintext(zh),
		BodyPlaintextEN:  toPlaintext(en),
	}, diagnostics, nil
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
		Severity:   "error",
		Code:       code,
		Field:      field,
		Message:    message,
	}
}

func warning(path, code, field, message string) ImportError {
	return ImportError{
		SourcePath: path,
		Severity:   "warning",
		Code:       code,
		Field:      field,
		Message:    message,
	}
}

func localizedFieldWarnings(path string, fm frontmatter) []ImportError {
	var out []ImportError

	if strings.TrimSpace(fm.Title) != "" {
		if strings.TrimSpace(fm.TitleZH) == "" {
			out = append(out, warning(path, "missing_localized_field", "title_zh", "Missing localized field 'title_zh'; using fallback title value."))
		}
		if strings.TrimSpace(fm.TitleEN) == "" {
			out = append(out, warning(path, "missing_localized_field", "title_en", "Missing localized field 'title_en'; using fallback title value."))
		}
	}

	if strings.TrimSpace(fm.Question) != "" {
		if strings.TrimSpace(fm.QuestionZH) == "" {
			out = append(out, warning(path, "missing_localized_field", "question_zh", "Missing localized field 'question_zh'; using fallback question value."))
		}
		if strings.TrimSpace(fm.QuestionEN) == "" {
			out = append(out, warning(path, "missing_localized_field", "question_en", "Missing localized field 'question_en'; using fallback question value."))
		}
	}

	return out
}

func localizedChoiceWarnings(path string, fm frontmatter, choicesZH, choicesEN []string) []ImportError {
	var out []ImportError

	if strings.ToLower(strings.TrimSpace(fm.Type)) != "single-choice" {
		return out
	}

	if len(fm.Choices) > 0 {
		if len(fm.ChoicesZH) == 0 {
			out = append(out, warning(path, "missing_localized_field", "choices_zh", "Missing localized field 'choices_zh'; using fallback choices."))
		}
		if len(fm.ChoicesEN) == 0 {
			out = append(out, warning(path, "missing_localized_field", "choices_en", "Missing localized field 'choices_en'; using fallback choices."))
		}
	}

	if len(choicesZH) > 0 && len(choicesEN) > 0 && len(choicesZH) != len(choicesEN) {
		out = append(out, warning(path, "bilingual_choice_count_mismatch", "choices", fmt.Sprintf("Localized choice counts differ: zh-TW=%d, en=%d.", len(choicesZH), len(choicesEN))))
	}

	return out
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

func normalizeList(items []string) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item != "" {
			out = append(out, item)
		}
	}
	return out
}

func localizeText(primary string, fallback string, secondary string) string {
	for _, candidate := range []string{primary, fallback, secondary} {
		if trimmed := strings.TrimSpace(candidate); trimmed != "" {
			return trimmed
		}
	}
	return ""
}

func localizeChoices(primary []string, fallback []string, secondary []string) []string {
	for _, candidate := range [][]string{primary, fallback, secondary} {
		if len(candidate) == 0 {
			continue
		}

		out := make([]string, 0, len(candidate))
		for _, item := range candidate {
			out = append(out, strings.TrimSpace(item))
		}
		return out
	}

	return nil
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

func writeGOB(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return gob.NewEncoder(file).Encode(v)
}

func LoadCache(path string) (CacheFile, error) {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".gob":
		file, err := os.Open(path)
		if err != nil {
			return CacheFile{}, err
		}
		defer file.Close()

		var cache CacheFile
		if err := gob.NewDecoder(file).Decode(&cache); err != nil {
			return CacheFile{}, err
		}
		return cache, nil
	default:
		bytes, err := os.ReadFile(path)
		if err != nil {
			return CacheFile{}, err
		}
		bytes = []byte(strings.TrimPrefix(string(bytes), "\uFEFF"))

		var cache CacheFile
		if err := json.Unmarshal(bytes, &cache); err != nil {
			return CacheFile{}, err
		}

		return cache, nil
	}
}

func SnapshotKnowledgeFiles(paths []string) ([]CacheSource, error) {
	files, err := ListMarkdownFiles(paths)
	if err != nil {
		return nil, err
	}

	sources := make([]CacheSource, 0, len(files))
	for _, path := range files {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		sources = append(sources, CacheSource{
			Path:             filepath.Clean(path),
			ModifiedUnixNano: info.ModTime().UnixNano(),
			Size:             info.Size(),
		})
	}
	return sources, nil
}

func fingerprintKnowledgeFiles(sources []CacheSource) string {
	h := sha256.New()
	for _, source := range sources {
		fmt.Fprintf(h, "%s|%d|%d\n", source.Path, source.ModifiedUnixNano, source.Size)
	}
	return fmt.Sprintf("sha256:%x", h.Sum(nil))
}
