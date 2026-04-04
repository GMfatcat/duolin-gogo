package cards

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestScanDirectoriesParsesValidBilingualCard(t *testing.T) {
	dir := t.TempDir()
	knowledgeDir := filepath.Join(dir, "knowledge", "git")

	if err := os.MkdirAll(knowledgeDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	cardPath := filepath.Join(knowledgeDir, "rebase.md")
	content := `---
id: git-rebase-vs-merge
title: Rebase vs Merge
title_zh: Rebase 跟 Merge 的差別
title_en: Rebase vs Merge
type: single-choice
body_format: bilingual-section
tags: [git, branching]
difficulty: 2
question_zh: "git rebase 主要是在做什麼？"
question_en: "What does git rebase mainly do?"
choices_zh:
  - "建立一個 merge commit"
  - "把 commits 重新接到新的 base 上"
choices_en:
  - "Creates a merge commit"
  - "Replays commits onto a new base"
answer: 1
clickbait_zh: "多數 Git 初學者其實沒搞懂 rebase"
clickbait_en: "Most Git beginners misunderstand rebase. Do you?"
review_hint_zh: "Rebase = 把 commits 重放到新的 base 上。"
review_hint_en: "Rebase = replay commits on top of another base."
enabled: true
---

## zh-TW

` + "`git rebase` 會把目前分支上的 commits 重新接到另一個 base 上。" + `

## en

` + "`git rebase` takes commits from your current branch and reapplies them onto another base commit." + `
`

	if err := os.WriteFile(cardPath, []byte(content), 0o644); err != nil {
		t.Fatalf("write file failed: %v", err)
	}

	result, err := ScanDirectories([]string{filepath.Join(dir, "knowledge")})
	if err != nil {
		t.Fatalf("scan failed: %v", err)
	}

	if len(result.Errors) != 0 {
		t.Fatalf("expected no import errors, got %d", len(result.Errors))
	}

	if len(result.Cards) != 1 {
		t.Fatalf("expected 1 card, got %d", len(result.Cards))
	}

	card := result.Cards[0]
	if card.ID != "git-rebase-vs-merge" {
		t.Fatalf("unexpected card id: %s", card.ID)
	}

	if card.BodyMarkdownZH == "" || card.BodyMarkdownEN == "" {
		t.Fatal("expected bilingual body sections to be parsed")
	}

	if card.QuestionType != "single-choice" {
		t.Fatalf("unexpected question type: %s", card.QuestionType)
	}

	if len(card.ChoicesEN) != 2 {
		t.Fatalf("expected 2 English choices, got %d", len(card.ChoicesEN))
	}

	if card.QuestionTextZH != "git rebase 主要是在做什麼？" {
		t.Fatalf("unexpected zh question: %s", card.QuestionTextZH)
	}

	if card.ClickbaitZH != "多數 Git 初學者其實沒搞懂 rebase" {
		t.Fatalf("unexpected zh clickbait: %s", card.ClickbaitZH)
	}

	if len(card.ChoicesZH) != 2 || card.ChoicesZH[1] != "把 commits 重新接到新的 base 上" {
		t.Fatalf("unexpected zh choices: %#v", card.ChoicesZH)
	}
}

func TestScanDirectoriesReportsMissingLanguageSection(t *testing.T) {
	dir := t.TempDir()
	knowledgeDir := filepath.Join(dir, "knowledge", "git")

	if err := os.MkdirAll(knowledgeDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	cardPath := filepath.Join(knowledgeDir, "broken.md")
	content := `---
id: git-broken-card
title: Broken Card
type: true-false
question: "Broken?"
answer: true
---

## zh-TW

只有中文，沒有英文。`

	if err := os.WriteFile(cardPath, []byte(content), 0o644); err != nil {
		t.Fatalf("write file failed: %v", err)
	}

	result, err := ScanDirectories([]string{filepath.Join(dir, "knowledge")})
	if err != nil {
		t.Fatalf("scan failed: %v", err)
	}

	if len(result.Cards) != 0 {
		t.Fatalf("expected 0 cards, got %d", len(result.Cards))
	}

	if len(result.Errors) != 1 {
		t.Fatalf("expected 1 import error, got %d", len(result.Errors))
	}

	if result.Errors[0].Code != "missing_language_section" {
		t.Fatalf("unexpected error code: %s", result.Errors[0].Code)
	}
}

func TestWriteOutputsPersistsCacheAndErrors(t *testing.T) {
	dir := t.TempDir()
	cachePath := filepath.Join(dir, "cards-cache.json")
	errorPath := filepath.Join(dir, "import-errors.json")

	result := ImportResult{
		Cards: []Card{
			{
				ID:             "git-cherry-pick-purpose",
				Title:          "Cherry-pick Purpose",
				TitleZH:        "Cherry-pick 的用途",
				TitleEN:        "Cherry-pick Purpose",
				QuestionType:   "true-false",
				QuestionText:   "Cherry-pick copies a chosen commit onto the current branch.",
				QuestionTextZH: "git cherry-pick 會把指定 commit 套到目前分支。",
				QuestionTextEN: "Cherry-pick copies a chosen commit onto the current branch.",
				AnswerValue:    true,
				BodyFormat:     "bilingual-section",
				BodyMarkdownZH: "中文說明",
				BodyMarkdownEN: "English",
				Enabled:        true,
			},
		},
		Errors: []ImportError{
			{
				SourcePath: "D:\\duolin-gogo\\knowledge\\git\\broken.md",
				Code:       "missing_language_section",
				Field:      "body",
				Message:    "Missing ## en section.",
			},
		},
	}

	if err := WriteCache(cachePath, result.Cards); err != nil {
		t.Fatalf("write cache failed: %v", err)
	}

	if err := WriteImportErrors(errorPath, result.Errors); err != nil {
		t.Fatalf("write import errors failed: %v", err)
	}

	cacheBytes, err := os.ReadFile(cachePath)
	if err != nil {
		t.Fatalf("read cache failed: %v", err)
	}

	var cache CacheFile
	if err := json.Unmarshal(cacheBytes, &cache); err != nil {
		t.Fatalf("unmarshal cache failed: %v", err)
	}

	if len(cache.Cards) != 1 {
		t.Fatalf("expected 1 cached card, got %d", len(cache.Cards))
	}

	errorBytes, err := os.ReadFile(errorPath)
	if err != nil {
		t.Fatalf("read import errors failed: %v", err)
	}

	var errorsFile ImportErrorsFile
	if err := json.Unmarshal(errorBytes, &errorsFile); err != nil {
		t.Fatalf("unmarshal errors failed: %v", err)
	}

	if len(errorsFile.Errors) != 1 {
		t.Fatalf("expected 1 import error, got %d", len(errorsFile.Errors))
	}
}

func TestWriteImportErrorsUsesEmptyArrayForNoErrors(t *testing.T) {
	dir := t.TempDir()
	errorPath := filepath.Join(dir, "import-errors.json")

	if err := WriteImportErrors(errorPath, nil); err != nil {
		t.Fatalf("write import errors failed: %v", err)
	}

	errorBytes, err := os.ReadFile(errorPath)
	if err != nil {
		t.Fatalf("read import errors failed: %v", err)
	}

	var errorsFile ImportErrorsFile
	if err := json.Unmarshal(errorBytes, &errorsFile); err != nil {
		t.Fatalf("unmarshal errors failed: %v", err)
	}

	if errorsFile.Errors == nil {
		t.Fatal("expected empty errors array, got nil")
	}

	if len(errorsFile.Errors) != 0 {
		t.Fatalf("expected 0 import errors, got %d", len(errorsFile.Errors))
	}
}

func TestRefreshKnowledgeWritesDataFiles(t *testing.T) {
	dir := t.TempDir()
	knowledgeDir := filepath.Join(dir, "knowledge", "git")
	dataDir := filepath.Join(dir, "data")

	if err := os.MkdirAll(knowledgeDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	validCard := `---
id: git-fetch-basic
title: Git Fetch
title_zh: Git Fetch
title_en: Git Fetch
type: true-false
question_zh: "` + "`git fetch` 只會更新 remote-tracking references，不會直接合併到目前分支。" + `"
question_en: "` + "`git fetch` updates local remote-tracking references without merging into the current branch." + `"
clickbait_zh: "你真的分得清 fetch 跟 pull 嗎？"
clickbait_en: "Do you actually know the difference between fetch and pull?"
answer: true
---

## zh-TW

` + "`git fetch` 會更新本地的 remote-tracking references，但不會直接合併到目前分支。" + `

## en

` + "`git fetch` updates local remote-tracking references without merging into the current branch." + `
`

	if err := os.WriteFile(filepath.Join(knowledgeDir, "fetch.md"), []byte(validCard), 0o644); err != nil {
		t.Fatalf("write valid card failed: %v", err)
	}

	invalidCard := `---
id: git-invalid-card
title: Invalid
type: true-false
question: "Broken?"
answer: false
---

## zh-TW

只有中文。`

	if err := os.WriteFile(filepath.Join(knowledgeDir, "broken.md"), []byte(invalidCard), 0o644); err != nil {
		t.Fatalf("write invalid card failed: %v", err)
	}

	result, err := RefreshKnowledge(filepath.Join(dir, "knowledge"), dataDir)
	if err != nil {
		t.Fatalf("refresh knowledge failed: %v", err)
	}

	if len(result.Cards) != 1 {
		t.Fatalf("expected 1 valid card, got %d", len(result.Cards))
	}

	if len(result.Errors) != 1 {
		t.Fatalf("expected 1 import error, got %d", len(result.Errors))
	}

	if _, err := os.Stat(filepath.Join(dataDir, "cards-cache.json")); err != nil {
		t.Fatalf("expected cards-cache.json to exist: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dataDir, "import-errors.json")); err != nil {
		t.Fatalf("expected import-errors.json to exist: %v", err)
	}
}

func TestScanDirectoriesReportsMissingLocalizedFieldsAsWarnings(t *testing.T) {
	dir := t.TempDir()
	knowledgeDir := filepath.Join(dir, "knowledge", "git")

	if err := os.MkdirAll(knowledgeDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	cardPath := filepath.Join(knowledgeDir, "legacy.md")
	content := `---
id: git-legacy-card
title: Legacy Card
type: single-choice
question: "Which command updates your index before commit?"
choices:
  - "git add"
  - "git clone"
answer: 0
---

## zh-TW

這是一張只有 fallback 欄位的舊卡片。

## en

This is a legacy card that only uses fallback fields.
`

	if err := os.WriteFile(cardPath, []byte(content), 0o644); err != nil {
		t.Fatalf("write file failed: %v", err)
	}

	result, err := ScanDirectories([]string{filepath.Join(dir, "knowledge")})
	if err != nil {
		t.Fatalf("scan failed: %v", err)
	}

	if len(result.Cards) != 1 {
		t.Fatalf("expected 1 card, got %d", len(result.Cards))
	}

	if len(result.Errors) != 6 {
		t.Fatalf("expected 6 warnings, got %d", len(result.Errors))
	}

	for _, diagnostic := range result.Errors {
		if diagnostic.Severity != "warning" {
			t.Fatalf("expected warning severity, got %s", diagnostic.Severity)
		}
		if diagnostic.Code != "missing_localized_field" {
			t.Fatalf("unexpected warning code: %s", diagnostic.Code)
		}
	}
}

func TestScanDirectoriesReportsChoiceCountMismatchAsWarning(t *testing.T) {
	dir := t.TempDir()
	knowledgeDir := filepath.Join(dir, "knowledge", "git")

	if err := os.MkdirAll(knowledgeDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	cardPath := filepath.Join(knowledgeDir, "mismatch.md")
	content := `---
id: git-choice-mismatch
title_zh: 選項數量不一致
title_en: Choice Count Mismatch
type: single-choice
question_zh: "哪個指令會抓遠端更新？"
question_en: "Which command fetches remote updates?"
choices_zh:
  - "git fetch"
choices_en:
  - "git fetch"
  - "git pull"
answer: 0
---

## zh-TW

這張卡故意讓雙語選項數量不一致。

## en

This card intentionally has mismatched localized choice counts.
`

	if err := os.WriteFile(cardPath, []byte(content), 0o644); err != nil {
		t.Fatalf("write file failed: %v", err)
	}

	result, err := ScanDirectories([]string{filepath.Join(dir, "knowledge")})
	if err != nil {
		t.Fatalf("scan failed: %v", err)
	}

	if len(result.Cards) != 1 {
		t.Fatalf("expected 1 card, got %d", len(result.Cards))
	}

	if len(result.Errors) != 1 {
		t.Fatalf("expected 1 warning, got %d", len(result.Errors))
	}

	if result.Errors[0].Severity != "warning" {
		t.Fatalf("expected warning severity, got %s", result.Errors[0].Severity)
	}
	if result.Errors[0].Code != "bilingual_choice_count_mismatch" {
		t.Fatalf("unexpected warning code: %s", result.Errors[0].Code)
	}
}

func TestPreviewDraftReturnsNormalizedCardAndWarnings(t *testing.T) {
	raw := `---
id: git-draft-preview
title: Draft Preview
type: single-choice
question: "Which command stages changes?"
choices:
  - "git add"
  - "git log"
answer: 0
---

## zh-TW

這是一張只靠 fallback 欄位的草稿卡。

## en

This draft card relies on fallback fields.`

	result, err := PreviewDraft("draft://ai-card.md", raw)
	if err != nil {
		t.Fatalf("preview draft failed: %v", err)
	}

	if result.Card == nil {
		t.Fatal("expected normalized draft card")
	}
	if result.Card.ID != "git-draft-preview" {
		t.Fatalf("unexpected card id: %s", result.Card.ID)
	}
	if len(result.Errors) == 0 {
		t.Fatal("expected fallback warnings")
	}
	if result.Errors[0].Severity != "warning" {
		t.Fatalf("expected warning severity, got %s", result.Errors[0].Severity)
	}
}
