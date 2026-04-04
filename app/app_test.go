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

	if len(dashboard.Summary.WeakTopics) == 0 {
		t.Fatal("expected weak topics summary")
	}

	if len(dashboard.ImportErrors) != 0 {
		t.Fatalf("expected no import errors, got %d", len(dashboard.ImportErrors))
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

	if sender.message.Title != "測試通知" {
		t.Fatalf("unexpected notification title: %s", sender.message.Title)
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
