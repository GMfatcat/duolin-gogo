package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/notifications"
	"duolin-gogo/internal/progress"
)

func TestNewAppReturnsApp(t *testing.T) {
	app := NewApp()

	if app == nil {
		t.Fatal("expected app instance")
	}
}

func TestAppInfo(t *testing.T) {
	app := NewApp()

	info := app.AppInfo()

	if info.Name != "duolin-gogo" {
		t.Fatalf("expected app name duolin-gogo, got %q", info.Name)
	}

	if info.FocusTopic != "git" {
		t.Fatalf("expected focus topic git, got %q", info.FocusTopic)
	}

	if info.DefaultLanguage != "zh-TW" {
		t.Fatalf("expected default language zh-TW, got %q", info.DefaultLanguage)
	}
}

func TestBeforeClosePreventsQuitByDefault(t *testing.T) {
	app := NewApp()

	prevent := app.beforeClose(nil)

	if !prevent {
		t.Fatal("expected close to be prevented by default")
	}
}

func TestBeforeCloseAllowsQuitWhenRequested(t *testing.T) {
	app := NewApp()
	app.requestQuit()

	prevent := app.beforeClose(nil)

	if prevent {
		t.Fatal("expected close to be allowed after explicit quit request")
	}
}

func TestExitApplicationRequestsQuitAndInvokesQuitHook(t *testing.T) {
	app := NewApp()
	quitCalled := false
	app.quitRuntime = func() {
		quitCalled = true
	}

	app.ExitApplication()

	if !quitCalled {
		t.Fatal("expected quit hook to be called")
	}

	if !app.allowQuit {
		t.Fatal("expected explicit exit to allow application quit")
	}
}

func TestDefaultAppPathsFindsRepoRootFromExecutableTree(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "knowledge"), 0o755); err != nil {
		t.Fatalf("mkdir knowledge failed: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(root, "data"), 0o755); err != nil {
		t.Fatalf("mkdir data failed: %v", err)
	}

	executablePath := filepath.Join(root, "app", "build", "bin", "app.exe")
	knowledgeDir, dataDir := defaultAppPaths(func() (string, error) {
		return executablePath, nil
	}, func() (string, error) {
		return filepath.Join(root, "app", "build", "bin"), nil
	})

	if knowledgeDir != filepath.Join(root, "knowledge") {
		t.Fatalf("unexpected knowledge dir: %s", knowledgeDir)
	}
	if dataDir != filepath.Join(root, "data") {
		t.Fatalf("unexpected data dir: %s", dataDir)
	}
}

func TestDefaultAppPathsFallsBackToRelativePaths(t *testing.T) {
	knowledgeDir, dataDir := defaultAppPaths(func() (string, error) {
		return "", os.ErrNotExist
	}, func() (string, error) {
		return "", os.ErrNotExist
	})

	if knowledgeDir != filepath.Clean(filepath.Join("..", "knowledge")) {
		t.Fatalf("unexpected fallback knowledge dir: %s", knowledgeDir)
	}
	if dataDir != filepath.Clean(filepath.Join("..", "data")) {
		t.Fatalf("unexpected fallback data dir: %s", dataDir)
	}
}

