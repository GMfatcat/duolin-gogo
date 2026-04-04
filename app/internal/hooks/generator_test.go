package hooks

import (
	"strings"
	"testing"

	"duolin-gogo/internal/cards"
)

func TestGenerateUsesMetaphorAndStyleDeterministically(t *testing.T) {
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
