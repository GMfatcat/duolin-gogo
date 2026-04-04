package scheduler

import (
	"testing"
	"time"

	"duolin-gogo/internal/settings"
)

func TestShouldSendLearningNotificationWithinActiveHours(t *testing.T) {
	config := testSettings()
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	if !ShouldSendLearningNotification(config, State{}, now) {
		t.Fatal("expected notification to be allowed")
	}
}

func TestShouldNotSendLearningNotificationOutsideActiveHours(t *testing.T) {
	config := testSettings()
	now := time.Date(2026, 4, 5, 23, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	if ShouldSendLearningNotification(config, State{}, now) {
		t.Fatal("expected notification to be blocked outside active hours")
	}
}

func TestShouldNotSendLearningNotificationWhenIntervalNotElapsed(t *testing.T) {
	config := testSettings()
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	last := now.Add(-5 * time.Minute)

	if ShouldSendLearningNotification(config, State{LastNotificationAt: &last}, now) {
		t.Fatal("expected notification to be blocked by interval")
	}
}

func TestShouldNotSendLearningNotificationWhenSnoozed(t *testing.T) {
	config := testSettings()
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
	snoozedUntil := now.Add(10 * time.Minute)

	if ShouldSendLearningNotification(config, State{SnoozedUntil: &snoozedUntil}, now) {
		t.Fatal("expected notification to be blocked by snooze")
	}
}

func testSettings() settings.File {
	file := settings.File{}
	file.NotificationIntervalMinutes = 10
	file.ActiveHours.Enabled = true
	file.ActiveHours.Start = "09:00"
	file.ActiveHours.End = "22:00"
	return file
}
