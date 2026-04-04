package progress

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRecordAttemptCreatesProgressForNewCard(t *testing.T) {
	store := NewStore()
	shownAt := time.Date(2026, 4, 4, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	answeredAt := shownAt.Add(12 * time.Second)

	attempt, err := store.RecordAttempt(RecordAttemptInput{
		CardID:         "git-rebase-vs-merge",
		SessionType:    "learn",
		SelectedAnswer: 1,
		IsCorrect:      true,
		ShownAt:        shownAt,
		AnsweredAt:     answeredAt,
	})
	if err != nil {
		t.Fatalf("record attempt failed: %v", err)
	}

	card, ok := store.State.Cards["git-rebase-vs-merge"]
	if !ok {
		t.Fatal("expected progress state for card")
	}

	if card.SeenCount != 1 {
		t.Fatalf("expected seen_count 1, got %d", card.SeenCount)
	}

	if card.CorrectCount != 1 || card.WrongCount != 0 {
		t.Fatalf("unexpected counts: correct=%d wrong=%d", card.CorrectCount, card.WrongCount)
	}

	if card.MasteryScore != 1 {
		t.Fatalf("expected mastery score 1, got %d", card.MasteryScore)
	}

	if card.StreakCorrect != 1 {
		t.Fatalf("expected streak_correct 1, got %d", card.StreakCorrect)
	}

	if card.NextReviewAt == nil {
		t.Fatal("expected next_review_at to be set")
	}

	if attempt.MasteryDelta != 1 {
		t.Fatalf("expected mastery delta 1, got %d", attempt.MasteryDelta)
	}
}

func TestRecordAttemptUpdatesWrongAnswerAndDailySummary(t *testing.T) {
	store := NewStore()
	shownAt := time.Date(2026, 4, 4, 21, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	answeredAt := shownAt.Add(8 * time.Second)

	_, err := store.RecordAttempt(RecordAttemptInput{
		CardID:         "git-cherry-pick-purpose",
		SessionType:    "review",
		SelectedAnswer: true,
		IsCorrect:      false,
		ShownAt:        shownAt,
		AnsweredAt:     answeredAt,
	})
	if err != nil {
		t.Fatalf("record attempt failed: %v", err)
	}

	card := store.State.Cards["git-cherry-pick-purpose"]
	if card.WrongCount != 1 {
		t.Fatalf("expected wrong_count 1, got %d", card.WrongCount)
	}

	if card.MasteryScore != -2 {
		t.Fatalf("expected mastery score -2, got %d", card.MasteryScore)
	}

	if card.StreakCorrect != 0 {
		t.Fatalf("expected streak_correct reset to 0, got %d", card.StreakCorrect)
	}

	if card.LastSessionType != "review" {
		t.Fatalf("expected last session type review, got %s", card.LastSessionType)
	}

	day := store.State.DailySummary["2026-04-04"]
	if day.CardsSeen != 1 || day.Answered != 1 {
		t.Fatalf("unexpected daily totals: %+v", day)
	}

	if day.Wrong != 1 || day.ReviewAnswered != 1 {
		t.Fatalf("unexpected daily wrong/review totals: %+v", day)
	}
}

func TestSaveAndLoadStateRoundTrip(t *testing.T) {
	dir := t.TempDir()
	progressPath := filepath.Join(dir, "progress.json")
	attemptsPath := filepath.Join(dir, "attempts.jsonl")
	store := NewStore()

	shownAt := time.Date(2026, 4, 4, 11, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	answeredAt := shownAt.Add(10 * time.Second)

	attempt, err := store.RecordAttempt(RecordAttemptInput{
		CardID:         "git-fetch-basic",
		SessionType:    "learn",
		SelectedAnswer: true,
		IsCorrect:      true,
		ShownAt:        shownAt,
		AnsweredAt:     answeredAt,
	})
	if err != nil {
		t.Fatalf("record attempt failed: %v", err)
	}

	if err := store.Save(progressPath); err != nil {
		t.Fatalf("save failed: %v", err)
	}

	if err := AppendAttempt(attemptsPath, attempt); err != nil {
		t.Fatalf("append attempt failed: %v", err)
	}

	loaded, err := Load(progressPath)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	card := loaded.Cards["git-fetch-basic"]
	if card.CorrectCount != 1 {
		t.Fatalf("expected loaded correct_count 1, got %d", card.CorrectCount)
	}

	file, err := os.Open(attemptsPath)
	if err != nil {
		t.Fatalf("open attempts failed: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		t.Fatal("expected one attempt line")
	}

	var logged AttemptEvent
	if err := json.Unmarshal(scanner.Bytes(), &logged); err != nil {
		t.Fatalf("unmarshal attempt failed: %v", err)
	}

	if logged.CardID != "git-fetch-basic" {
		t.Fatalf("unexpected logged card id: %s", logged.CardID)
	}
}

func TestRecordAndPersistWritesProgressAndAttemptsFiles(t *testing.T) {
	dir := t.TempDir()
	progressPath := filepath.Join(dir, "progress.json")
	attemptsPath := filepath.Join(dir, "attempts.jsonl")

	shownAt := time.Date(2026, 4, 5, 9, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	answeredAt := shownAt.Add(9 * time.Second)

	_, state, err := RecordAndPersist(progressPath, attemptsPath, RecordAttemptInput{
		CardID:         "git-cherry-pick-purpose",
		SessionType:    "learn",
		SelectedAnswer: true,
		IsCorrect:      true,
		ShownAt:        shownAt,
		AnsweredAt:     answeredAt,
	})
	if err != nil {
		t.Fatalf("record and persist failed: %v", err)
	}

	card := state.Cards["git-cherry-pick-purpose"]
	if card.CorrectCount != 1 {
		t.Fatalf("expected correct_count 1, got %d", card.CorrectCount)
	}

	if _, err := os.Stat(progressPath); err != nil {
		t.Fatalf("expected progress file to exist: %v", err)
	}

	if _, err := os.Stat(attemptsPath); err != nil {
		t.Fatalf("expected attempts file to exist: %v", err)
	}
}
