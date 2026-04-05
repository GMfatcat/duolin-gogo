package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/dashboard"
	"duolin-gogo/internal/diagnostics"
	"duolin-gogo/internal/notifications"
	"duolin-gogo/internal/pet"
	"duolin-gogo/internal/progress"
	"duolin-gogo/internal/review"
	"duolin-gogo/internal/scheduler"
	"duolin-gogo/internal/selection"
	"duolin-gogo/internal/settings"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx                context.Context
	knowledgeDir       string
	dataDir            string
	nowFunc            func() time.Time
	notificationSender notifications.Sender
	schedulerState     scheduler.State
	tickerFactory      func(time.Duration) *time.Ticker
	lastReviewRunAt    *time.Time
	allowQuit          bool
	quitRuntime        func()
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

type NotificationSettings struct {
	Style     string `json:"style"`
	TitleMode string `json:"titleMode"`
}

type ScheduleSettings struct {
	NotificationIntervalMinutes int    `json:"notificationIntervalMinutes"`
	ReviewTime                  string `json:"reviewTime"`
	ActiveHoursEnabled          bool   `json:"activeHoursEnabled"`
	ActiveHoursStart            string `json:"activeHoursStart"`
	ActiveHoursEnd              string `json:"activeHoursEnd"`
	RevealSpeed                 string `json:"revealSpeed"`
}

type AuthoringPreviewFile struct {
	Path       string `json:"path"`
	Name       string `json:"name"`
	ModifiedAt string `json:"modifiedAt"`
	CardID     string `json:"cardId"`
	TitleZH    string `json:"titleZh"`
	TitleEN    string `json:"titleEn"`
	Topic      string `json:"topic"`
	SearchText string `json:"searchText"`
}

type AuthoringPreviewData struct {
	Files        []AuthoringPreviewFile `json:"files"`
	SelectedPath string                 `json:"selectedPath"`
	CurrentCard  *StudyCard             `json:"currentCard"`
	ImportErrors []diagnostics.Error    `json:"importErrors"`
}

type DraftReviewData struct {
	Items        []DraftReviewItem     `json:"items"`
	CurrentCard  *StudyCard          `json:"currentCard"`
	ImportErrors []diagnostics.Error `json:"importErrors"`
}

type DraftReviewItem struct {
	Index        int                 `json:"index"`
	CurrentCard  *StudyCard          `json:"currentCard"`
	ImportErrors []diagnostics.Error `json:"importErrors"`
	Valid        bool                `json:"valid"`
}

type SaveDraftStatus struct {
	Message    string             `json:"message"`
	SavedPath  string             `json:"savedPath"`
	Topic      string             `json:"topic"`
	Successful bool               `json:"successful"`
	Report     *BatchImportReport `json:"report,omitempty"`
}

type BatchImportReport struct {
	SavedCount   int               `json:"savedCount"`
	SkippedCount int               `json:"skippedCount"`
	WarningCount int               `json:"warningCount"`
	ErrorCount   int               `json:"errorCount"`
	Items        []BatchImportItem `json:"items"`
}

type BatchImportItem struct {
	Index        int    `json:"index"`
	Status       string `json:"status"`
	CardID       string `json:"cardId"`
	SavedPath    string `json:"savedPath"`
	WarningCount int    `json:"warningCount"`
	ErrorCount   int    `json:"errorCount"`
}

type AuthoringPromptData struct {
	Content string `json:"content"`
}

type DraftScaffoldData struct {
	Raw string `json:"raw"`
}

type AnswerChoice struct {
	Value   string `json:"value"`
	LabelZH string `json:"labelZh"`
	LabelEN string `json:"labelEn"`
}

type StudyCard struct {
	ID             string         `json:"id"`
	Title          string         `json:"title"`
	TitleZH        string         `json:"titleZh"`
	TitleEN        string         `json:"titleEn"`
	QuestionType   string         `json:"questionType"`
	QuestionText   string         `json:"questionText"`
	QuestionTextZH string         `json:"questionTextZh"`
	QuestionTextEN string         `json:"questionTextEn"`
	Choices        []AnswerChoice `json:"choices"`
	Clickbait      string         `json:"clickbait"`
	ClickbaitZH    string         `json:"clickbaitZh"`
	ClickbaitEN    string         `json:"clickbaitEn"`
	ReviewHint     string         `json:"reviewHint"`
	ReviewHintZH   string         `json:"reviewHintZh"`
	ReviewHintEN   string         `json:"reviewHintEn"`
	ExplanationZH  string         `json:"explanationZh"`
	ExplanationEN  string         `json:"explanationEn"`
	ShownAt        string         `json:"shownAt"`
}

type DashboardData struct {
	Info                 AppInfo              `json:"info"`
	PreferredLanguage    string               `json:"preferredLanguage"`
	SelectedTopic        string               `json:"selectedTopic"`
	AvailableTopics      []string             `json:"availableTopics"`
	Stats                DashboardStats       `json:"stats"`
	Summary              dashboard.Summary    `json:"summary"`
	ImportErrors         []diagnostics.Error  `json:"importErrors"`
	NotificationSettings NotificationSettings `json:"notificationSettings"`
	ScheduleSettings     ScheduleSettings     `json:"scheduleSettings"`
	CurrentCard          *StudyCard           `json:"currentCard"`
	ReviewQueue          []StudyCard          `json:"reviewQueue"`
	ReviewMode           bool                 `json:"reviewMode"`
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

type ActionStatus struct {
	Message string `json:"message"`
}

type ValidationStatus struct {
	Message      string              `json:"message"`
	ImportErrors []diagnostics.Error `json:"importErrors"`
}

type LearnBreakStatus struct {
	Message   string `json:"message"`
	UnlockAt  string `json:"unlockAt"`
	DurationM int    `json:"durationMinutes"`
}

type DGInteractionStatus struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Variant string `json:"variant"`
	Pose    string `json:"pose"`
	Stage   int    `json:"stage"`
}