func TestLoadDashboardReturnsStudyCardAndStats(t *testing.T) {
	app := newTestApp(t)

	_, _, err := progress.RecordAndPersist(
		filepath.Join(app.dataDir, "progress.json"),
		filepath.Join(app.dataDir, "attempts.jsonl"),
		progress.RecordAttemptInput{
			CardID:         "git-rebase-vs-merge",
			SessionType:    "learn",
			SelectedAnswer: "0",
			IsCorrect:      false,
			ShownAt:        app.nowFunc().Add(-2 * time.Hour),
			AnsweredAt:     app.nowFunc().Add(-2*time.Hour + 8*time.Second),
		},
	)
	if err != nil {
		t.Fatalf("seed progress failed: %v", err)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if dashboard.PreferredLanguage != "zh-TW" {
		t.Fatalf("expected preferred language zh-TW, got %s", dashboard.PreferredLanguage)
	}
	if dashboard.SelectedTopic != "all" {
		t.Fatalf("expected selected topic all, got %s", dashboard.SelectedTopic)
	}
	if len(dashboard.AvailableTopics) == 0 || dashboard.AvailableTopics[0] != "all" {
		t.Fatalf("expected available topics to start with all, got %#v", dashboard.AvailableTopics)
	}

	if dashboard.CurrentCard == nil {
		t.Fatal("expected current study card")
	}

	if dashboard.CurrentCard.ID != "git-cherry-pick-purpose" {
		t.Fatalf("expected cherry-pick card, got %s", dashboard.CurrentCard.ID)
	}

	if dashboard.CurrentCard.ExplanationZH == "" || dashboard.CurrentCard.ExplanationEN == "" {
		t.Fatal("expected bilingual explanations")
	}

	if dashboard.CurrentCard.QuestionTextZH == "" || dashboard.CurrentCard.ClickbaitZH == "" {
		t.Fatal("expected bilingual question and clickbait")
	}

	if dashboard.NotificationSettings.Style != "playful" {
		t.Fatalf("expected default notification style playful, got %s", dashboard.NotificationSettings.Style)
	}

	if len(dashboard.Summary.WeakTopics) == 0 {
		t.Fatal("expected weak topics summary")
	}

	if len(dashboard.ImportErrors) != 0 {
		t.Fatalf("expected no import errors, got %d", len(dashboard.ImportErrors))
	}
}

func TestUpdateSelectedTopicPersistsValueAndFiltersCurrentCard(t *testing.T) {
	app := newTestApp(t)

	dockerCard := `---
id: docker-run-start-container
title_zh: docker run 啟動容器
title_en: Docker Run
type: true-false
tags: [docker, container]
question_zh: "docker run 會建立並啟動一個容器。"
question_en: "docker run creates and starts a container."
clickbait_zh: "第一個 Docker 指令，很多人其實沒真的懂。"
clickbait_en: "The first Docker command many people never really understand."
review_hint_zh: "run = 建立並啟動容器。"
review_hint_en: "run creates and starts a container."
answer: true
enabled: true
---

## zh-TW

docker run 會根據 image 建立並啟動一個容器。
## en

docker run creates and starts a container from an image.
`
	if err := os.MkdirAll(filepath.Join(app.knowledgeDir, "docker"), 0o755); err != nil {
		t.Fatalf("mkdir docker failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(app.knowledgeDir, "docker", "run.md"), []byte(dockerCard), 0o644); err != nil {
		t.Fatalf("write docker card failed: %v", err)
	}
	if _, err := app.RescanKnowledge(); err != nil {
		t.Fatalf("rescan knowledge failed: %v", err)
	}

	status, err := app.UpdateSelectedTopic("docker")
	if err != nil {
		t.Fatalf("update selected topic failed: %v", err)
	}
	if status.Message != "Topic filter updated." {
		t.Fatalf("unexpected status: %s", status.Message)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if dashboard.SelectedTopic != "docker" {
		t.Fatalf("expected selected topic docker, got %s", dashboard.SelectedTopic)
	}
	if dashboard.CurrentCard == nil {
		t.Fatal("expected current card")
	}
	if dashboard.CurrentCard.ID != "docker-run-start-container" {
		t.Fatalf("expected docker card, got %s", dashboard.CurrentCard.ID)
	}
}

func TestUpdateSelectedTopicSupportsPresetGroups(t *testing.T) {
	app := newTestApp(t)

	status, err := app.UpdateSelectedTopic("backend-tools")
	if err != nil {
		t.Fatalf("update selected topic failed: %v", err)
	}
	if status.Message != "Topic filter updated." {
		t.Fatalf("unexpected status: %s", status.Message)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if dashboard.SelectedTopic != "backend-tools" {
		t.Fatalf("expected selected topic backend-tools, got %s", dashboard.SelectedTopic)
	}
	if len(dashboard.AvailableTopics) < 2 {
		t.Fatalf("expected preset topics in available topics, got %#v", dashboard.AvailableTopics)
	}
	if dashboard.AvailableTopics[0] != "all" || dashboard.AvailableTopics[1] != "backend-tools" {
		t.Fatalf("expected presets near the front, got %#v", dashboard.AvailableTopics)
	}
}

func TestSubmitAnswerPersistsFeedback(t *testing.T) {
	app := newTestApp(t)

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	result, err := app.SubmitAnswer(dashboard.CurrentCard.ID, "learn", "true", dashboard.CurrentCard.ShownAt)
	if err != nil {
		t.Fatalf("submit answer failed: %v", err)
	}

	if !result.IsCorrect {
		t.Fatal("expected answer to be correct")
	}

	if result.Stats.StudiedToday != 1 {
		t.Fatalf("expected studied today 1, got %d", result.Stats.StudiedToday)
	}

	if _, err := os.Stat(filepath.Join(app.dataDir, "progress.json")); err != nil {
		t.Fatalf("expected progress.json to exist: %v", err)
	}

	if _, err := os.Stat(filepath.Join(app.dataDir, "attempts.jsonl")); err != nil {
		t.Fatalf("expected attempts.jsonl to exist: %v", err)
	}
}

func TestGetStudyCardReturnsSpecificCard(t *testing.T) {
	app := newTestApp(t)

	card, err := app.GetStudyCard("git-rebase-vs-merge")
	if err != nil {
		t.Fatalf("get study card failed: %v", err)
	}

	if card.ID != "git-rebase-vs-merge" {
		t.Fatalf("unexpected card id: %s", card.ID)
	}
}

func TestCheckAndSendNotificationSendsSelectedCard(t *testing.T) {
	app := newTestApp(t)
	sender := &stubSender{}
	app.notificationSender = sender

	sent, err := app.CheckAndSendNotification()
	if err != nil {
		t.Fatalf("check and send notification failed: %v", err)
	}

	if !sent {
		t.Fatal("expected notification to be sent")
	}

	if sender.message.Title != "哪個 Git 指令可以只拿走一個 commit？" {
		t.Fatalf("unexpected localized title: %s", sender.message.Title)
	}

	if sender.message.ActivationArgument != "duolin-gogo://study/git-cherry-pick-purpose" {
		t.Fatalf("unexpected activation argument: %s", sender.message.ActivationArgument)
	}

	if app.schedulerState.LastNotificationAt == nil {
		t.Fatal("expected last notification timestamp")
	}
}

func TestCheckAndSendNotificationRespectsIntervalAndSnooze(t *testing.T) {
	app := newTestApp(t)
	sender := &stubSender{}
	app.notificationSender = sender

	sent, err := app.CheckAndSendNotification()
	if err != nil {
		t.Fatalf("first send failed: %v", err)
	}
	if !sent {
		t.Fatal("expected first notification to send")
	}

	sent, err = app.CheckAndSendNotification()
	if err != nil {
		t.Fatalf("second send failed: %v", err)
	}
	if sent {
		t.Fatal("expected second notification to be blocked by interval")
	}

	status, err := app.SnoozeNotifications()
	if err != nil {
		t.Fatalf("snooze failed: %v", err)
	}
	if status.Message == "" {
		t.Fatal("expected snooze status message")
	}

	app.schedulerState.LastNotificationAt = nil
	sent, err = app.CheckAndSendNotification()
	if err != nil {
		t.Fatalf("send after snooze failed: %v", err)
	}
	if sent {
		t.Fatal("expected notification to be blocked by snooze")
	}
}

func TestSendTestNotificationUsesSelectedCard(t *testing.T) {
	app := newTestApp(t)
	sender := &stubSender{}
	app.notificationSender = sender

	status, err := app.SendTestNotification()
	if err != nil {
		t.Fatalf("send test notification failed: %v", err)
	}

	if status.Message != "Test notification sent." {
		t.Fatalf("unexpected status: %s", status.Message)
	}

	if sender.message.Title != "哪個 Git 指令可以只拿走一個 commit？" {
		t.Fatalf("unexpected notification title: %s", sender.message.Title)
	}
}

func TestSendTestNotificationReportsPresetScope(t *testing.T) {
	app := newTestApp(t)
	sender := &stubSender{}
	app.notificationSender = sender

	goCard := `---
id: go-goroutine-concurrency
title_zh: goroutine 併發單位
title_en: Goroutine
type: true-false
tags: [go, concurrency]
question_zh: "goroutine 是 Go 裡很輕量的併發執行單位。"
question_en: "A goroutine is a lightweight concurrent execution unit in Go."
clickbait_zh: "Go 很快，不一定是因為你真的會用 goroutine。"
clickbait_en: "Go feels fast, but only if you really understand goroutines."
review_hint_zh: "goroutine 是 Go 的輕量併發單位。"
review_hint_en: "A goroutine is Go's lightweight concurrency primitive."
answer: true
enabled: true
---

## zh-TW

goroutine 是 Go 用來啟動併發工作的輕量執行單位。

## en

A goroutine is Go's lightweight unit for concurrent execution.
`
	if err := os.MkdirAll(filepath.Join(app.knowledgeDir, "go"), 0o755); err != nil {
		t.Fatalf("mkdir go failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(app.knowledgeDir, "go", "goroutine.md"), []byte(goCard), 0o644); err != nil {
		t.Fatalf("write go card failed: %v", err)
	}
	if _, err := app.RescanKnowledge(); err != nil {
		t.Fatalf("rescan knowledge failed: %v", err)
	}
	if _, err := app.UpdateSelectedTopic("languages"); err != nil {
		t.Fatalf("update selected topic failed: %v", err)
	}
	if _, err := app.UpdatePreferredLanguage("en"); err != nil {
		t.Fatalf("update preferred language failed: %v", err)
	}

	status, err := app.SendTestNotification()
	if err != nil {
		t.Fatalf("send test notification failed: %v", err)
	}

	if status.Message != "Test notification sent for Languages." {
		t.Fatalf("unexpected status: %s", status.Message)
	}
	if sender.message.Title == "" {
		t.Fatal("expected notification title")
	}
}

func TestCheckAndSendNotificationUsesTopicAwareReviewCopy(t *testing.T) {
	app := newTestApp(t)
	sender := &stubSender{}
	app.notificationSender = sender
	now := time.Date(2026, 4, 5, 21, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	app.nowFunc = func() time.Time { return now }

	dockerCard := `---
id: docker-run-start-container
title_zh: docker run 建立容器
title_en: Docker Run
type: true-false
tags: [docker, container]
question_zh: "docker run 會建立並啟動容器。"
question_en: "docker run creates and starts a container."
clickbait_zh: "很多人第一個 Docker 指令就沒有真的搞懂。"
clickbait_en: "A lot of people never really understand their first Docker command."
review_hint_zh: "run = 建立並啟動容器。"
review_hint_en: "run creates and starts a container."
answer: true
enabled: true
---

## zh-TW

docker run 會根據 image 建立新的 container 並立即啟動。

## en

docker run creates a container from an image and starts it right away.
`
	if err := os.MkdirAll(filepath.Join(app.knowledgeDir, "docker"), 0o755); err != nil {
		t.Fatalf("mkdir docker failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(app.knowledgeDir, "docker", "run.md"), []byte(dockerCard), 0o644); err != nil {
		t.Fatalf("write docker card failed: %v", err)
	}
	if _, err := app.RescanKnowledge(); err != nil {
		t.Fatalf("rescan knowledge failed: %v", err)
	}
	if _, err := app.UpdateSelectedTopic("backend-tools"); err != nil {
		t.Fatalf("update selected topic failed: %v", err)
	}
	if _, err := app.UpdatePreferredLanguage("en"); err != nil {
		t.Fatalf("update preferred language failed: %v", err)
	}

	shownAt := now.Add(-49 * time.Hour)
	answeredAt := now.Add(-48 * time.Hour)
	if _, _, err := progress.RecordAndPersist(
		filepath.Join(app.dataDir, "progress.json"),
		filepath.Join(app.dataDir, "attempts.jsonl"),
		progress.RecordAttemptInput{
			CardID:         "docker-run-start-container",
			SessionType:    "learn",
			SelectedAnswer: "true",
			IsCorrect:      true,
			ShownAt:        shownAt,
			AnsweredAt:     answeredAt,
		},
	); err != nil {
		t.Fatalf("record attempt failed: %v", err)
	}

	sent, err := app.CheckAndSendNotification()
	if err != nil {
		t.Fatalf("check and send notification failed: %v", err)
	}
	if !sent {
		t.Fatal("expected review notification to be sent")
	}

	if sender.message.Title != "Review time" {
		t.Fatalf("unexpected review title: %s", sender.message.Title)
	}
	if sender.message.Body != "Time to review your recent backend tools concepts." {
		t.Fatalf("unexpected review body: %s", sender.message.Body)
	}
}
func TestSendTestNotificationReportsWhenNoCardsExist(t *testing.T) {
	app := newTestApp(t)
	sender := &stubSender{}
	app.notificationSender = sender

	if err := os.RemoveAll(app.knowledgeDir); err != nil {
		t.Fatalf("remove knowledge dir failed: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(app.knowledgeDir, "git"), 0o755); err != nil {
		t.Fatalf("recreate knowledge dir failed: %v", err)
	}

	status, err := app.SendTestNotification()
	if err != nil {
		t.Fatalf("send test notification failed: %v", err)
	}

	if status.Message != "No card available for test notification." {
		t.Fatalf("unexpected status: %s", status.Message)
	}
}

func TestRescanKnowledgeRefreshesCacheWithNewCard(t *testing.T) {
	app := newTestApp(t)

	newCard := `---
id: git-fast-forward-merge
title: Fast-forward Merge
title_zh: Fast-forward Merge
title_en: Fast-forward Merge
type: true-false
tags: [git, branching]
question_zh: "` + "`git merge --ff-only`" + ` 不允許產生 merge commit。"
question_en: "` + "`git merge --ff-only`" + ` refuses to create a merge commit."
clickbait_zh: "這個 merge 選項其實比你想的更嚴格"
clickbait_en: "This merge flag is stricter than you think"
answer: true
enabled: true
---

## zh-TW

` + "`git merge --ff-only`" + ` 只允許 fast-forward merge，不能產生 merge commit。
## en

` + "`git merge --ff-only`" + ` only allows a fast-forward merge and refuses to create a merge commit.
`

	if err := os.WriteFile(filepath.Join(app.knowledgeDir, "git", "fast-forward.md"), []byte(newCard), 0o644); err != nil {
		t.Fatalf("write new card failed: %v", err)
	}

	status, err := app.RescanKnowledge()
	if err != nil {
		t.Fatalf("rescan knowledge failed: %v", err)
	}

	if status.Message != "Knowledge refreshed: 3 cards, 0 errors." {
		t.Fatalf("unexpected rescan status: %s", status.Message)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if dashboard.CurrentCard == nil {
		t.Fatal("expected current card after rescan")
	}

	cache, err := loadCache(filepath.Join(app.dataDir, "cards-cache.json"))
	if err != nil {
		t.Fatalf("load cache failed: %v", err)
	}

	if len(cache.Cards) != 3 {
		t.Fatalf("expected 3 cached cards after rescan, got %d", len(cache.Cards))
	}
}

func TestValidateKnowledgeUpdatesDiagnosticsWithoutRefreshingCache(t *testing.T) {
	app := newTestApp(t)

	cacheBefore, err := loadCache(filepath.Join(app.dataDir, "cards-cache.json"))
	if err != nil {
		t.Fatalf("load cache before validate failed: %v", err)
	}

	broken := `---
id: git-validate-broken
title: Broken
type: true-false
question: "Broken?"
answer: true
---

## zh-TW

只有中文`
	if err := os.WriteFile(filepath.Join(app.knowledgeDir, "git", "broken.md"), []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken file failed: %v", err)
	}

	status, err := app.ValidateKnowledge()
	if err != nil {
		t.Fatalf("validate knowledge failed: %v", err)
	}

	if status.Message != "Knowledge validated: 2 cards, 1 diagnostics." {
		t.Fatalf("unexpected validate status: %s", status.Message)
	}
	if len(status.ImportErrors) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(status.ImportErrors))
	}
	if status.ImportErrors[0].Code != "missing_language_section" {
		t.Fatalf("unexpected diagnostic code: %s", status.ImportErrors[0].Code)
	}

	cacheAfter, err := loadCache(filepath.Join(app.dataDir, "cards-cache.json"))
	if err != nil {
		t.Fatalf("load cache after validate failed: %v", err)
	}
	if len(cacheAfter.Cards) != len(cacheBefore.Cards) {
		t.Fatalf("expected cache size to stay %d, got %d", len(cacheBefore.Cards), len(cacheAfter.Cards))
	}
}

func TestLoadAuthoringPreviewReturnsFilesAndSelectedCard(t *testing.T) {
	app := newTestApp(t)

	preview, err := app.LoadAuthoringPreview()
	if err != nil {
		t.Fatalf("load authoring preview failed: %v", err)
	}

	if len(preview.Files) != 2 {
		t.Fatalf("expected 2 preview files, got %d", len(preview.Files))
	}
	if preview.SelectedPath == "" {
		t.Fatal("expected selected preview path")
	}
	if preview.CurrentCard == nil {
		t.Fatal("expected preview card")
	}
	if preview.CurrentCard.ID == "" {
		t.Fatal("expected preview card id")
	}
}

func TestLoadAuthoringPreviewIncludesModifiedAt(t *testing.T) {
	app := newTestApp(t)

	preview, err := app.LoadAuthoringPreview()
	if err != nil {
		t.Fatalf("load authoring preview failed: %v", err)
	}

	if len(preview.Files) == 0 {
		t.Fatal("expected preview files")
	}

	for _, file := range preview.Files {
		if file.ModifiedAt == "" {
			t.Fatalf("expected modifiedAt for %s", file.Path)
		}
		if _, err := time.Parse(time.RFC3339, file.ModifiedAt); err != nil {
			t.Fatalf("expected RFC3339 modifiedAt for %s: %v", file.Path, err)
		}
	}
}

func TestPreviewKnowledgeCardReturnsDiagnosticsForBrokenCard(t *testing.T) {
	app := newTestApp(t)

	brokenPath := filepath.Join(app.knowledgeDir, "git", "preview-broken.md")
	broken := `---
id: git-preview-broken
title: Broken Preview
type: true-false
question: "Broken?"
answer: true
---

## zh-TW

Only zh body.`
	if err := os.WriteFile(brokenPath, []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken preview file failed: %v", err)
	}

	preview, err := app.PreviewKnowledgeCard(brokenPath)
	if err != nil {
		t.Fatalf("preview knowledge card failed: %v", err)
	}

	if preview.SelectedPath != brokenPath {
		t.Fatalf("expected selected path %s, got %s", brokenPath, preview.SelectedPath)
	}
	if preview.CurrentCard != nil {
		t.Fatal("expected no preview card for broken file")
	}
	if len(preview.ImportErrors) == 0 {
		t.Fatal("expected preview diagnostics")
	}
	if preview.ImportErrors[0].Code != "missing_language_section" {
		t.Fatalf("unexpected diagnostic code: %s", preview.ImportErrors[0].Code)
	}
}

func TestReviewDraftReturnsPreviewAndDiagnostics(t *testing.T) {
	app := newTestApp(t)

	raw := `---
id: git-ai-draft
title: AI Draft
type: single-choice
question: "Which command downloads remote refs without merging?"
choices:
  - "git fetch"
  - "git pull"
answer: 0
---

## zh-TW

這是一張 AI 草稿卡，使用 fallback 欄位。

## en

This AI draft card relies on fallback fields.`

	result, err := app.ReviewDraft(raw)
	if err != nil {
		t.Fatalf("review draft failed: %v", err)
	}

	if result.CurrentCard == nil {
		t.Fatal("expected normalized draft card")
	}
	if result.CurrentCard.ID != "git-ai-draft" {
		t.Fatalf("unexpected draft card id: %s", result.CurrentCard.ID)
	}
	if len(result.ImportErrors) == 0 {
		t.Fatal("expected draft diagnostics")
	}
}

func TestSaveDraftPersistsMarkdownIntoTopicFolder(t *testing.T) {
	app := newTestApp(t)

	raw := `---
id: git-save-draft
title_zh: 儲存草稿
title_en: Save Draft
type: true-false
question_zh: "git fetch 會直接 merge 到目前分支。"
question_en: "git fetch merges into the current branch."
clickbait_zh: "這個指令看起來沒做事，但很多人會先按"
clickbait_en: "This command looks quiet, but many people use it first"
review_hint_zh: "fetch 不會直接 merge。"
review_hint_en: "Fetch does not merge directly."
answer: false
---

## zh-TW

git fetch 只會更新追蹤資訊，不會直接 merge。

## en

git fetch only updates tracking refs and does not merge directly.`

	status, err := app.SaveDraft(raw, "git")
	if err != nil {
		t.Fatalf("save draft failed: %v", err)
	}

	if !status.Successful {
		t.Fatal("expected successful save")
	}

	expectedPath := filepath.Join(app.knowledgeDir, "git", "git-save-draft.md")
	if status.SavedPath != expectedPath {
		t.Fatalf("unexpected saved path: %s", status.SavedPath)
	}
	if _, err := os.Stat(expectedPath); err != nil {
		t.Fatalf("expected saved draft file: %v", err)
	}
}

func TestUpdateNotificationSettingsPersistsValues(t *testing.T) {
	app := newTestApp(t)

	status, err := app.UpdateNotificationSettings("chaotic", "prefer_generated")
	if err != nil {
		t.Fatalf("update notification settings failed: %v", err)
	}

	if status.Message != "Notification settings updated." {
		t.Fatalf("unexpected status: %s", status.Message)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if dashboard.NotificationSettings.Style != "chaotic" {
		t.Fatalf("expected updated style, got %s", dashboard.NotificationSettings.Style)
	}
	if dashboard.NotificationSettings.TitleMode != "prefer_generated" {
		t.Fatalf("expected updated title mode, got %s", dashboard.NotificationSettings.TitleMode)
	}
}

func TestUpdatePreferredLanguagePersistsValue(t *testing.T) {
	app := newTestApp(t)

	status, err := app.UpdatePreferredLanguage("en")
	if err != nil {
		t.Fatalf("update preferred language failed: %v", err)
	}

	if status.Message != "Language updated." {
		t.Fatalf("unexpected status: %s", status.Message)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if dashboard.PreferredLanguage != "en" {
		t.Fatalf("expected preferred language en, got %s", dashboard.PreferredLanguage)
	}
}

func TestUpdateScheduleSettingsPersistsValues(t *testing.T) {
	app := newTestApp(t)

	previous := app.nowFunc().Add(-30 * time.Minute)
	snoozedUntil := app.nowFunc().Add(20 * time.Minute)
	app.schedulerState.LastNotificationAt = &previous
	app.schedulerState.SnoozedUntil = &snoozedUntil

	status, err := app.UpdateScheduleSettings(30, "20:30", true, "08:30", "23:30")
	if err != nil {
		t.Fatalf("update schedule settings failed: %v", err)
	}

	if status.Message != "Schedule settings updated." {
		t.Fatalf("unexpected status: %s", status.Message)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if dashboard.ScheduleSettings.NotificationIntervalMinutes != 30 {
		t.Fatalf("expected notification interval 30, got %d", dashboard.ScheduleSettings.NotificationIntervalMinutes)
	}
	if dashboard.ScheduleSettings.ReviewTime != "20:30" {
		t.Fatalf("expected review time 20:30, got %s", dashboard.ScheduleSettings.ReviewTime)
	}
	if !dashboard.ScheduleSettings.ActiveHoursEnabled {
		t.Fatal("expected active hours to stay enabled")
	}
	if dashboard.ScheduleSettings.ActiveHoursStart != "08:30" {
		t.Fatalf("expected active hours start 08:30, got %s", dashboard.ScheduleSettings.ActiveHoursStart)
	}
	if dashboard.ScheduleSettings.ActiveHoursEnd != "23:30" {
		t.Fatalf("expected active hours end 23:30, got %s", dashboard.ScheduleSettings.ActiveHoursEnd)
	}

	if app.schedulerState.LastNotificationAt == nil {
		t.Fatal("expected scheduler timestamp to reset")
	}
	if app.schedulerState.SnoozedUntil != nil {
		t.Fatal("expected snooze to clear after schedule update")
	}
}

func TestResetStudyDataClearsProgressAndAttempts(t *testing.T) {
	app := newTestApp(t)

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if _, err := app.SubmitAnswer(dashboard.CurrentCard.ID, "learn", "true", dashboard.CurrentCard.ShownAt); err != nil {
		t.Fatalf("submit answer failed: %v", err)
	}

	if _, err := os.Stat(filepath.Join(app.dataDir, "progress.json")); err != nil {
		t.Fatalf("expected progress file before reset: %v", err)
	}
	if _, err := os.Stat(filepath.Join(app.dataDir, "attempts.jsonl")); err != nil {
		t.Fatalf("expected attempts file before reset: %v", err)
	}

	status, err := app.ResetStudyData()
	if err != nil {
		t.Fatalf("reset study data failed: %v", err)
	}

	if status.Message != "Study data reset." {
		t.Fatalf("unexpected reset status: %s", status.Message)
	}

	if _, err := os.Stat(filepath.Join(app.dataDir, "progress.json")); !os.IsNotExist(err) {
		t.Fatalf("expected progress file to be removed, got %v", err)
	}
	if _, err := os.Stat(filepath.Join(app.dataDir, "attempts.jsonl")); !os.IsNotExist(err) {
		t.Fatalf("expected attempts file to be removed, got %v", err)
	}

	after, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard after reset failed: %v", err)
	}

	if after.Stats.StudiedToday != 0 {
		t.Fatalf("expected studiedToday 0 after reset, got %d", after.Stats.StudiedToday)
	}
	if after.Stats.CorrectRate != 0 {
		t.Fatalf("expected correctRate 0 after reset, got %f", after.Stats.CorrectRate)
	}
}

func TestLoadDashboardEntersReviewModeWhenReviewIsDue(t *testing.T) {
	app := newTestApp(t)
	now := time.Date(2026, 4, 5, 21, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	app.nowFunc = func() time.Time { return now }

	_, _, err := progress.RecordAndPersist(
		filepath.Join(app.dataDir, "progress.json"),
		filepath.Join(app.dataDir, "attempts.jsonl"),
		progress.RecordAttemptInput{
			CardID:         "git-rebase-vs-merge",
			SessionType:    "learn",
			SelectedAnswer: "1",
			IsCorrect:      true,
			ShownAt:        now.Add(-24 * time.Hour),
			AnsweredAt:     now.Add(-24*time.Hour + 10*time.Second),
		},
	)
	if err != nil {
		t.Fatalf("seed progress failed: %v", err)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if !dashboard.ReviewMode {
		t.Fatal("expected review mode")
	}

	if len(dashboard.ReviewQueue) == 0 {
		t.Fatal("expected review queue")
	}
}

func TestLoadDashboardShowsImportDiagnosticsWithoutCrashing(t *testing.T) {
	app := newTestApp(t)

	broken := `---
id: git-broken-card
title: Broken Card
type: true-false
question: "Broken?"
answer: true
---

## zh-TW

只有中文。`

	if err := os.WriteFile(filepath.Join(app.knowledgeDir, "git", "broken.md"), []byte(broken), 0o644); err != nil {
		t.Fatalf("write broken card failed: %v", err)
	}

	dashboard, err := app.LoadDashboard()
	if err != nil {
		t.Fatalf("load dashboard failed: %v", err)
	}

	if len(dashboard.ImportErrors) == 0 {
		t.Fatal("expected import diagnostics")
	}

	if dashboard.CurrentCard == nil {
		t.Fatal("expected dashboard to still load a valid card")
	}
}

func newTestApp(t *testing.T) *App {
	t.Helper()

	root := t.TempDir()
	knowledgeDir := filepath.Join(root, "knowledge", "git")
	dataDir := filepath.Join(root, "data")

	if err := os.MkdirAll(knowledgeDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	rebase := `---
id: git-rebase-vs-merge
title: Rebase vs Merge
title_zh: Rebase 跟 Merge 的差別
title_en: Rebase vs Merge
type: single-choice
tags: [git, branching]
question_zh: "git rebase 主要是在做什麼？"
question_en: "What does git rebase mainly do?"
choices_zh:
  - "建立一個 merge commit"
  - "把 commits 重新接到新的 base 上"
choices_en:
  - "Creates a merge commit between branches"
  - "Replays commits onto a new base"
answer: 1
clickbait_zh: "你真的懂 rebase 跟 merge 的差別嗎？"
clickbait_en: "Do you really know the difference between rebase and merge?"
review_hint_zh: "Rebase = 把 commits 重放到新的 base 上。"
review_hint_en: "Rebase = replay commits on top of another base."
enabled: true
---

## zh-TW

` + "`git rebase` 會把目前分支上的 commits 重新接到另一個 base 上。" + `
## en

English rebase explanation.
`

	cherryPick := `---
id: git-cherry-pick-purpose
title: Cherry-pick Purpose
title_zh: Cherry-pick 的用途
title_en: Cherry-pick Purpose
type: true-false
tags: [git, commits]
question_zh: "` + "`git cherry-pick` 會把指定的一個 commit 套用到目前分支上。" + `"
question_en: "` + "`git cherry-pick` applies a chosen commit to the current branch." + `"
clickbait_zh: "哪個 Git 指令可以只拿走一個 commit？"
clickbait_en: "One Git command can steal just one commit. Know which?"
review_hint_zh: "Cherry-pick 會把選定 commit 的變更套到目前分支。"
review_hint_en: "Cherry-pick copies selected commit changes onto your current branch."
answer: true
enabled: true
---

## zh-TW

` + "`git cherry-pick` 會把你指定的一個 commit 套用到目前分支上。" + `
## en

English cherry-pick explanation.
`

	if err := os.WriteFile(filepath.Join(knowledgeDir, "rebase.md"), []byte(rebase), 0o644); err != nil {
		t.Fatalf("write rebase failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(knowledgeDir, "cherry-pick.md"), []byte(cherryPick), 0o644); err != nil {
		t.Fatalf("write cherry-pick failed: %v", err)
	}

	settings := `{
  "notification_interval_minutes": 10,
  "active_hours": {
    "enabled": true,
    "start": "09:00",
    "end": "22:00"
  },
  "review_schedule": {
    "mode": "daily",
    "weekday": null,
    "time": "21:00",
    "batch_size": 5
  },
  "study_rules": {
    "snooze_minutes": 15
  },
  "language": {
    "default": "zh-TW"
  }
}`
	if err := os.WriteFile(filepath.Join(dataDir, "settings.json"), []byte(settings), 0o644); err != nil {
		t.Fatalf("write settings failed: %v", err)
	}

	app := NewAppWithPaths(filepath.Join(root, "knowledge"), dataDir)
	app.nowFunc = func() time.Time {
		return time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	}

	if _, err := cards.RefreshKnowledge(filepath.Join(root, "knowledge"), dataDir); err != nil {
		t.Fatalf("refresh knowledge failed: %v", err)
	}

	return app
}

type stubSender struct {
	message notifications.Message
}

func (s *stubSender) Send(message notifications.Message) error {
	s.message = message
	return nil
}
