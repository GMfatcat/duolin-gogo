package review

import (
	"testing"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/progress"
	"duolin-gogo/internal/settings"
)

func TestShouldStartReviewDailyAtConfiguredTime(t *testing.T) {
	config := settings.File{}
	config.ReviewSchedule.Mode = "daily"
	config.ReviewSchedule.Time = "21:00"

	now := time.Date(2026, 4, 5, 21, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	if !ShouldStartReview(config, nil, now) {
		t.Fatal("expected daily review to start")
	}
}

func TestShouldStartReviewWeeklyOnConfiguredWeekday(t *testing.T) {
	config := settings.File{}
	config.ReviewSchedule.Mode = "weekly"
	weekday := "sun"
	config.ReviewSchedule.Weekday = &weekday
	config.ReviewSchedule.Time = "21:00"

	now := time.Date(2026, 4, 5, 21, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	if !ShouldStartReview(config, nil, now) {
		t.Fatal("expected weekly review to start")
	}
}

func TestShouldNotStartReviewTwiceSameWindow(t *testing.T) {
	config := settings.File{}
	config.ReviewSchedule.Mode = "daily"
	config.ReviewSchedule.Time = "21:00"

	now := time.Date(2026, 4, 5, 21, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	lastRun := now.Add(-10 * time.Minute)
	if ShouldStartReview(config, &lastRun, now) {
		t.Fatal("expected review to be blocked in same window")
	}
}

func TestBuildQueuePrioritizesDueAndWeakCards(t *testing.T) {
	now := time.Date(2026, 4, 5, 21, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	due := now.Add(-1 * time.Hour).Format(time.RFC3339)
	old := now.Add(-48 * time.Hour).Format(time.RFC3339)

	queue := BuildQueue([]cards.Card{
		{ID: "git-rebase-vs-merge", Enabled: true},
		{ID: "git-cherry-pick-purpose", Enabled: true},
		{ID: "git-fetch-basic", Enabled: true},
	}, map[string]progress.CardProgress{
		"git-rebase-vs-merge": {
			SeenCount:    4,
			CorrectCount: 1,
			WrongCount:   3,
			NextReviewAt: &due,
			LastSeenAt:   &old,
		},
		"git-cherry-pick-purpose": {
			SeenCount:    4,
			CorrectCount: 4,
			WrongCount:   0,
			NextReviewAt: &due,
		},
		"git-fetch-basic": {
			SeenCount: 0,
		},
	}, now, 2)

	if len(queue) != 2 {
		t.Fatalf("expected 2 cards, got %d", len(queue))
	}

	if queue[0].ID != "git-rebase-vs-merge" {
		t.Fatalf("expected weakest due card first, got %s", queue[0].ID)
	}
}

func TestBuildQueueSkipsUnseenCardsForReview(t *testing.T) {
	now := time.Date(2026, 4, 5, 21, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	queue := BuildQueue([]cards.Card{
		{ID: "git-new", Enabled: true},
	}, map[string]progress.CardProgress{
		"git-new": {SeenCount: 0},
	}, now, 5)

	if len(queue) != 0 {
		t.Fatalf("expected unseen cards to be excluded from review, got %d", len(queue))
	}
}