type diagnosticSuggestion struct {
	ZH string
	EN string
}

// NewApp creates a new App application struct
func NewApp() *App {
	knowledgeDir, dataDir := defaultAppPaths(os.Executable, os.Getwd)
	return NewAppWithPaths(knowledgeDir, dataDir)
}

func NewAppWithPaths(knowledgeDir, dataDir string) *App {
	return &App{
		knowledgeDir:       knowledgeDir,
		dataDir:            dataDir,
		nowFunc:            time.Now,
		notificationSender: notifications.WindowsSender{},
		tickerFactory:      time.NewTicker,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	_ = notifications.ConfigureApp()
	notifications.RegisterActivationHandler(func(cardID string) {
		if a.ctx != nil {
			a.revealWindow()
			runtime.EventsEmit(a.ctx, "notification:open-card", cardID)
		}
	})
	go a.startNotificationLoop()
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	if a.allowQuit {
		return false
	}

	if ctx != nil {
		runtime.WindowHide(ctx)
	} else if a.ctx != nil {
		runtime.WindowHide(a.ctx)
	}

	return true
}

func (a *App) requestQuit() {
	a.allowQuit = true
}

func (a *App) ExitApplication() {
	a.requestQuit()

	if a.quitRuntime != nil {
		a.quitRuntime()
		return
	}

	if a.ctx != nil {
		runtime.Quit(a.ctx)
	}
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
	diagnosticFile, err := diagnostics.Load(filepath.Join(a.dataDir, "import-errors.json"))
	if err != nil {
		return DashboardData{}, err
	}

	config := a.mustLoadSettings()
	selectedTopic := normalizeSelectedTopic(config.SelectedTopic, cache.Cards)
	filteredCards := selection.FilterCardsByTopic(cache.Cards, selectedTopic)
	now := a.nowFunc()
	reviewQueue := a.buildReviewQueue(filteredCards, state, now)
	reviewMode := review.ShouldStartReview(config, a.lastReviewRunAt, now) && len(reviewQueue) > 0

	card, ok := selection.SelectNextCardForTopic(cache.Cards, state.Cards, selectedTopic, now)
	var currentCard *StudyCard
	if reviewMode {
		first := reviewQueue[0]
		currentCard = &first
	} else if ok {
		currentCard = studyCardFromCard(card, now)
	}

	return DashboardData{
		Info:              a.AppInfo(),
		PreferredLanguage: preferredLanguage,
		SelectedTopic:     selectedTopic,
		AvailableTopics:   availableTopics(cache.Cards),
		Stats:             calculateStats(state, now),
		Summary:           dashboard.BuildSummary(filteredCards, state, now),
		ImportErrors:      enrichDiagnosticsErrors(diagnosticFile.Errors),
		NotificationSettings: NotificationSettings{
			Style:     normalizeNotificationStyle(config.Notifications.Style),
			TitleMode: normalizeNotificationTitleMode(config.Notifications.TitleMode),
		},
		ScheduleSettings: ScheduleSettings{
			NotificationIntervalMinutes: normalizeNotificationInterval(config.NotificationIntervalMinutes),
			ReviewTime:                  normalizeReviewTime(config.ReviewSchedule.Time),
			ActiveHoursEnabled:          config.ActiveHours.Enabled,
			ActiveHoursStart:            normalizeClockTime(config.ActiveHours.Start, "09:00"),
			ActiveHoursEnd:              normalizeClockTime(config.ActiveHours.End, "22:00"),
			RevealSpeed:                 normalizeRevealSpeed(config.StudyRules.RevealSpeed),
		},
		CurrentCard: currentCard,
		ReviewQueue: reviewQueue,
		ReviewMode:  reviewMode,
	}, nil
}

func (a *App) GetStudyCard(cardID string) (*StudyCard, error) {
	cache, _, _, err := a.loadState()
	if err != nil {
		return nil, err
	}

	card, err := findCard(cache.Cards, cardID)
	if err != nil {
		return nil, err
	}

	return studyCardFromCard(card, a.nowFunc()), nil
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

	petEvent := pet.StudyEventAnsweredWrong
	if isCorrect {
		petEvent = pet.StudyEventAnsweredCorrect
	}
	if _, err := pet.RecordStudyEvent(filepath.Join(a.dataDir, "pet.json"), petEvent, now); err != nil {
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

func (a *App) CheckAndSendNotification() (bool, error) {
	config, err := settings.Load(filepath.Join(a.dataDir, "settings.json"))
	if err != nil {
		return false, err
	}

	now := a.nowFunc()
	if !scheduler.ShouldSendLearningNotification(config, a.schedulerState, now) {
		return false, nil
	}

	cache, state, preferredLanguage, err := a.loadState()
	if err != nil {
		return false, err
	}
	selectedTopic := normalizeSelectedTopic(config.SelectedTopic, cache.Cards)

	if review.ShouldStartReview(config, a.lastReviewRunAt, now) {
		queue := review.BuildQueue(selection.FilterCardsByTopic(cache.Cards, selectedTopic), state.Cards, now, config.ReviewSchedule.BatchSize)
		if len(queue) > 0 {
			if err := a.notificationSender.Send(buildReviewNotificationMessage(preferredLanguage, selectedTopic)); err != nil {
				return false, err
			}
			a.lastReviewRunAt = &now
			a.schedulerState.LastNotificationAt = &now
			return true, nil
		}
	}

	card, ok := selection.SelectNextCardForTopic(cache.Cards, state.Cards, selectedTopic, now)
	if !ok {
		return false, nil
	}

	if err := a.notificationSender.Send(notifications.BuildStudyMessageForLanguage(card, preferredLanguage, config.Notifications.Style, config.Notifications.TitleMode)); err != nil {
		return false, err
	}

	a.schedulerState.LastNotificationAt = &now
	return true, nil
}
func (a *App) SnoozeNotifications() (ActionStatus, error) {
	config, err := settings.Load(filepath.Join(a.dataDir, "settings.json"))
	if err != nil {
		return ActionStatus{}, err
	}

	now := a.nowFunc()
	snoozeMinutes := config.StudyRules.SnoozeMinutes
	if snoozeMinutes <= 0 {
		snoozeMinutes = 15
	}

	snoozedUntil := now.Add(time.Duration(snoozeMinutes) * time.Minute)
	a.schedulerState.SnoozedUntil = &snoozedUntil
	return ActionStatus{
		Message: fmt.Sprintf("Notifications snoozed until %s.", snoozedUntil.Format("15:04")),
	}, nil
}

func (a *App) SendTestNotification() (ActionStatus, error) {
	cache, state, preferredLanguage, err := a.loadState()
	if err != nil {
		return ActionStatus{}, err
	}
	selectedTopic := normalizeSelectedTopic(a.mustLoadSettings().SelectedTopic, cache.Cards)

	card, ok := selection.SelectNextCardForTopic(cache.Cards, state.Cards, selectedTopic, a.nowFunc())
	if !ok {
		filtered := selection.FilterCardsByTopic(cache.Cards, selectedTopic)
		if len(filtered) > 0 {
			card = filtered[0]
			ok = true
		}
	}
	if !ok {
		return ActionStatus{Message: "No card available for test notification."}, nil
	}

	config := a.mustLoadSettings()
	message := notifications.BuildStudyMessageForLanguage(card, preferredLanguage, config.Notifications.Style, config.Notifications.TitleMode)
	if err := a.notificationSender.Send(message); err != nil {
		return ActionStatus{}, err
	}

	if selectedTopic == "all" {
		return ActionStatus{Message: "Test notification sent."}, nil
	}

	return ActionStatus{Message: fmt.Sprintf("Test notification sent for %s.", topicScopeLabel(preferredLanguage, selectedTopic))}, nil
}

func (a *App) StartLearnBreak() (LearnBreakStatus, error) {
	config, err := settings.Load(filepath.Join(a.dataDir, "settings.json"))
	if err != nil {
		return LearnBreakStatus{}, err
	}

	now := a.nowFunc()
	durationMinutes := normalizeNotificationInterval(config.NotificationIntervalMinutes)
	unlockAt := now.Add(time.Duration(durationMinutes) * time.Minute)
	a.schedulerState.SnoozedUntil = &unlockAt

	return LearnBreakStatus{
		Message:   fmt.Sprintf("Learn break started until %s.", unlockAt.Format("15:04")),
		UnlockAt:  unlockAt.Format(time.RFC3339),
		DurationM: durationMinutes,
	}, nil
}

func (a *App) InteractWithDG() (DGInteractionStatus, error) {
	config := a.mustLoadSettings()
	cache, err := cards.LoadCache(filepath.Join(a.dataDir, "cards-cache.gob"))
	if err != nil {
		return DGInteractionStatus{}, err
	}

	selectedTopic := normalizeSelectedTopic(config.SelectedTopic, cache.Cards)
	result, err := pet.Interact(filepath.Join(a.dataDir, "pet.json"), a.loadPreferredLanguage(), selectedTopic, a.nowFunc())
	if err != nil {
		return DGInteractionStatus{}, err
	}

	return DGInteractionStatus{
		Title:   result.Reaction.Title,
		Body:    result.Reaction.Body,
		Variant: result.Reaction.Variant,
		Pose:    result.Reaction.Pose,
		Stage:   result.State.Stage,
	}, nil
}

func (a *App) GetDGReaction(trigger string) (DGInteractionStatus, error) {
	config := a.mustLoadSettings()
	cache, err := cards.LoadCache(filepath.Join(a.dataDir, "cards-cache.gob"))
	if err != nil {
		return DGInteractionStatus{}, err
	}

	selectedTopic := normalizeSelectedTopic(config.SelectedTopic, cache.Cards)
	result, err := pet.ReactionForTrigger(filepath.Join(a.dataDir, "pet.json"), trigger, a.loadPreferredLanguage(), selectedTopic, a.nowFunc())
	if err != nil {
		return DGInteractionStatus{}, err
	}

	return DGInteractionStatus{
		Title:   result.Reaction.Title,
		Body:    result.Reaction.Body,
		Variant: result.Reaction.Variant,
		Pose:    result.Reaction.Pose,
		Stage:   result.State.Stage,
	}, nil
}

func (a *App) RescanKnowledge() (ActionStatus, error) {
	result, err := cards.RefreshKnowledge(a.knowledgeDir, a.dataDir)
	if err != nil {
		return ActionStatus{}, err
	}

	return ActionStatus{
		Message: fmt.Sprintf("Knowledge refreshed: %d cards, %d errors.", len(result.Cards), len(result.Errors)),
	}, nil
}

func (a *App) ValidateKnowledge() (ValidationStatus, error) {
	result, err := cards.ScanDirectories([]string{a.knowledgeDir})
	if err != nil {
		return ValidationStatus{}, err
	}

	if err := cards.WriteImportErrors(filepath.Join(a.dataDir, "import-errors.json"), result.Errors); err != nil {
		return ValidationStatus{}, err
	}

	diagnosticItems := make([]diagnostics.Error, 0, len(result.Errors))
	for _, item := range result.Errors {
		diagnosticItems = append(diagnosticItems, diagnostics.Error{
			SourcePath: item.SourcePath,
			Severity:   item.Severity,
			Code:       item.Code,
			Field:      item.Field,
			Message:    item.Message,
		})
	}

	return ValidationStatus{
		Message:      fmt.Sprintf("Knowledge validated: %d cards, %d diagnostics.", len(result.Cards), len(result.Errors)),
		ImportErrors: enrichDiagnosticsErrors(diagnosticItems),
	}, nil
}

func (a *App) ResetStudyData() (ActionStatus, error) {
	paths := []string{
		filepath.Join(a.dataDir, "progress.json"),
		filepath.Join(a.dataDir, "attempts.jsonl"),
		filepath.Join(a.dataDir, "pet.json"),
	}

	for _, path := range paths {
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			return ActionStatus{}, err
		}
	}

	a.schedulerState.LastNotificationAt = nil
	a.schedulerState.SnoozedUntil = nil
	a.lastReviewRunAt = nil

	return ActionStatus{Message: "Study data reset."}, nil
}

func (a *App) LoadAuthoringPreview() (AuthoringPreviewData, error) {
	files, err := cards.ListMarkdownFiles([]string{a.knowledgeDir})
	if err != nil {
		return AuthoringPreviewData{}, err
	}

	if len(files) == 0 {
		return AuthoringPreviewData{
			Files:        []AuthoringPreviewFile{},
			ImportErrors: []diagnostics.Error{},
		}, nil
	}

	return a.previewKnowledgeCard(files[0], files)
}

func (a *App) PreviewKnowledgeCard(path string) (AuthoringPreviewData, error) {
	files, err := cards.ListMarkdownFiles([]string{a.knowledgeDir})
	if err != nil {
		return AuthoringPreviewData{}, err
	}

	return a.previewKnowledgeCard(path, files)
}

func (a *App) ReviewDraft(raw string) (DraftReviewData, error) {
	drafts := splitDraftBatch(raw)
	if len(drafts) == 0 {
		drafts = []string{strings.TrimSpace(raw)}
	}

	items := make([]DraftReviewItem, 0, len(drafts))
	for index, draft := range drafts {
		result, err := cards.PreviewDraft(fmt.Sprintf("draft://ai-card-%d.md", index+1), draft)
		if err != nil {
			return DraftReviewData{}, err
		}

		var previewCard *StudyCard
		if result.Card != nil {
			previewCard = studyCardFromCard(*result.Card, a.nowFunc())
		}

		items = append(items, DraftReviewItem{
			Index:        index + 1,
			CurrentCard:  previewCard,
			ImportErrors: toDiagnosticsErrors(result.Errors),
			Valid:        result.Card != nil,
		})
	}

	first := DraftReviewItem{}
	if len(items) > 0 {
		first = items[0]
	}

	return DraftReviewData{
		Items:        items,
		CurrentCard:  first.CurrentCard,
		ImportErrors: enrichDiagnosticsErrors(first.ImportErrors),
	}, nil
}

func (a *App) SaveDraft(raw string, topic string) (SaveDraftStatus, error) {
	drafts := splitDraftBatch(raw)
	if len(drafts) == 0 {
		drafts = []string{strings.TrimSpace(raw)}
	}

	topic = normalizeTopic(topic)
	targetDir := filepath.Join(a.knowledgeDir, topic)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return SaveDraftStatus{}, err
	}

	report := &BatchImportReport{
		Items: make([]BatchImportItem, 0, len(drafts)),
	}
	firstSavedPath := ""

	for index, draft := range drafts {
		result, err := cards.PreviewDraft(fmt.Sprintf("draft://ai-card-%d.md", index+1), draft)
		if err != nil {
			return SaveDraftStatus{}, err
		}

		warningCount, errorCount := countImportDiagnostics(result.Errors)
		item := BatchImportItem{
			Index:        index + 1,
			WarningCount: warningCount,
			ErrorCount:   errorCount,
		}
		report.WarningCount += warningCount
		report.ErrorCount += errorCount

		if result.Card == nil {
			item.Status = "skipped"
			report.SkippedCount++
			report.Items = append(report.Items, item)
			continue
		}

		targetPath := filepath.Join(targetDir, fmt.Sprintf("%s.md", result.Card.ID))
		if err := os.WriteFile(targetPath, []byte(draft), 0o644); err != nil {
			return SaveDraftStatus{}, err
		}

		item.Status = "saved"
		item.CardID = result.Card.ID
		item.SavedPath = targetPath
		report.SavedCount++
		report.Items = append(report.Items, item)

		if firstSavedPath == "" {
			firstSavedPath = targetPath
		}
	}

	if len(drafts) == 1 && report.SavedCount == 0 {
		return SaveDraftStatus{}, fmt.Errorf("draft has blocking diagnostics")
	}

	message := "No drafts were saved."
	successful := report.SavedCount > 0
	if len(drafts) == 1 && report.SavedCount == 1 {
		message = fmt.Sprintf("Draft saved to %s.", firstSavedPath)
	} else if report.SavedCount > 0 {
		message = fmt.Sprintf("Saved %d drafts. Skipped %d drafts.", report.SavedCount, report.SkippedCount)
	}

	return SaveDraftStatus{
		Message:    message,
		SavedPath:  firstSavedPath,
		Topic:      topic,
		Successful: successful,
		Report:     report,
	}, nil
}

func (a *App) LoadAuthoringPrompt() (AuthoringPromptData, error) {
	root := filepath.Dir(a.knowledgeDir)
	content, err := os.ReadFile(filepath.Join(root, "AI_CARD_PROMPT.md"))
	if err != nil {
		return AuthoringPromptData{}, err
	}

	return AuthoringPromptData{
		Content: string(content),
	}, nil
}

func (a *App) GenerateDraftScaffold(sourceNotes string, topic string) (DraftScaffoldData, error) {
	topic = normalizeTopic(topic)
	title := inferScaffoldTitle(sourceNotes, topic)
	id := scaffoldID(topic, title)
	scaffold := fmt.Sprintf(`---
id: %s
title_zh: %s
title_en: %s
type: true-false
question_zh: "TODO：把這段筆記改成一個可以判斷對錯的觀念題目。"
question_en: "TODO: turn this note into a true-or-false concept check."
clickbait_zh: "這段筆記真正想你記住的是哪個觀念？"
clickbait_en: "What is the one idea this note is actually trying to teach?"
review_hint_zh: "TODO：補一句短而好記的提示。"
review_hint_en: "TODO: add one short memorable hint."
tags: [%s]
difficulty: 2
answer: false
enabled: true
body_format: bilingual-section
---

## zh-TW

TODO：把下面的原始筆記整理成精簡的繁體中文說明。

原始筆記：

%s

## en

TODO: rewrite the note below into a concise English explanation.

Source note:

%s
`, id, title, title, topic, indentScaffoldBody(sourceNotes), indentScaffoldBody(sourceNotes))

	return DraftScaffoldData{Raw: scaffold}, nil
}

func splitDraftBatch(raw string) []string {
	normalized := strings.ReplaceAll(raw, "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\n<!-- draft-break -->\n", "\n===\n")

	parts := strings.Split(normalized, "\n===\n")
	drafts := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		drafts = append(drafts, trimmed)
	}

	return drafts
}

