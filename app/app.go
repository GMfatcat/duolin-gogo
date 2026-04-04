package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/progress"
	"duolin-gogo/internal/selection"
)

// App struct
type App struct {
	ctx          context.Context
	knowledgeDir string
	dataDir      string
	nowFunc      func() time.Time
}

type AppInfo struct {
	Name            string `json:"name"`
	FocusTopic      string `json:"focusTopic"`
	DefaultLanguage string `json:"defaultLanguage"`
}

type DashboardStats struct {
	StudiedToday int     `json:"studiedToday"`
	CorrectRate  float64 `json:"correctRate"`
}

type AnswerChoice struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type StudyCard struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	QuestionType  string         `json:"questionType"`
	QuestionText  string         `json:"questionText"`
	Choices       []AnswerChoice `json:"choices"`
	Clickbait     string         `json:"clickbait"`
	ReviewHint    string         `json:"reviewHint"`
	ExplanationZH string         `json:"explanationZh"`
	ExplanationEN string         `json:"explanationEn"`
	ShownAt       string         `json:"shownAt"`
}

type DashboardData struct {
	Info              AppInfo        `json:"info"`
	PreferredLanguage string         `json:"preferredLanguage"`
	Stats             DashboardStats `json:"stats"`
	CurrentCard       *StudyCard     `json:"currentCard"`
}

type SubmitAnswerResult struct {
	CardID            string         `json:"cardId"`
	IsCorrect         bool           `json:"isCorrect"`
	CorrectAnswer     string         `json:"correctAnswer"`
	Feedback          string         `json:"feedback"`
	ReviewHint        string         `json:"reviewHint"`
	PreferredLanguage string         `json:"preferredLanguage"`
	Stats             DashboardStats `json:"stats"`
}

