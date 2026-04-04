package scheduler

import (
	"time"

	"duolin-gogo/internal/settings"
)

type State struct {
	LastNotificationAt *time.Time
	SnoozedUntil       *time.Time
}

func ShouldSendLearningNotification(config settings.File, state State, now time.Time) bool {
	if config.ActiveHours.Enabled && !withinActiveHours(config.ActiveHours.Start, config.ActiveHours.End, now) {
		return false
	}

	if state.SnoozedUntil != nil && state.SnoozedUntil.After(now) {
		return false
	}

	if state.LastNotificationAt == nil {
		return true
	}

	interval := time.Duration(config.NotificationIntervalMinutes) * time.Minute
	if interval <= 0 {
		interval = 10 * time.Minute
	}

	return now.Sub(*state.LastNotificationAt) >= interval
}

func withinActiveHours(start, end string, now time.Time) bool {
	startTime, err := time.Parse("15:04", start)
	if err != nil {
		return true
	}
	endTime, err := time.Parse("15:04", end)
	if err != nil {
		return true
	}

	currentMinutes := now.Hour()*60 + now.Minute()
	startMinutes := startTime.Hour()*60 + startTime.Minute()
	endMinutes := endTime.Hour()*60 + endTime.Minute()

	if startMinutes <= endMinutes {
		return currentMinutes >= startMinutes && currentMinutes <= endMinutes
	}

	return currentMinutes >= startMinutes || currentMinutes <= endMinutes
}
