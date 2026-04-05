package dashboard

import (
	"path/filepath"
	"slices"
	"strings"
	"time"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/progress"
)

type WeakTopic struct {
	Tag        string  `json:"tag"`
	WrongCount int     `json:"wrongCount"`
	SeenCount  int     `json:"seenCount"`
	Accuracy   float64 `json:"accuracy"`
}

type Summary struct {
	StudiedToday  int             `json:"studiedToday"`
	CorrectRate   float64         `json:"correctRate"`
	NextReviewAt  string          `json:"nextReviewAt"`
	WeakTopics    []WeakTopic     `json:"weakTopics"`
	TopicProgress []TopicProgress `json:"topicProgress"`
	WeakestDeck   *TopicProgress  `json:"weakestDeck,omitempty"`
}

type TopicProgress struct {
	Topic       string  `json:"topic"`
	SeenCount   int     `json:"seenCount"`
	CorrectCount int    `json:"correctCount"`
	WrongCount  int     `json:"wrongCount"`
	Accuracy    float64 `json:"accuracy"`
}

func BuildSummary(allCards []cards.Card, state progress.ProgressFile, now time.Time) Summary {
	day := state.DailySummary[now.Format("2006-01-02")]
	correctRate := 0.0
	if day.Answered > 0 {
		correctRate = float64(day.Correct) / float64(day.Answered)
	}

	progressItems := topicProgress(allCards, state.Cards)

	return Summary{
		StudiedToday:  day.Answered,
		CorrectRate:   correctRate,
		NextReviewAt:  nextReviewTime(allCards, state.Cards, now),
		WeakTopics:    weakTopics(allCards, state.Cards),
		TopicProgress: progressItems,
		WeakestDeck:   weakestDeck(progressItems),
	}
}

func weakestDeck(items []TopicProgress) *TopicProgress {
	if len(items) <= 1 {
		return nil
	}

	weakest := items[0]
	return &weakest
}

func nextReviewTime(allCards []cards.Card, states map[string]progress.CardProgress, now time.Time) string {
	var next *time.Time
	allowed := map[string]struct{}{}
	for _, card := range allCards {
		allowed[card.ID] = struct{}{}
	}

	for cardID, state := range states {
		if len(allowed) > 0 {
			if _, ok := allowed[cardID]; !ok {
				continue
			}
		}

		if state.NextReviewAt == nil || *state.NextReviewAt == "" {
			continue
		}
		parsed, err := time.Parse(time.RFC3339, *state.NextReviewAt)
		if err != nil {
			continue
		}
		if parsed.Before(now) {
			continue
		}
		if next == nil || parsed.Before(*next) {
			copy := parsed
			next = &copy
		}
	}

	if next == nil {
		return ""
	}

	return next.Format(time.RFC3339)
}

func weakTopics(allCards []cards.Card, states map[string]progress.CardProgress) []WeakTopic {
	type aggregate struct {
		wrong int
		seen  int
	}

	tagScores := map[string]aggregate{}
	for _, card := range allCards {
		state, ok := states[card.ID]
		if !ok || state.SeenCount == 0 {
			continue
		}

		for _, tag := range card.Tags {
			agg := tagScores[tag]
			agg.wrong += state.WrongCount
			agg.seen += state.SeenCount
			tagScores[tag] = agg
		}
	}

	topics := make([]WeakTopic, 0, len(tagScores))
	for tag, agg := range tagScores {
		accuracy := 0.0
		if agg.seen > 0 {
			accuracy = float64(agg.seen-agg.wrong) / float64(agg.seen)
		}
		topics = append(topics, WeakTopic{
			Tag:        tag,
			WrongCount: agg.wrong,
			SeenCount:  agg.seen,
			Accuracy:   accuracy,
		})
	}

	slices.SortFunc(topics, func(a, b WeakTopic) int {
		if a.Accuracy == b.Accuracy {
			if a.WrongCount == b.WrongCount {
				if a.Tag < b.Tag {
					return -1
				}
				return 1
			}
			if a.WrongCount > b.WrongCount {
				return -1
			}
			return 1
		}
		if a.Accuracy < b.Accuracy {
			return -1
		}
		return 1
	})

	if len(topics) > 5 {
		topics = topics[:5]
	}

	return topics
}

func topicProgress(allCards []cards.Card, states map[string]progress.CardProgress) []TopicProgress {
	type aggregate struct {
		seen    int
		correct int
		wrong   int
	}

	topicScores := map[string]aggregate{}
	for _, card := range allCards {
		state, ok := states[card.ID]
		if !ok || state.SeenCount == 0 {
			continue
		}

		topic := topicForCard(card)
		if topic == "" {
			continue
		}

		agg := topicScores[topic]
		agg.seen += state.SeenCount
		agg.correct += state.CorrectCount
		agg.wrong += state.WrongCount
		topicScores[topic] = agg
	}

	items := make([]TopicProgress, 0, len(topicScores))
	for topic, agg := range topicScores {
		accuracy := 0.0
		if agg.seen > 0 {
			accuracy = float64(agg.correct) / float64(agg.seen)
		}
		items = append(items, TopicProgress{
			Topic:        topic,
			SeenCount:    agg.seen,
			CorrectCount: agg.correct,
			WrongCount:   agg.wrong,
			Accuracy:     accuracy,
		})
	}

	slices.SortFunc(items, func(a, b TopicProgress) int {
		if a.Accuracy == b.Accuracy {
			if a.SeenCount == b.SeenCount {
				if a.Topic < b.Topic {
					return -1
				}
				return 1
			}
			if a.SeenCount > b.SeenCount {
				return -1
			}
			return 1
		}
		if a.Accuracy < b.Accuracy {
			return -1
		}
		return 1
	})

	return items
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