func inferScaffoldTitle(sourceNotes string, topic string) string {
	lines := strings.Split(strings.ReplaceAll(sourceNotes, "\r\n", "\n"), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(strings.TrimLeft(line, "#-*0123456789. "))
		if trimmed != "" {
			return trimmed
		}
	}

	return strings.Title(topic) + " Note"
}

func scaffoldID(topic string, title string) string {
	normalized := strings.ToLower(title)
	replacer := strings.NewReplacer(" ", "-", "_", "-", "/", "-", "\\", "-", ":", "", ".", "", ",", "", "`", "", "'", "")
	normalized = replacer.Replace(normalized)
	normalized = regexp.MustCompile(`[^a-z0-9-]+`).ReplaceAllString(normalized, "-")
	normalized = regexp.MustCompile(`-+`).ReplaceAllString(normalized, "-")
	normalized = strings.Trim(normalized, "-")
	if normalized == "" {
		normalized = "note-card"
	}
	return fmt.Sprintf("%s-%s", topic, normalized)
}

func indentScaffoldBody(source string) string {
	source = strings.TrimSpace(source)
	if source == "" {
		return "- TODO: paste or summarize your note here."
	}
	lines := strings.Split(strings.ReplaceAll(source, "\r\n", "\n"), "\n")
	for index, line := range lines {
		lines[index] = "> " + strings.TrimRight(line, " ")
	}
	return strings.Join(lines, "\n")
}

