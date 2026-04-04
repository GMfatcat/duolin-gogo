package cards

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanDirectoriesParsesHookMetadata(t *testing.T) {
	dir := t.TempDir()
	knowledgeDir := filepath.Join(dir, "knowledge", "git")

	if err := os.MkdirAll(knowledgeDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	content := `---
id: git-fetch-basic
title_zh: git fetch 的用途
title_en: Git Fetch
type: true-false
question_zh: "` + "`git fetch` 只會更新遠端追蹤資訊，不會直接 merge。" + `"
question_en: "` + "`git fetch` updates remote-tracking references without merging directly." + `"
answer: true
confusion_with: [git-pull-composition]
metaphor_seed: [先看貨, 先觀察]
hook_style_tags: [safer-first, misunderstood]
---

## zh-TW

這是中文說明。

## en

This is the English explanation.
`

	cardPath := filepath.Join(knowledgeDir, "fetch.md")
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

	card := result.Cards[0]
	if len(card.ConfusionWith) != 1 || card.ConfusionWith[0] != "git-pull-composition" {
		t.Fatalf("unexpected confusion_with: %#v", card.ConfusionWith)
	}
	if len(card.MetaphorSeed) != 2 || card.MetaphorSeed[0] != "先看貨" {
		t.Fatalf("unexpected metaphor_seed: %#v", card.MetaphorSeed)
	}
	if len(card.HookStyleTags) != 2 || card.HookStyleTags[0] != "safer-first" {
		t.Fatalf("unexpected hook_style_tags: %#v", card.HookStyleTags)
	}
}
