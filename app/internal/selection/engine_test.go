package selection

import (
	"testing"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/progress"
)

func TestPriorityScoreGivesNewCardBoost(t *testing.T) {
	now := fixedNow()
	card := cards.Card{ID: "git-new-card", Enabled: true, Tags: []string{"git"}}

	score := PriorityScore(card, progress.CardProgress{}, now)

	if score < 30 {
		t.Fatalf("expected new card score to include base new-card boost, got %d", score)
	}
}

func TestPriorityScoreBoostsReviewDueAndRecentWrong(t *testing.T) {
	now := fixedNow()
	lastWrong := now.Add(-2 * time.Hour).Format(time.RFC3339)
	nextReview := now.Add(-30 * time.Minute).Format(time.RFC3339)
	lastSeen := now.Add(-48 * time.Hour).Format(time.RFC3339)

	score := PriorityScore(cards.Card{ID: "git-review", Enabled: true}, progress.CardProgress{
		SeenCount:    3,
		CorrectCount: 1,
		WrongCount:   2,
		MasteryScore: -1,
		LastWrongAt:  &lastWrong,
		LastSeenAt:   &lastSeen,
		NextReviewAt: &nextReview,
		IsMastered:   false,
	}, now)

	if score < 50 {
		t.Fatalf("expected overdue + recent wrong card to have high score, got %d", score)
	}
}

func TestPriorityScoreAppliesMasteryPenalty(t *testing.T) {
	now := fixedNow()
	lastSeen := now.Add(-48 * time.Hour).Format(time.RFC3339)

	weakScore := PriorityScore(cards.Card{ID: "git-weak", Enabled: true}, progress.CardProgress{
		SeenCount:    5,
		CorrectCount: 2,
		WrongCount:   3,
		MasteryScore: 0,
		LastSeenAt:   &lastSeen,
	}, now)

	masteredScore := PriorityScore(cards.Card{ID: "git-mastered", Enabled: true}, progress.CardProgress{
		SeenCount:    8,
		CorrectCount: 8,
		WrongCount:   0,
		MasteryScore: 8,
		LastSeenAt:   &lastSeen,
		IsMastered:   true,
	}, now)

	if masteredScore >= weakScore {
		t.Fatalf("expected mastered score (%d) to be lower than weak score (%d)", masteredScore, weakScore)
	}
}

func TestSelectNextCardReturnsHighestPriorityEnabledCard(t *testing.T) {
	now := fixedNow()
	overdue := now.Add(-1 * time.Hour).Format(time.RFC3339)
	lastWrong := now.Add(-15 * time.Minute).Format(time.RFC3339)

	candidate, ok := SelectNextCard([]cards.Card{
		{ID: "git-disabled", Enabled: false},
		{ID: "git-mastered", Enabled: true},
		{ID: "git-review", Enabled: true},
	}, map[string]progress.CardProgress{
		"git-mastered": {
			SeenCount:    8,
			CorrectCount: 8,
			WrongCount:   0,
			MasteryScore: 8,
			IsMastered:   true,
		},
		"git-review": {
			SeenCount:    4,
			CorrectCount: 2,
			WrongCount:   2,
			MasteryScore: 0,
			NextReviewAt: &overdue,
			LastWrongAt:  &lastWrong,
		},
	}, now)

	if !ok {
		t.Fatal("expected a candidate card")
	}

	if candidate.ID != "git-review" {
		t.Fatalf("expected git-review, got %s", candidate.ID)
	}
}

func TestSelectNextCardReturnsFalseWhenNoEnabledCardsExist(t *testing.T) {
	now := fixedNow()

	_, ok := SelectNextCard([]cards.Card{
		{ID: "git-disabled", Enabled: false},
	}, map[string]progress.CardProgress{}, now)

	if ok {
		t.Fatal("expected no candidate")
	}
}

func TestSelectNextCardWorksWithParsedCardsAndProgressState(t *testing.T) {
	now := fixedNow()
	reviewDue := now.Add(-10 * time.Minute).Format(time.RFC3339)

	allCards := []cards.Card{
		{
			ID:      "git-rebase-vs-merge",
			Enabled: true,
			Tags:    []string{"git", "branching"},
		},
		{
			ID:      "git-cherry-pick-purpose",
			Enabled: true,
			Tags:    []string{"git", "commits"},
		},
	}

	states := map[string]progress.CardProgress{
		"git-rebase-vs-merge": {
			SeenCount:    3,
			CorrectCount: 3,
			WrongCount:   0,
			MasteryScore: 5,
		},
		"git-cherry-pick-purpose": {
			SeenCount:    2,
			CorrectCount: 1,
			WrongCount:   1,
			MasteryScore: 0,
			NextReviewAt: &reviewDue,
		},
	}

	card, ok := SelectNextCard(allCards, states, now)
	if !ok {
		t.Fatal("expected candidate")
	}

	if card.ID != "git-cherry-pick-purpose" {
		t.Fatalf("expected due review card, got %s", card.ID)
	}
}

func TestSelectNextCardAvoidsShortTermRepeatWhenAlternativeExists(t *testing.T) {
	now := fixedNow()
	recentlySeen := now.Add(-5 * time.Minute).Format(time.RFC3339)
	olderSeen := now.Add(-48 * time.Hour).Format(time.RFC3339)

	card, ok := SelectNextCard([]cards.Card{
		{ID: "git-recent-repeat", Enabled: true},
		{ID: "git-alternative", Enabled: true},
	}, map[string]progress.CardProgress{
		"git-recent-repeat": {
			SeenCount:    3,
			CorrectCount: 2,
			WrongCount:   1,
			MasteryScore: 0,
			LastSeenAt:   &recentlySeen,
		},
		"git-alternative": {
			SeenCount:    3,
			CorrectCount: 2,
			WrongCount:   1,
			MasteryScore: 0,
			LastSeenAt:   &olderSeen,
		},
	}, now)
	if !ok {
		t.Fatal("expected candidate")
	}

	if card.ID != "git-alternative" {
		t.Fatalf("expected alternative card, got %s", card.ID)
	}
}

func TestSelectNextCardStillReturnsRecentCardWhenItIsOnlyOption(t *testing.T) {
	now := fixedNow()
	recentlySeen := now.Add(-2 * time.Minute).Format(time.RFC3339)

	card, ok := SelectNextCard([]cards.Card{
		{ID: "git-only-card", Enabled: true},
	}, map[string]progress.CardProgress{
		"git-only-card": {
			SeenCount:    1,
			CorrectCount: 1,
			WrongCount:   0,
			MasteryScore: 0,
			LastSeenAt:   &recentlySeen,
		},
	}, now)
	if !ok {
		t.Fatal("expected candidate")
	}

	if card.ID != "git-only-card" {
		t.Fatalf("expected only card, got %s", card.ID)
	}
}

func fixedNow() time.Time {
	return time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
}