func toDiagnosticsErrors(items []cards.ImportError) []diagnostics.Error {
	diagnosticItems := make([]diagnostics.Error, 0, len(items))
	for _, item := range items {
		diagnosticItems = append(diagnosticItems, toDiagnosticsError(item))
	}

	return diagnosticItems
}

func toDiagnosticsError(item cards.ImportError) diagnostics.Error {
	suggestion := suggestionForDiagnostic(item.Code, item.Field)
	return diagnostics.Error{
		SourcePath:   item.SourcePath,
		Severity:     item.Severity,
		Code:         item.Code,
		Field:        item.Field,
		Message:      item.Message,
		SuggestionZH: suggestion.ZH,
		SuggestionEN: suggestion.EN,
	}
}

func enrichDiagnosticsErrors(items []diagnostics.Error) []diagnostics.Error {
	out := make([]diagnostics.Error, 0, len(items))
	for _, item := range items {
		suggestion := suggestionForDiagnostic(item.Code, item.Field)
		if item.SuggestionZH == "" {
			item.SuggestionZH = suggestion.ZH
		}
		if item.SuggestionEN == "" {
			item.SuggestionEN = suggestion.EN
		}
		out = append(out, item)
	}
	return out
}

func suggestionForDiagnostic(code string, field string) diagnosticSuggestion {
	switch code {
	case "missing_language_section":
		return diagnosticSuggestion{
			ZH: "補上完整的 `## zh-TW` 和 `## en` 兩段，而且兩邊都要有內容。",
			EN: "Add both `## zh-TW` and `## en` sections, and make sure neither section is empty.",
		}
	case "missing_localized_field":
		return diagnosticSuggestion{
			ZH: fmt.Sprintf("補上 `%s`，避免只靠 fallback 值撐過匯入。", field),
			EN: fmt.Sprintf("Add `%s` so the card does not rely on fallback content.", field),
		}
	case "bilingual_choice_count_mismatch":
		return diagnosticSuggestion{
			ZH: "讓 `choices_zh` 和 `choices_en` 的選項數量一致，並保持順序對齊。",
			EN: "Make `choices_zh` and `choices_en` the same length and keep the option order aligned.",
		}
	case "missing_required_field":
		return diagnosticSuggestion{
			ZH: fmt.Sprintf("先補齊必要欄位 `%s`，這張卡才能通過基本 schema 驗證。", field),
			EN: fmt.Sprintf("Fill in the required `%s` field before this card can pass schema validation.", field),
		}
	case "missing_choices":
		return diagnosticSuggestion{
			ZH: "單選題至少提供 2 個選項，並讓 `answer` 指向有效索引。",
			EN: "Provide at least 2 choices for a single-choice card and keep `answer` within range.",
		}
	case "invalid_answer_type":
		return diagnosticSuggestion{
			ZH: "檢查 `type` 和 `answer` 是否配對：單選題用整數索引，true-false 用布林值。",
			EN: "Check that `type` matches `answer`: single-choice needs an integer index, true-false needs a boolean.",
		}
	case "answer_out_of_range":
		return diagnosticSuggestion{
			ZH: "把 `answer` 改成有效索引，範圍要落在選項數量內。",
			EN: "Change `answer` to a valid zero-based index that exists in the choice list.",
		}
	case "frontmatter_parse_failed":
		return diagnosticSuggestion{
			ZH: "檢查 YAML frontmatter 的縮排、冒號和清單格式，先讓 frontmatter 能被正常解析。",
			EN: "Check YAML frontmatter indentation, colons, and list formatting so the frontmatter parses cleanly.",
		}
	case "duplicate_id":
		return diagnosticSuggestion{
			ZH: "換一個新的 `id`，避免和既有卡片衝突。",
			EN: "Use a new unique `id` so this card does not collide with an existing one.",
		}
	case "unsupported_type":
		return diagnosticSuggestion{
			ZH: "目前只支援 `single-choice` 和 `true-false`，請先改成其中一種。",
			EN: "Use one of the currently supported types: `single-choice` or `true-false`.",
		}
	default:
		return diagnosticSuggestion{}
	}
}

