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

func TestSelectNextCardForTopicFiltersCandidates(t *testing.T) {
	now := fixedNow()

	card, ok := SelectNextCardForTopic([]cards.Card{
		{ID: "git-card", Enabled: true, Tags: []string{"git"}},
		{ID: "docker-card", Enabled: true, Tags: []string{"docker"}},
	}, map[string]progress.CardProgress{}, "docker", now)
	if !ok {
		t.Fatal("expected candidate")
	}

	if card.ID != "docker-card" {
		t.Fatalf("expected docker card, got %s", card.ID)
	}
}

func TestSelectNextCardInMixedModePrefersWeakerTopicWhenScoresAreClose(t *testing.T) {
	now := fixedNow()
	olderSeen := now.Add(-48 * time.Hour).Format(time.RFC3339)
	recentlySeen := now.Add(-3 * time.Minute).Format(time.RFC3339)

	card, ok := SelectNextCardForTopic([]cards.Card{
		{ID: "git-card", Enabled: true, Tags: []string{"git"}},
		{ID: "git-second", Enabled: true, Tags: []string{"git"}},
		{ID: "docker-card", Enabled: true, Tags: []string{"docker"}},
		{ID: "docker-second", Enabled: true, Tags: []string{"docker"}},
	}, map[string]progress.CardProgress{
		"git-card": {
			SeenCount:    4,
			CorrectCount: 2,
			WrongCount:   2,
			MasteryScore: 1,
			LastSeenAt:   &olderSeen,
		},
		"git-second": {
			SeenCount:    6,
			CorrectCount: 6,
			WrongCount:   0,
			MasteryScore: 6,
			LastSeenAt:   &recentlySeen,
			IsMastered:   true,
		},
		"docker-card": {
			SeenCount:    4,
			CorrectCount: 2,
			WrongCount:   2,
			MasteryScore: 1,
			LastSeenAt:   &olderSeen,
		},
		"docker-second": {
			SeenCount:    4,
			CorrectCount: 0,
			WrongCount:   4,
			MasteryScore: 0,
			LastSeenAt:   &recentlySeen,
		},
	}, "all", now)
	if !ok {
		t.Fatal("expected candidate")
	}

	if card.ID != "docker-card" {
		t.Fatalf("expected mixed mode to prefer docker-card due to weaker topic, got %s", card.ID)
	}
}

func TestSelectNextCardForFocusedTopicIgnoresCrossTopicWeaknessBoost(t *testing.T) {
	now := fixedNow()
	olderSeen := now.Add(-48 * time.Hour).Format(time.RFC3339)

	card, ok := SelectNextCardForTopic([]cards.Card{
		{ID: "git-stronger-card", Enabled: true, Tags: []string{"git"}},
		{ID: "git-weaker-card", Enabled: true, Tags: []string{"git"}},
		{ID: "docker-card", Enabled: true, Tags: []string{"docker"}},
	}, map[string]progress.CardProgress{
		"git-stronger-card": {
			SeenCount:    4,
			CorrectCount: 1,
			WrongCount:   3,
			MasteryScore: 0,
			LastSeenAt:   &olderSeen,
		},
		"git-weaker-card": {
			SeenCount:    4,
			CorrectCount: 0,
			WrongCount:   4,
			MasteryScore: 0,
			LastSeenAt:   &olderSeen,
		},
		"docker-card": {
			SeenCount:    4,
			CorrectCount: 0,
			WrongCount:   4,
			MasteryScore: 0,
			LastSeenAt:   &olderSeen,
		},
	}, "git", now)
	if !ok {
		t.Fatal("expected candidate")
	}

	if card.ID != "git-weaker-card" {
		t.Fatalf("expected git-weaker-card within focused topic, got %s", card.ID)
	}
}

func TestFilterCardsByTopicReturnsOnlyMatchingCards(t *testing.T) {
	filtered := FilterCardsByTopic([]cards.Card{
		{ID: "git-card", Enabled: true, Tags: []string{"git"}},
		{ID: "python-card", Enabled: true, Tags: []string{"python"}},
		{ID: "docker-card", Enabled: true, SourcePath: "D:/duolin-gogo/knowledge/docker/run.md"},
	}, "docker")

	if len(filtered) != 1 {
		t.Fatalf("expected 1 filtered card, got %d", len(filtered))
	}

	if filtered[0].ID != "docker-card" {
		t.Fatalf("expected docker-card, got %s", filtered[0].ID)
	}
}

func TestFilterCardsByTopicSupportsTopicPresetGroups(t *testing.T) {
	filtered := FilterCardsByTopic([]cards.Card{
		{ID: "git-card", Enabled: true, Tags: []string{"git"}},
		{ID: "docker-card", Enabled: true, Tags: []string{"docker"}},
		{ID: "linux-card", Enabled: true, Tags: []string{"linux"}},
		{ID: "go-card", Enabled: true, Tags: []string{"go"}},
		{ID: "python-card", Enabled: true, Tags: []string{"python"}},
	}, "backend-tools")

	if len(filtered) != 3 {
		t.Fatalf("expected 3 backend tools cards, got %d", len(filtered))
	}

	filtered = FilterCardsByTopic([]cards.Card{
		{ID: "git-card", Enabled: true, Tags: []string{"git"}},
		{ID: "docker-card", Enabled: true, Tags: []string{"docker"}},
		{ID: "linux-card", Enabled: true, Tags: []string{"linux"}},
		{ID: "go-card", Enabled: true, Tags: []string{"go"}},
		{ID: "python-card", Enabled: true, Tags: []string{"python"}},
	}, "languages")

	if len(filtered) != 2 {
		t.Fatalf("expected 2 language cards, got %d", len(filtered))
	}
}

func fixedNow() time.Time {
	return time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
}
