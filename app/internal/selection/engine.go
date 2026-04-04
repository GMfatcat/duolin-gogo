package selection

import (
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/progress"
)

func PriorityScore(card cards.Card, state progress.CardProgress, now time.Time) int {
	score := 0

	if !card.Enabled {
		return -1
	}

	if state.SeenCount == 0 {
		score += 30
	}

	if isReviewDue(state.NextReviewAt, now) {
		score += 25
	}

	if isRecentWrong(state.LastWrongAt, now) {
		score += 20
	}

	score += timeSinceSeenBonus(state.LastSeenAt, now)
	score += weaknessBonus(state)
	score -= masteryPenalty(state)

	return score
}

func SelectNextCard(allCards []cards.Card, states map[string]progress.CardProgress, now time.Time) (cards.Card, bool) {
	var best cards.Card
	bestScore := -1
	found := false

	for _, card := range allCards {
		if !card.Enabled {
			continue
		}

		score := PriorityScore(card, states[card.ID], now)
		if !found || score > bestScore {
			best = card
			bestScore = score
			found = true
		}
	}

	return best, found
}

func isReviewDue(nextReviewAt *string, now time.Time) bool {
	if nextReviewAt == nil || *nextReviewAt == "" {
		return false
	}

	t, err := time.Parse(time.RFC3339, *nextReviewAt)
	if err != nil {
		return false
	}

	return !t.After(now)
}

func isRecentWrong(lastWrongAt *string, now time.Time) bool {
	if lastWrongAt == nil || *lastWrongAt == "" {
		return false
	}

	t, err := time.Parse(time.RFC3339, *lastWrongAt)
	if err != nil {
		return false
	}

	return now.Sub(t) <= 24*time.Hour
}

func timeSinceSeenBonus(lastSeenAt *string, now time.Time) int {
	if lastSeenAt == nil || *lastSeenAt == "" {
		return 0
	}

	t, err := time.Parse(time.RFC3339, *lastSeenAt)
	if err != nil {
		return 0
	}

	days := int(now.Sub(t).Hours() / 24)
	if days <= 0 {
		return 0
	}

	bonus := days * 5
	if bonus > 20 {
		return 20
	}

	return bonus
}

func weaknessBonus(state progress.CardProgress) int {
	if state.SeenCount == 0 {
		return 0
	}

	bonus := state.WrongCount*8 - state.CorrectCount*2
	if bonus < 0 {
		return 0
	}

	return bonus
}

func masteryPenalty(state progress.CardProgress) int {
	penalty := state.MasteryScore * 3
	if state.IsMastered {
		penalty += 10
	}
	if penalty < 0 {
		return 0
	}

	return penalty
}
