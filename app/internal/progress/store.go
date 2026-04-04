package progress

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ProgressFile struct {
	Version      int                     `json:"version"`
	UpdatedAt    *string                 `json:"updated_at"`
	Cards        map[string]CardProgress `json:"cards"`
	DailySummary map[string]DailySummary `json:"daily_summary"`
}

type CardProgress struct {
	SeenCount       int     `json:"seen_count"`
	CorrectCount    int     `json:"correct_count"`
	WrongCount      int     `json:"wrong_count"`
	MasteryScore    int     `json:"mastery_score"`
	StreakCorrect   int     `json:"streak_correct"`
	LastSeenAt      *string `json:"last_seen_at"`
	LastCorrectAt   *string `json:"last_correct_at"`
	LastWrongAt     *string `json:"last_wrong_at"`
	LastSessionType string  `json:"last_session_type"`
	IntroducedAt    *string `json:"introduced_at"`
	NextReviewAt    *string `json:"next_review_at"`
	SnoozedUntil    *string `json:"snoozed_until"`
	IsMastered      bool    `json:"is_mastered"`
}

type DailySummary struct {
	CardsSeen      int `json:"cards_seen"`
	Answered       int `json:"answered"`
	Correct        int `json:"correct"`
	Wrong          int `json:"wrong"`
	ReviewAnswered int `json:"review_answered"`
	LearnAnswered  int `json:"learn_answered"`
}

type AttemptEvent struct {
	AttemptID      string `json:"attempt_id"`
	CardID         string `json:"card_id"`
	SessionType    string `json:"session_type"`
	ShownAt        string `json:"shown_at"`
	AnsweredAt     string `json:"answered_at"`
	SelectedAnswer any    `json:"selected_answer"`
	IsCorrect      bool   `json:"is_correct"`
	ResponseTimeMS int64  `json:"response_time_ms"`
	MasteryDelta   int    `json:"mastery_delta"`
}

type RecordAttemptInput struct {
	CardID         string
	SessionType    string
	SelectedAnswer any
	IsCorrect      bool
	ShownAt        time.Time
	AnsweredAt     time.Time
}

type Store struct {
	State ProgressFile
}

func NewStore() *Store {
	return &Store{
		State: ProgressFile{
			Version:      1,
			Cards:        map[string]CardProgress{},
			DailySummary: map[string]DailySummary{},
		},
	}
}

func Load(path string) (ProgressFile, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return ProgressFile{}, err
	}

	var state ProgressFile
	if err := json.Unmarshal(bytes, &state); err != nil {
		return ProgressFile{}, err
	}

	if state.Cards == nil {
		state.Cards = map[string]CardProgress{}
	}
	if state.DailySummary == nil {
		state.DailySummary = map[string]DailySummary{}
	}
	if state.Version == 0 {
		state.Version = 1
	}

	return state, nil
}

func (s *Store) Save(path string) error {
	now := time.Now().Format(time.RFC3339)
	s.State.UpdatedAt = &now

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(s.State, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, append(bytes, '\n'), 0o644)
}

func AppendAttempt(path string, attempt AttemptEvent) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	bytes, err := json.Marshal(attempt)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(append(bytes, '\n')); err != nil {
		return err
	}

	return nil
}

func RecordAndPersist(progressPath, attemptsPath string, input RecordAttemptInput) (AttemptEvent, ProgressFile, error) {
	store := NewStore()

	if _, err := os.Stat(progressPath); err == nil {
		state, loadErr := Load(progressPath)
		if loadErr != nil {
			return AttemptEvent{}, ProgressFile{}, loadErr
		}
		store.State = state
	}

	attempt, err := store.RecordAttempt(input)
	if err != nil {
		return AttemptEvent{}, ProgressFile{}, err
	}

	if err := store.Save(progressPath); err != nil {
		return AttemptEvent{}, ProgressFile{}, err
	}

	if err := AppendAttempt(attemptsPath, attempt); err != nil {
		return AttemptEvent{}, ProgressFile{}, err
	}

	return attempt, store.State, nil
}

func (s *Store) RecordAttempt(input RecordAttemptInput) (AttemptEvent, error) {
	if input.CardID == "" {
		return AttemptEvent{}, errors.New("card id is required")
	}
	if input.SessionType != "learn" && input.SessionType != "review" {
		return AttemptEvent{}, fmt.Errorf("unsupported session type %q", input.SessionType)
	}
	if input.AnsweredAt.Before(input.ShownAt) {
		return AttemptEvent{}, errors.New("answered_at cannot be before shown_at")
	}

	card := s.State.Cards[input.CardID]

	card.SeenCount++
	card.LastSessionType = input.SessionType

	shownAt := input.ShownAt.Format(time.RFC3339)
	answeredAt := input.AnsweredAt.Format(time.RFC3339)

	if card.IntroducedAt == nil {
		card.IntroducedAt = &shownAt
	}

	card.LastSeenAt = &shownAt

	masteryDelta := -2
	if input.IsCorrect {
		card.CorrectCount++
		card.StreakCorrect++
		card.LastCorrectAt = &answeredAt
		masteryDelta = 1
	} else {
		card.WrongCount++
		card.StreakCorrect = 0
		card.LastWrongAt = &answeredAt
	}

	card.MasteryScore += masteryDelta
	if card.MasteryScore < -5 {
		card.MasteryScore = -5
	}
	if card.MasteryScore > 10 {
		card.MasteryScore = 10
	}
	card.IsMastered = card.MasteryScore >= 6

	nextReview := scheduleNextReview(input.AnsweredAt, input.IsCorrect, card.StreakCorrect)
	nextReviewString := nextReview.Format(time.RFC3339)
	card.NextReviewAt = &nextReviewString

	s.State.Cards[input.CardID] = card

	dayKey := input.ShownAt.Format("2006-01-02")
	day := s.State.DailySummary[dayKey]
	day.CardsSeen++
	day.Answered++
	if input.IsCorrect {
		day.Correct++
	} else {
		day.Wrong++
	}
	if input.SessionType == "review" {
		day.ReviewAnswered++
	} else {
		day.LearnAnswered++
	}
	s.State.DailySummary[dayKey] = day

	attempt := AttemptEvent{
		AttemptID:      fmt.Sprintf("att_%s_%d", input.AnsweredAt.Format("20060102T150405"), input.AnsweredAt.UnixNano()),
		CardID:         input.CardID,
		SessionType:    input.SessionType,
		ShownAt:        shownAt,
		AnsweredAt:     answeredAt,
		SelectedAnswer: input.SelectedAnswer,
		IsCorrect:      input.IsCorrect,
		ResponseTimeMS: input.AnsweredAt.Sub(input.ShownAt).Milliseconds(),
		MasteryDelta:   masteryDelta,
	}

	return attempt, nil
}

func scheduleNextReview(answeredAt time.Time, isCorrect bool, streakCorrect int) time.Time {
	if !isCorrect {
		return answeredAt.Add(24 * time.Hour)
	}

	switch streakCorrect {
	case 1:
		return answeredAt.Add(24 * time.Hour)
	case 2:
		return answeredAt.Add(72 * time.Hour)
	default:
		return answeredAt.Add(7 * 24 * time.Hour)
	}
}