func countImportDiagnostics(items []cards.ImportError) (warnings int, errors int) {
	for _, item := range items {
		if item.Severity == "warning" {
			warnings++
			continue
		}
		errors++
	}
	return warnings, errors
}

func (a *App) UpdateNotificationSettings(style string, titleMode string) (ActionStatus, error) {
	path := filepath.Join(a.dataDir, "settings.json")
	file, err := settings.Load(path)
	if err != nil {
		return ActionStatus{}, err
	}

	file.Notifications.Style = normalizeNotificationStyle(style)
	file.Notifications.TitleMode = normalizeNotificationTitleMode(titleMode)

	if err := writeSettingsFile(path, file); err != nil {
		return ActionStatus{}, err
	}

	return ActionStatus{Message: "Notification settings updated."}, nil
}

func (a *App) UpdatePreferredLanguage(language string) (ActionStatus, error) {
	path := filepath.Join(a.dataDir, "settings.json")
	file, err := settings.Load(path)
	if err != nil {
		return ActionStatus{}, err
	}

	file.Language.Default = normalizePreferredLanguage(language)

	if err := writeSettingsFile(path, file); err != nil {
		return ActionStatus{}, err
	}

	return ActionStatus{Message: "Language updated."}, nil
}

func (a *App) UpdateSelectedTopic(topic string) (ActionStatus, error) {
	path := filepath.Join(a.dataDir, "settings.json")
	file, err := settings.Load(path)
	if err != nil {
		return ActionStatus{}, err
	}

	cache, err := cards.LoadCache(filepath.Join(a.dataDir, "cards-cache.gob"))
	if err != nil {
		return ActionStatus{}, err
	}

	file.SelectedTopic = normalizeSelectedTopic(topic, cache.Cards)

	if err := writeSettingsFile(path, file); err != nil {
		return ActionStatus{}, err
	}

	return ActionStatus{Message: "Topic filter updated."}, nil
}

