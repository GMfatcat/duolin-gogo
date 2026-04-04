package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/notifications"
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

func TestLoadDashboardReturnsStudyCardAndStats(t *testing.T) {
	app := newTestApp(t)

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

	if err := app.SnoozeNotifications(); err != nil {
		t.Fatalf("snooze failed: %v", err)
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
type: single-choice
question: "What does git rebase mainly do?"
choices:
  - "Creates a merge commit between branches"
  - "Replays commits onto a new base"
answer: 1
enabled: true
---

## zh-TW

中文 rebase 解釋。

## en

English rebase explanation.
`

	cherryPick := `---
id: git-cherry-pick-purpose
title: Cherry-pick Purpose
type: true-false
question: "` + "`git cherry-pick` applies a chosen commit to the current branch." + `"
answer: true
enabled: true
---

## zh-TW

中文 cherry-pick 解釋。

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