type settingsFile struct {
	Language struct {
		Default string `json:"default"`
	} `json:"language"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return NewAppWithPaths(filepath.Clean(filepath.Join("..", "knowledge")), filepath.Clean(filepath.Join("..", "data")))
}

func NewAppWithPaths(knowledgeDir, dataDir string) *App {
	return &App{
		knowledgeDir: knowledgeDir,
		dataDir:      dataDir,
		nowFunc:      time.Now,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// AppInfo returns basic app metadata for the UI shell.
func (a *App) AppInfo() AppInfo {
	return AppInfo{
		Name:            "duolin-gogo",
		FocusTopic:      "git",
		DefaultLanguage: "zh-TW",
	}
}

func (a *App) LoadDashboard() (DashboardData, error) {
	cache, state, preferredLanguage, err := a.loadState()
	if err != nil {
		return DashboardData{}, err
	}

	card, ok := selection.SelectNextCard(cache.Cards, state.Cards, a.nowFunc())
	var currentCard *StudyCard
	if ok {
		currentCard = studyCardFromCard(card, a.nowFunc())
	}

	return DashboardData{
		Info:              a.AppInfo(),
		PreferredLanguage: preferredLanguage,
		Stats:             calculateStats(state, a.nowFunc()),
		CurrentCard:       currentCard,
	}, nil
}

func (a *App) SubmitAnswer(cardID string, sessionType string, selectedAnswer string, shownAt string) (SubmitAnswerResult, error) {
	cache, _, preferredLanguage, err := a.loadState()
	if err != nil {
		return SubmitAnswerResult{}, err
	}

	card, err := findCard(cache.Cards, cardID)
	if err != nil {
		return SubmitAnswerResult{}, err
	}

	shownTime, err := time.Parse(time.RFC3339, shownAt)
	if err != nil {
		return SubmitAnswerResult{}, fmt.Errorf("invalid shownAt: %w", err)
	}

	now := a.nowFunc()
	isCorrect := selectedAnswer == answerValueString(card)
	_, state, err := progress.RecordAndPersist(
		filepath.Join(a.dataDir, "progress.json"),
		filepath.Join(a.dataDir, "attempts.jsonl"),
		progress.RecordAttemptInput{
			CardID:         card.ID,
			SessionType:    sessionType,
			SelectedAnswer: selectedAnswer,
			IsCorrect:      isCorrect,
			ShownAt:        shownTime,
			AnsweredAt:     now,
		},
	)
	if err != nil {
		return SubmitAnswerResult{}, err
	}

	feedback := "Not quite."
	if isCorrect {
		feedback = "Correct."
	}

	return SubmitAnswerResult{
		CardID:            card.ID,
		IsCorrect:         isCorrect,
		CorrectAnswer:     answerValueString(card),
		Feedback:          feedback,
		ReviewHint:        card.ReviewHint,
		PreferredLanguage: preferredLanguage,
		Stats:             calculateStats(state, now),
	}, nil
}

func (a *App) loadState() (cards.CacheFile, progress.ProgressFile, string, error) {
	if _, err := cards.RefreshKnowledge(a.knowledgeDir, a.dataDir); err != nil {
		return cards.CacheFile{}, progress.ProgressFile{}, "", err
	}

	cache, err := loadCache(filepath.Join(a.dataDir, "cards-cache.json"))
	if err != nil {
		return cards.CacheFile{}, progress.ProgressFile{}, "", err
	}

	state, err := loadProgress(filepath.Join(a.dataDir, "progress.json"))
	if err != nil {
		return cards.CacheFile{}, progress.ProgressFile{}, "", err
	}

	return cache, state, a.loadPreferredLanguage(), nil
}

func loadCache(path string) (cards.CacheFile, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return cards.CacheFile{}, err
	}

	var cache cards.CacheFile
	if err := json.Unmarshal(bytes, &cache); err != nil {
		return cards.CacheFile{}, err
	}

	return cache, nil
}

func loadProgress(path string) (progress.ProgressFile, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return progress.ProgressFile{
			Version:      1,
			Cards:        map[string]progress.CardProgress{},
			DailySummary: map[string]progress.DailySummary{},
		}, nil
	}

	return progress.Load(path)
}

func (a *App) loadPreferredLanguage() string {
	path := filepath.Join(a.dataDir, "settings.json")
	bytes, err := os.ReadFile(path)
	if err != nil {
		return a.AppInfo().DefaultLanguage
	}

	var settings settingsFile
	if err := json.Unmarshal(bytes, &settings); err != nil {
		return a.AppInfo().DefaultLanguage
	}

	if settings.Language.Default == "" {
		return a.AppInfo().DefaultLanguage
	}

	return settings.Language.Default
}

func calculateStats(state progress.ProgressFile, now time.Time) DashboardStats {
	day := state.DailySummary[now.Format("2006-01-02")]
	rate := 0.0
	if day.Answered > 0 {
		rate = float64(day.Correct) / float64(day.Answered)
	}

	return DashboardStats{
		StudiedToday: day.Answered,
		CorrectRate:  rate,
	}
}

func studyCardFromCard(card cards.Card, now time.Time) *StudyCard {
	return &StudyCard{
		ID:            card.ID,
		Title:         card.Title,
		QuestionType:  card.QuestionType,
		QuestionText:  card.QuestionText,
		Choices:       choicesFromCard(card),
		Clickbait:     card.Clickbait,
		ReviewHint:    card.ReviewHint,
		ExplanationZH: card.BodyMarkdownZH,
		ExplanationEN: card.BodyMarkdownEN,
		ShownAt:       now.Format(time.RFC3339),
	}
}

func choicesFromCard(card cards.Card) []AnswerChoice {
	if card.QuestionType == "true-false" {
		return []AnswerChoice{
			{Value: "true", Label: "True"},
			{Value: "false", Label: "False"},
		}
	}

	choices := make([]AnswerChoice, 0, len(card.Choices))
	for index, choice := range card.Choices {
		choices = append(choices, AnswerChoice{
			Value: fmt.Sprintf("%d", index),
			Label: choice,
		})
	}

	return choices
}

func answerValueString(card cards.Card) string {
	switch value := card.AnswerValue.(type) {
	case bool:
		if value {
			return "true"
		}
		return "false"
	case int:
		return fmt.Sprintf("%d", value)
	case float64:
		return fmt.Sprintf("%d", int(value))
	default:
		return fmt.Sprint(value)
	}
}

func findCard(all []cards.Card, id string) (cards.Card, error) {
	for _, card := range all {
		if card.ID == id {
			return card, nil
		}
	}

	return cards.Card{}, fmt.Errorf("card %q not found", id)
}
