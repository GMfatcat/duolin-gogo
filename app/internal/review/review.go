package review

import (
	"slices"
	"strings"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/progress"
	"duolin-gogo/internal/settings"
)

func ShouldStartReview(config settings.File, lastRun *time.Time, now time.Time) bool {
	mode := strings.ToLower(config.ReviewSchedule.Mode)
	if mode == "" || mode == "off" {
		return false
	}

	if !matchesReviewTime(config.ReviewSchedule.Time, now) {
		return false
	}

	if mode == "weekly" && !matchesWeekday(config.ReviewSchedule.Weekday, now) {
		return false
	}

	if lastRun != nil && sameReviewWindow(mode, *lastRun, now) {
		return false
	}

	return true
}

func BuildQueue(allCards []cards.Card, states map[string]progress.CardProgress, now time.Time, batchSize int) []cards.Card {
	if batchSize <= 0 {
		batchSize = 5
	}

	type candidate struct {
		card  cards.Card
		score int
	}

	var candidates []candidate
	for _, card := range allCards {
		if !card.Enabled {
			continue
		}

		state := states[card.ID]
		if state.SeenCount == 0 {
			continue
		}

		score := reviewPriority(state, now)
		if score <= 0 {
			continue
		}

		candidates = append(candidates, candidate{card: card, score: score})
	}

	slices.SortFunc(candidates, func(a, b candidate) int {
		if a.score == b.score {
			return strings.Compare(a.card.ID, b.card.ID)
		}
		if a.score > b.score {
			return -1
		}
		return 1
	})

	if len(candidates) > batchSize {
		candidates = candidates[:batchSize]
	}

	queue := make([]cards.Card, 0, len(candidates))
	for _, item := range candidates {
		queue = append(queue, item.card)
	}

	return queue
}

func matchesReviewTime(configured string, now time.Time) bool {
	t, err := time.Parse("15:04", configured)
	if err != nil {
		return false
	}

	return now.Hour() == t.Hour() && now.Minute() == t.Minute()
}

func matchesWeekday(weekday *string, now time.Time) bool {
	if weekday == nil {
		return false
	}

	return strings.EqualFold(*weekday, strings.ToLower(now.Weekday().String()[:3]))
}

func sameReviewWindow(mode string, lastRun time.Time, now time.Time) bool {
	switch mode {
	case "daily":
		y1, m1, d1 := lastRun.Date()
		y2, m2, d2 := now.Date()
		return y1 == y2 && m1 == m2 && d1 == d2
	case "weekly":
		y1, w1 := lastRun.ISOWeek()
		y2, w2 := now.ISOWeek()
		return y1 == y2 && w1 == w2
	default:
		return false
	}
}

func reviewPriority(state progress.CardProgress, now time.Time) int {
	score := 0

	if isDue(state.NextReviewAt, now) {
		score += 30
	}

	score += state.WrongCount * 8
	score -= state.CorrectCount * 2

	if state.LastSeenAt != nil {
		if t, err := time.Parse(time.RFC3339, *state.LastSeenAt); err == nil {
			days := int(now.Sub(t).Hours() / 24)
			if days > 0 {
				score += min(days*4, 16)
			}
		}
	}

	if score < 0 {
		return 0
	}

	return score
}

func isDue(nextReviewAt *string, now time.Time) bool {
	if nextReviewAt == nil || *nextReviewAt == "" {
		return false
	}

	t, err := time.Parse(time.RFC3339, *nextReviewAt)
	if err != nil {
		return false
	}

	return !t.After(now)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
