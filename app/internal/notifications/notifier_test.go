package notifications

import (
	"strings"
	"testing"

	"duolin-gogo/internal/cards"
)

func TestBuildStudyMessageUsesLocalizedManualClickbaitByDefault(t *testing.T) {
	message := BuildStudyMessageForLanguage(cards.Card{
		ID:             "git-rebase-vs-merge",
		TitleZH:        "Rebase 與 Merge 的差異",
		TitleEN:        "Rebase vs Merge",
		QuestionTextZH: "git rebase 最主要在做什麼？",
		QuestionTextEN: "What does git rebase mainly do?",
		ClickbaitZH:    "你真的分得清 rebase 和 merge 的差別嗎？",
		ClickbaitEN:    "Most Git beginners misunderstand rebase. Do you?",
	}, "zh-TW", "playful", "prefer_manual")

	if message.Title != "你真的分得清 rebase 和 merge 的差別嗎？" {
		t.Fatalf("unexpected title: %s", message.Title)
	}
	if message.Body != "" {
		t.Fatalf("expected empty body, got %q", message.Body)
	}
	if message.ActivationArgument != "duolin-gogo://study/git-rebase-vs-merge" {
		t.Fatalf("unexpected activation argument: %s", message.ActivationArgument)
	}
}

func TestBuildStudyMessageCanPreferGeneratedHook(t *testing.T) {
	message := BuildStudyMessageForLanguage(cards.Card{
		ID:            "git-fetch-basic",
		TitleZH:       "Git Fetch",
		TitleEN:       "Git Fetch",
		QuestionTextZH: "`git fetch` 會更新遠端追蹤參照。",
		QuestionTextEN: "`git fetch` updates remote-tracking references.",
		ClickbaitZH:   "這句手寫文案應該被 generator 蓋掉",
		ClickbaitEN:   "This manual line should be overridden",
		ConfusionWith: []string{"git-pull-composition"},
		MetaphorSeed:  []string{"先看貨"},
		HookStyleTags: []string{"safer-first"},
	}, "zh-TW", "playful", "prefer_generated")

	if strings.Contains(message.Title, "手寫文案") {
		t.Fatalf("expected generated title, got %q", message.Title)
	}
	if !strings.Contains(message.Title, "先看貨") {
		t.Fatalf("expected generated metaphor, got %q", message.Title)
	}
	if message.Body != "" {
		t.Fatalf("expected empty body, got %q", message.Body)
	}
}

func TestBuildStudyMessageGeneratesHookWhenManualClickbaitMissing(t *testing.T) {
	message := BuildStudyMessageForLanguage(cards.Card{
		ID:            "git-fetch-basic",
		TitleZH:       "Git Fetch",
		TitleEN:       "Git Fetch",
		QuestionTextZH: "`git fetch` 會更新遠端追蹤參照。",
		QuestionTextEN: "`git fetch` updates remote-tracking references.",
		ConfusionWith: []string{"git-pull-composition"},
		MetaphorSeed:  []string{"先看貨"},
		HookStyleTags: []string{"safer-first"},
	}, "zh-TW", "playful", "prefer_manual")

	if message.Title == "" {
		t.Fatal("expected generated title")
	}
	if message.Body != "" {
		t.Fatalf("expected empty body, got %q", message.Body)
	}
}
