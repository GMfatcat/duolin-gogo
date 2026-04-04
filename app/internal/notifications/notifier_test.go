package notifications

import (
	"testing"

	"duolin-gogo/internal/cards"
)

func TestBuildStudyMessageUsesLocalizedClickbaitAndCardID(t *testing.T) {
	message := BuildStudyMessageForLanguage(cards.Card{
		ID:             "git-rebase-vs-merge",
		TitleZH:        "Rebase vs Merge",
		TitleEN:        "Rebase vs Merge",
		QuestionTextZH: "git rebase 主要是在做什麼？",
		QuestionTextEN: "What does git rebase mainly do?",
		ClickbaitZH:    "多數人其實沒搞懂 rebase",
		ClickbaitEN:    "Most Git beginners misunderstand rebase. Do you?",
	}, "zh-TW")

	if message.Title != "多數人其實沒搞懂 rebase" {
		t.Fatalf("unexpected title: %s", message.Title)
	}

	if message.Body != "git rebase 主要是在做什麼？" {
		t.Fatalf("unexpected body: %s", message.Body)
	}

	if message.ActivationArgument != "duolin-gogo://study/git-rebase-vs-merge" {
		t.Fatalf("unexpected activation argument: %s", message.ActivationArgument)
	}
}