func (a *App) UpdateScheduleSettings(notificationIntervalMinutes int, reviewTime string, activeHoursEnabled bool, activeHoursStart string, activeHoursEnd string, revealSpeed string) (ActionStatus, error) {
	path := filepath.Join(a.dataDir, "settings.json")
	file, err := settings.Load(path)
	if err != nil {
		return ActionStatus{}, err
	}

	file.NotificationIntervalMinutes = normalizeNotificationInterval(notificationIntervalMinutes)
	file.ReviewSchedule.Time = normalizeReviewTime(reviewTime)
	file.ActiveHours.Enabled = activeHoursEnabled
	file.ActiveHours.Start = normalizeClockTime(activeHoursStart, "09:00")
	file.ActiveHours.End = normalizeClockTime(activeHoursEnd, "22:00")
	file.StudyRules.RevealSpeed = normalizeRevealSpeed(revealSpeed)

	if err := writeSettingsFile(path, file); err != nil {
		return ActionStatus{}, err
	}

	now := a.nowFunc()
	a.schedulerState.LastNotificationAt = &now
	a.schedulerState.SnoozedUntil = nil

	return ActionStatus{Message: "Schedule settings updated."}, nil
}

func (a *App) loadState() (cards.CacheFile, progress.ProgressFile, string, error) {
	cache, _, err := cards.EnsureKnowledgeCache(a.knowledgeDir, a.dataDir)
	if err != nil {
		return cards.CacheFile{}, progress.ProgressFile{}, "", err
	}

	state, err := loadProgress(filepath.Join(a.dataDir, "progress.json"))
	if err != nil {
		return cards.CacheFile{}, progress.ProgressFile{}, "", err
	}

	return cache, state, a.loadPreferredLanguage(), nil
}

