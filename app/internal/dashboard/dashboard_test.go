package dashboard

import (
	"testing"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/progress"
)

func TestBuildSummaryCalculatesStatsAndNextReview(t *testing.T) {
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	nextReview := now.Add(2 * time.Hour).Format(time.RFC3339)

	summary := BuildSummary([]cards.Card{
		{ID: "git-rebase-vs-merge", Tags: []string{"git", "branching"}},
	}, progress.ProgressFile{
		Cards: map[string]progress.CardProgress{
			"git-rebase-vs-merge": {
				SeenCount:    3,
				CorrectCount: 2,
				WrongCount:   1,
				NextReviewAt: &nextReview,
			},
		},
		DailySummary: map[string]progress.DailySummary{
			"2026-04-05": {
				Answered: 4,
				Correct:  3,
			},
		},
	}, now)

	if summary.StudiedToday != 4 {
		t.Fatalf("expected studied today 4, got %d", summary.StudiedToday)
	}

	if summary.CorrectRate != 0.75 {
		t.Fatalf("expected correct rate 0.75, got %.2f", summary.CorrectRate)
	}

	if summary.NextReviewAt == "" {
		t.Fatal("expected next review time")
	}
}

func TestBuildSummaryReturnsWeakTopicsOrderedByErrorRate(t *testing.T) {
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	summary := BuildSummary([]cards.Card{
		{ID: "git-rebase-vs-merge", Tags: []string{"git", "branching"}},
		{ID: "git-cherry-pick-purpose", Tags: []string{"git", "commits"}},
		{ID: "git-fetch-basic", Tags: []string{"git", "branching"}},
	}, progress.ProgressFile{
		Cards: map[string]progress.CardProgress{
			"git-rebase-vs-merge": {
				SeenCount:    4,
				CorrectCount: 1,
				WrongCount:   3,
			},
			"git-cherry-pick-purpose": {
				SeenCount:    4,
				CorrectCount: 3,
				WrongCount:   1,
			},
			"git-fetch-basic": {
				SeenCount:    2,
				CorrectCount: 0,
				WrongCount:   2,
			},
		},
	}, now)

	if len(summary.WeakTopics) == 0 {
		t.Fatal("expected weak topics")
	}

	if summary.WeakTopics[0].Tag != "branching" {
		t.Fatalf("expected branching to be weakest tag, got %s", summary.WeakTopics[0].Tag)
	}
}

func TestBuildSummaryReturnsTopicProgress(t *testing.T) {
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	summary := BuildSummary([]cards.Card{
		{ID: "git-rebase-vs-merge", Tags: []string{"git", "branching"}, SourcePath: "D:/duolin-gogo/knowledge/git/rebase.md"},
		{ID: "docker-run-start-container", Tags: []string{"docker", "container"}, SourcePath: "D:/duolin-gogo/knowledge/docker/run.md"},
	}, progress.ProgressFile{
		Cards: map[string]progress.CardProgress{
			"git-rebase-vs-merge": {
				SeenCount:    4,
				CorrectCount: 3,
				WrongCount:   1,
			},
			"docker-run-start-container": {
				SeenCount:    2,
				CorrectCount: 1,
				WrongCount:   1,
			},
		},
	}, now)

	if len(summary.TopicProgress) != 2 {
		t.Fatalf("expected 2 topic progress items, got %d", len(summary.TopicProgress))
	}

	if summary.TopicProgress[0].Topic != "docker" {
		t.Fatalf("expected docker to sort before git by lower accuracy, got %s", summary.TopicProgress[0].Topic)
	}

	if summary.TopicProgress[1].Topic != "git" {
		t.Fatalf("expected git topic, got %s", summary.TopicProgress[1].Topic)
	}

	if summary.WeakestDeck == nil {
		t.Fatal("expected weakest deck insight")
	}

	if summary.WeakestDeck.Topic != "docker" {
		t.Fatalf("expected docker as weakest deck, got %s", summary.WeakestDeck.Topic)
	}
}

func TestBuildSummaryOmitsWeakestDeckWhenOnlyOneDeckIsPresent(t *testing.T) {
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	summary := BuildSummary([]cards.Card{
		{ID: "git-rebase-vs-merge", Tags: []string{"git", "branching"}, SourcePath: "D:/duolin-gogo/knowledge/git/rebase.md"},
	}, progress.ProgressFile{
		Cards: map[string]progress.CardProgress{
			"git-rebase-vs-merge": {
				SeenCount:    4,
				CorrectCount: 3,
				WrongCount:   1,
			},
		},
	}, now)

	if summary.WeakestDeck != nil {
		t.Fatalf("expected no weakest deck for single-topic summary, got %#v", summary.WeakestDeck)
	}
}
