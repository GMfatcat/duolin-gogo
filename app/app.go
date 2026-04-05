package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/dashboard"
	"duolin-gogo/internal/diagnostics"
	"duolin-gogo/internal/notifications"
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
}

type AuthoringPreviewFile struct {
	Path       string `json:"path"`
	Name       string `json:"name"`
	ModifiedAt string `json:"modifiedAt"`
}

type AuthoringPreviewData struct {
	Files        []AuthoringPreviewFile `json:"files"`
	SelectedPath string                 `json:"selectedPath"`
	CurrentCard  *StudyCard             `json:"currentCard"`
	ImportErrors []diagnostics.Error    `json:"importErrors"`
}

type DraftReviewData struct {
	CurrentCard  *StudyCard          `json:"currentCard"`
	ImportErrors []diagnostics.Error `json:"importErrors"`
}

type SaveDraftStatus struct {
	Message    string `json:"message"`
	SavedPath  string `json:"savedPath"`
	Topic      string `json:"topic"`
	Successful bool   `json:"successful"`
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
		ImportErrors:      diagnosticFile.Errors,
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
		ImportErrors: diagnosticItems,
	}, nil
}

func (a *App) ResetStudyData() (ActionStatus, error) {
	paths := []string{
		filepath.Join(a.dataDir, "progress.json"),
		filepath.Join(a.dataDir, "attempts.jsonl"),
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
	result, err := cards.PreviewDraft("draft://ai-card.md", raw)
	if err != nil {
		return DraftReviewData{}, err
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

	return DraftReviewData{
		CurrentCard:  previewCard,
		ImportErrors: diagnosticItems,
	}, nil
}

func (a *App) SaveDraft(raw string, topic string) (SaveDraftStatus, error) {
	result, err := cards.PreviewDraft("draft://ai-card.md", raw)
	if err != nil {
		return SaveDraftStatus{}, err
	}

	if result.Card == nil {
		return SaveDraftStatus{}, fmt.Errorf("draft has blocking diagnostics")
	}

	topic = normalizeTopic(topic)
	targetDir := filepath.Join(a.knowledgeDir, topic)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return SaveDraftStatus{}, err
	}

	targetPath := filepath.Join(targetDir, fmt.Sprintf("%s.md", result.Card.ID))
	if err := os.WriteFile(targetPath, []byte(raw), 0o644); err != nil {
		return SaveDraftStatus{}, err
	}

	return SaveDraftStatus{
		Message:    fmt.Sprintf("Draft saved to %s.", targetPath),
		SavedPath:  targetPath,
		Topic:      topic,
		Successful: true,
	}, nil
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

	cache, err := loadCache(filepath.Join(a.dataDir, "cards-cache.json"))
	if err != nil {
		return ActionStatus{}, err
	}

	file.SelectedTopic = normalizeSelectedTopic(topic, cache.Cards)

	if err := writeSettingsFile(path, file); err != nil {
		return ActionStatus{}, err
	}

	return ActionStatus{Message: "Topic filter updated."}, nil
}

func (a *App) UpdateScheduleSettings(notificationIntervalMinutes int, reviewTime string, activeHoursEnabled bool, activeHoursStart string, activeHoursEnd string) (ActionStatus, error) {
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

	if err := writeSettingsFile(path, file); err != nil {
		return ActionStatus{}, err
	}

	now := a.nowFunc()
	a.schedulerState.LastNotificationAt = &now
	a.schedulerState.SnoozedUntil = nil

	return ActionStatus{Message: "Schedule settings updated."}, nil
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

func (a *App) previewKnowledgeCard(path string, files []string) (AuthoringPreviewData, error) {
	previewFiles := make([]AuthoringPreviewFile, 0, len(files))
	for _, file := range files {
		modifiedAt := ""
		if info, err := os.Stat(file); err == nil {
			modifiedAt = info.ModTime().Format(time.RFC3339)
		}
		previewFiles = append(previewFiles, AuthoringPreviewFile{
			Path:       file,
			Name:       filepath.Base(file),
			ModifiedAt: modifiedAt,
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
		ImportErrors: diagnosticItems,
	}, nil
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
	topic = filepath.Base(topic)
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