func (a *App) previewKnowledgeCard(path string, files []string) (AuthoringPreviewData, error) {
	cache, _ := cards.LoadCache(filepath.Join(a.dataDir, "cards-cache.gob"))
	cacheByPath := make(map[string]cards.Card, len(cache.Cards))
	for _, card := range cache.Cards {
		if card.SourcePath == "" {
			continue
		}
		cacheByPath[filepath.Clean(card.SourcePath)] = card
	}

	previewFiles := make([]AuthoringPreviewFile, 0, len(files))
	for _, file := range files {
		modifiedAt := ""
		if info, err := os.Stat(file); err == nil {
			modifiedAt = info.ModTime().Format(time.RFC3339)
		}
		cleanPath := filepath.Clean(file)
		cachedCard := cacheByPath[cleanPath]
		topic := authoringTopicFromPath(cleanPath)
		searchParts := []string{
			filepath.Base(cleanPath),
			cleanPath,
			cachedCard.ID,
			cachedCard.TitleZH,
			cachedCard.TitleEN,
			topic,
		}
		previewFiles = append(previewFiles, AuthoringPreviewFile{
			Path:       file,
			Name:       filepath.Base(file),
			ModifiedAt: modifiedAt,
			CardID:     cachedCard.ID,
			TitleZH:    cachedCard.TitleZH,
			TitleEN:    cachedCard.TitleEN,
			Topic:      topic,
			SearchText: strings.ToLower(strings.Join(searchParts, "\n")),
		})
	}

	if len(files) == 0 {
		return AuthoringPreviewData{
			Files:        previewFiles,
			ImportErrors: []diagnostics.Error{},
		}, nil
	}

	selectedPath := path
	if selectedPath == "" {
		selectedPath = files[0]
	}

	found := false
	for _, file := range files {
		if file == selectedPath {
			found = true
			break
		}
	}
	if !found {
		selectedPath = files[0]
	}

	result, err := cards.PreviewFile(selectedPath)
	if err != nil {
		return AuthoringPreviewData{}, err
	}

	diagnosticItems := make([]diagnostics.Error, 0, len(result.Errors))
	for _, item := range result.Errors {
		diagnosticItems = append(diagnosticItems, diagnostics.Error{
			SourcePath: item.SourcePath,
			Severity:   item.Severity,
			Code:       item.Code,
			Field:      item.Field,
			Message:    item.Message,
		})
	}

	var previewCard *StudyCard
	if result.Card != nil {
		previewCard = studyCardFromCard(*result.Card, a.nowFunc())
	}

	return AuthoringPreviewData{
		Files:        previewFiles,
		SelectedPath: selectedPath,
		CurrentCard:  previewCard,
		ImportErrors: enrichDiagnosticsErrors(diagnosticItems),
	}, nil
}

func authoringTopicFromPath(path string) string {
	normalized := filepath.ToSlash(path)
	match := regexp.MustCompile(`/knowledge/([^/]+)/`).FindStringSubmatch(normalized)
	if len(match) > 1 {
		return match[1]
	}
	return ""
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
	file, err := settings.Load(filepath.Join(a.dataDir, "settings.json"))
	if err != nil {
		return a.AppInfo().DefaultLanguage
	}

	return normalizePreferredLanguage(file.Language.Default)
}

func normalizeNotificationStyle(style string) string {
	switch style {
	case "safe", "playful", "aggressive", "chaotic":
		return style
	default:
		return "playful"
	}
}

func normalizeNotificationTitleMode(titleMode string) string {
	switch titleMode {
	case "prefer_generated", "prefer_manual":
		return titleMode
	default:
		return "prefer_manual"
	}
}

func normalizePreferredLanguage(language string) string {
	switch language {
	case "en":
		return "en"
	default:
		return "zh-TW"
	}
}

func normalizeSelectedTopic(topic string, allCards []cards.Card) string {
	normalized := strings.ToLower(strings.TrimSpace(topic))
	if normalized == "" || normalized == "all" {
		return "all"
	}

	if isTopicPreset(normalized) {
		return normalized
	}

	for _, available := range availableTopics(allCards) {
		if available == normalized {
			return normalized
		}
	}

	return "all"
}

func topicScopeLabel(language string, topic string) string {
	switch normalizeTopicKey(topic) {
	case "all":
		if language == "zh-TW" {
			return "混合模式"
		}
		return "Mixed mode"
	case "backend-tools":
		if language == "zh-TW" {
			return "後端工具"
		}
		return "Backend tools"
	case "languages":
		if language == "zh-TW" {
			return "程式語言"
		}
		return "Languages"
	case "git":
		return "Git"
	case "docker":
		return "Docker"
	case "linux":
		return "Linux"
	case "go":
		return "Go"
	case "python":
		return "Python"
	default:
		return topic
	}
}

func buildReviewNotificationMessage(language string, topic string) notifications.Message {
	return notifications.Message{
		Title: reviewNotificationTitle(language),
		Body:  reviewNotificationBody(language, topic),
	}
}

func reviewNotificationTitle(language string) string {
	if language == "zh-TW" {
		return "複習時間到了"
	}
	return "Review time"
}

func reviewNotificationBody(language string, topic string) string {
	switch normalizeTopicKey(topic) {
	case "backend-tools":
		if language == "zh-TW" {
			return "回頭看一下最近學過的後端工具觀念。"
		}
		return "Time to review your recent backend tools concepts."
	case "languages":
		if language == "zh-TW" {
			return "回頭看一下最近學過的程式語言觀念。"
		}
		return "Time to review your recent language concepts."
	case "all":
		if language == "zh-TW" {
			return "回頭看一下最近學過的主題。"
		}
		return "Time to review the concepts you studied recently."
	default:
		scopeLabel := topicScopeLabel(language, topic)
		if language == "zh-TW" {
			return fmt.Sprintf("回頭看一下最近學過的 %s 觀念。", scopeLabel)
		}
		return fmt.Sprintf("Time to review your recent %s concepts.", scopeLabel)
	}
}

func normalizeTopicKey(topic string) string {
	normalized := strings.ToLower(strings.TrimSpace(topic))
	if normalized == "" {
		return "all"
	}
	return normalized
}

func normalizeTopic(topic string) string {
	topic = strings.ReplaceAll(topic, "\\", " ")
	topic = strings.ReplaceAll(topic, "/", " ")
	topic = strings.TrimSpace(strings.ToLower(topic))
	topic = regexp.MustCompile(`[^a-z0-9._-]+`).ReplaceAllString(topic, "-")
	topic = strings.Trim(topic, "-")
	if topic == "." || topic == "" {
		return "git"
	}
	return topic
}

func availableTopics(allCards []cards.Card) []string {
	seen := map[string]struct{}{}
	topics := []string{"all", "backend-tools", "languages"}

	for _, card := range allCards {
		topic := topicForCard(card)
		if topic == "" || topic == "all" {
			continue
		}
		if _, ok := seen[topic]; ok {
			continue
		}
		seen[topic] = struct{}{}
		topics = append(topics, topic)
	}

	if len(topics) > 1 {
		slices.Sort(topics[1:])
	}

	return topics
}

