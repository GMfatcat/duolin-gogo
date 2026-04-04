package hooks

import (
	"strings"
	"testing"

	"duolin-gogo/internal/cards"
)

func TestGeneratePlayfulUsesMetaphorDeterministically(t *testing.T) {
	card := cards.Card{
		ID:             "git-fetch-basic",
		TitleZH:        "git fetch 的用途",
		TitleEN:        "Git Fetch",
		QuestionTextZH: "`git fetch` 只會更新遠端追蹤資訊。",
		QuestionTextEN: "`git fetch` updates remote-tracking references.",
		Tags:           []string{"git", "remote"},
		ConfusionWith:  []string{"git-pull-composition"},
		MetaphorSeed:   []string{"先看貨"},
		HookStyleTags:  []string{"safer-first", "misunderstood"},
	}

	title, body := Generate(card, "zh-TW", "playful")

	if title == "" || body == "" {
		t.Fatal("expected hook output")
	}
	if !strings.Contains(title, "先看貨") {
		t.Fatalf("expected metaphor in title, got %q", title)
	}
	if !strings.Contains(body, "git fetch") {
		t.Fatalf("expected card question/body context, got %q", body)
	}
}

func TestGenerateAggressiveUsesComparisonWhenAvailable(t *testing.T) {
	card := cards.Card{
		ID:             "git-pull-composition",
		TitleZH:        "git pull 的組成",
		TitleEN:        "Git Pull",
		QuestionTextZH: "`git pull` 通常等於什麼？",
		QuestionTextEN: "What is `git pull` usually made of?",
		Tags:           []string{"git", "remote"},
		ConfusionWith:  []string{"git-fetch-basic"},
		HookStyleTags:  []string{"comparison", "misunderstood"},
	}

	title, _ := Generate(card, "en", "aggressive")

	if !strings.Contains(strings.ToLower(title), "fetch basic") && !strings.Contains(strings.ToLower(title), "fetch") {
		t.Fatalf("expected comparison-oriented aggressive title, got %q", title)
	}
}

func TestGenerateChaoticFeelsMoreHeadlineLike(t *testing.T) {
	card := cards.Card{
		ID:             "git-checkout-legacy",
		TitleZH:        "git checkout 的舊式用途",
		TitleEN:        "Git Checkout",
		QuestionTextZH: "`git checkout` 可以做不只一件事。",
		QuestionTextEN: "`git checkout` can do more than one job.",
		Tags:           []string{"git", "branches"},
		MetaphorSeed:   []string{"瑞士刀"},
		HookStyleTags:  []string{"chaotic", "misunderstood"},
	}

	title, _ := Generate(card, "zh-TW", "chaotic")

	if !strings.Contains(title, "瑞士刀") && !strings.Contains(title, "事故") && !strings.Contains(title, "心測") {
		t.Fatalf("expected headline-like chaotic title, got %q", title)
	}
}

func TestGenerateFallsBackWhenMetadataIsSparse(t *testing.T) {
	card := cards.Card{
		ID:             "git-status-purpose",
		TitleZH:        "git status 的用途",
		TitleEN:        "Git Status",
		QuestionTextZH: "`git status` 最主要是在做什麼？",
		QuestionTextEN: "What does `git status` mainly do?",
		Tags:           []string{"git", "basics"},
	}

	title, body := Generate(card, "en", "aggressive")

	if title == "" || body == "" {
		t.Fatal("expected fallback hook output")
	}
	if !strings.Contains(strings.ToLower(title), "git") {
		t.Fatalf("expected title to retain technical context, got %q", title)
	}
}
