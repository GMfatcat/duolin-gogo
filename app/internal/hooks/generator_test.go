package hooks

import (
	"strings"
	"testing"

	"duolin-gogo/internal/cards"
)

func TestGeneratePlayfulUsesMetaphorDeterministically(t *testing.T) {
	card := cards.Card{
		ID:             "git-fetch-basic",
		TitleZH:        "Git Fetch",
		TitleEN:        "Git Fetch",
		QuestionTextZH: "`git fetch` 會更新遠端追蹤分支。",
		QuestionTextEN: "`git fetch` updates remote-tracking branches.",
		MetaphorSeed:   []string{"先看貨"},
		HookStyleTags:  []string{"safer-first"},
	}

	titleOne, bodyOne := Generate(card, "zh-TW", "playful")
	titleTwo, bodyTwo := Generate(card, "zh-TW", "playful")

	if titleOne != titleTwo {
		t.Fatalf("expected deterministic title, got %q and %q", titleOne, titleTwo)
	}
	if !strings.Contains(titleOne, "先看貨") {
		t.Fatalf("expected metaphor in title, got %q", titleOne)
	}
	if bodyOne != card.QuestionTextZH || bodyTwo != card.QuestionTextZH {
		t.Fatalf("expected localized body, got %q and %q", bodyOne, bodyTwo)
	}
}

func TestGenerateAggressiveUsesComparisonWhenAvailable(t *testing.T) {
	card := cards.Card{
		ID:             "git-pull-composition",
		TitleZH:        "Git Pull",
		TitleEN:        "Git Pull",
		QuestionTextZH: "`git pull` 其實是兩個動作的組合。",
		QuestionTextEN: "`git pull` is a composition of two Git actions.",
		ConfusionWith:  []string{"git-fetch-basic"},
		HookStyleTags:  []string{"comparison"},
	}

	title, _ := Generate(card, "zh-TW", "aggressive")

	if !strings.Contains(title, "fetch basic") {
		t.Fatalf("expected comparison label in title, got %q", title)
	}
}

func TestGenerateChaoticFeelsHeadlineLike(t *testing.T) {
	card := cards.Card{
		ID:             "git-stash-purpose",
		TitleZH:        "Git Stash",
		TitleEN:        "Git Stash",
		QuestionTextZH: "`git stash` 可以暫時收起尚未提交的改動。",
		QuestionTextEN: "`git stash` temporarily shelves uncommitted changes.",
		MetaphorSeed:   []string{"先藏起來"},
	}

	title, _ := Generate(card, "zh-TW", "chaotic")

	matches := []string{
		"事故開頭",
		"熱門誤用榜",
		"心測一下",
		"翻車",
	}
	for _, fragment := range matches {
		if strings.Contains(title, fragment) {
			return
		}
	}

	t.Fatalf("expected chaotic headline-like title, got %q", title)
}

func TestGenerateFallsBackToTechnicalContextWhenMetadataIsSparse(t *testing.T) {
	card := cards.Card{
		ID:             "git-status-purpose",
		TitleZH:        "Git Status",
		TitleEN:        "Git Status",
		QuestionTextZH: "`git status` 會顯示工作目錄與暫存區的狀態。",
		QuestionTextEN: "`git status` shows the state of the working tree and staging area.",
		HookStyleTags:  []string{"misunderstood"},
	}

	title, body := Generate(card, "en", "chaotic")

	if title == "" {
		t.Fatal("expected fallback title")
	}
	if !strings.Contains(strings.ToLower(title), "status") && !strings.Contains(strings.ToLower(title), "git") {
		t.Fatalf("expected technical context in title, got %q", title)
	}
	if body != card.QuestionTextEN {
		t.Fatalf("expected english body, got %q", body)
	}
}