func isTopicPreset(topic string) bool {
	switch topic {
	case "backend-tools", "languages":
		return true
	default:
		return false
	}
}

func topicForCard(card cards.Card) string {
	if card.SourcePath != "" {
		parent := strings.ToLower(strings.TrimSpace(filepath.Base(filepath.Dir(card.SourcePath))))
		if parent != "" && parent != "." {
			return parent
		}
	}

	for _, tag := range card.Tags {
		tag = strings.ToLower(strings.TrimSpace(tag))
		if tag != "" && tag != "all" {
			return tag
		}
	}

	return ""
}

func normalizeNotificationInterval(minutes int) int {
	if minutes <= 0 {
		return 20
	}
	if minutes < 5 {
		return 5
	}
	if minutes > 120 {
		return 120
	}
	return minutes
}

func normalizeReviewTime(value string) string {
	return normalizeClockTime(value, "21:00")
}

func normalizeClockTime(value string, fallback string) string {
	if len(value) != 5 || value[2] != ':' {
		return fallback
	}

	hour := value[:2]
	minute := value[3:]
	if hour < "00" || hour > "23" || minute < "00" || minute > "59" {
		return fallback
	}

	return value
}

func normalizeRevealSpeed(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "fast", "slow":
		return strings.ToLower(strings.TrimSpace(value))
	default:
		return "normal"
	}
}

func writeSettingsFile(path string, file settings.File) error {
	bytes, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(bytes, '\n'), 0o644)
}

func (a *App) mustLoadSettings() settings.File {
	file, err := settings.Load(filepath.Join(a.dataDir, "settings.json"))
	if err != nil {
		return settings.File{}
	}
	return file
}

func (a *App) buildReviewQueue(allCards []cards.Card, state progress.ProgressFile, now time.Time) []StudyCard {
	config := a.mustLoadSettings()
	queue := review.BuildQueue(allCards, state.Cards, now, config.ReviewSchedule.BatchSize)
	out := make([]StudyCard, 0, len(queue))
	for _, card := range queue {
		out = append(out, *studyCardFromCard(card, now))
	}
	return out
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
		ID:             card.ID,
		Title:          card.TitleEN,
		TitleZH:        card.TitleZH,
		TitleEN:        card.TitleEN,
		QuestionType:   card.QuestionType,
		QuestionText:   card.QuestionTextEN,
		QuestionTextZH: card.QuestionTextZH,
		QuestionTextEN: card.QuestionTextEN,
		Choices:        choicesFromCard(card),
		Clickbait:      card.ClickbaitEN,
		ClickbaitZH:    card.ClickbaitZH,
		ClickbaitEN:    card.ClickbaitEN,
		ReviewHint:     card.ReviewHintEN,
		ReviewHintZH:   card.ReviewHintZH,
		ReviewHintEN:   card.ReviewHintEN,
		ExplanationZH:  card.BodyMarkdownZH,
		ExplanationEN:  card.BodyMarkdownEN,
		ShownAt:        now.Format(time.RFC3339),
	}
}

func choicesFromCard(card cards.Card) []AnswerChoice {
	if card.QuestionType == "true-false" {
		return []AnswerChoice{
			{Value: "true", LabelZH: "是", LabelEN: "True"},
			{Value: "false", LabelZH: "否", LabelEN: "False"},
		}
	}

	choices := make([]AnswerChoice, 0, len(card.ChoicesEN))
	for index, choice := range card.ChoicesEN {
		labelZH := choice
		if index < len(card.ChoicesZH) {
			labelZH = card.ChoicesZH[index]
		}
		choices = append(choices, AnswerChoice{
			Value:   fmt.Sprintf("%d", index),
			LabelZH: labelZH,
			LabelEN: choice,
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

func (a *App) startNotificationLoop() {
	if a.tickerFactory == nil {
		return
	}

	ticker := a.tickerFactory(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		_, _ = a.CheckAndSendNotification()
	}
}

func (a *App) revealWindow() {
	if a.ctx == nil {
		return
	}

	runtime.WindowShow(a.ctx)
	runtime.WindowUnminimise(a.ctx)
	runtime.WindowSetAlwaysOnTop(a.ctx, true)
	time.Sleep(150 * time.Millisecond)
	runtime.WindowSetAlwaysOnTop(a.ctx, false)
}

func defaultAppPaths(executablePathFunc func() (string, error), workingDirFunc func() (string, error)) (string, string) {
	const (
		knowledgeName = "knowledge"
		dataName      = "data"
	)

	fallbackKnowledge := filepath.Clean(filepath.Join("..", knowledgeName))
	fallbackData := filepath.Clean(filepath.Join("..", dataName))

	searchRoots := []string{}

	if executablePath, err := executablePathFunc(); err == nil && executablePath != "" {
		searchRoots = append(searchRoots, filepath.Dir(executablePath))
	}

	if workingDir, err := workingDirFunc(); err == nil && workingDir != "" {
		searchRoots = append(searchRoots, workingDir)
	}

	for _, root := range searchRoots {
		if knowledgeDir, dataDir, ok := findResourceDirs(root, knowledgeName, dataName); ok {
			return knowledgeDir, dataDir
		}
	}

	return fallbackKnowledge, fallbackData
}

func findResourceDirs(startDir string, knowledgeName string, dataName string) (string, string, bool) {
	current := filepath.Clean(startDir)

	for {
		knowledgeDir := filepath.Join(current, knowledgeName)
		dataDir := filepath.Join(current, dataName)

		if isDirectory(knowledgeDir) && isDirectory(dataDir) {
			return knowledgeDir, dataDir, true
		}

		parent := filepath.Dir(current)
		if parent == current {
			return "", "", false
		}
		current = parent
	}
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
