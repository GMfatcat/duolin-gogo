package notifications

import (
	"testing"

	"duolin-gogo/internal/cards"
)

func TestBuildStudyMessageUsesClickbaitAndCardID(t *testing.T) {
	message := BuildStudyMessage(cards.Card{
		ID:           "git-rebase-vs-merge",
		Title:        "Rebase vs Merge",
		QuestionText: "What does git rebase mainly do?",
		Clickbait:    "Most Git beginners misunderstand rebase. Do you?",
	})

	if message.Title != "Most Git beginners misunderstand rebase. Do you?" {
		t.Fatalf("unexpected title: %s", message.Title)
	}

	if message.ActivationArgument != "duolin-gogo://study/git-rebase-vs-merge" {
		t.Fatalf("unexpected activation argument: %s", message.ActivationArgument)
	}
}
