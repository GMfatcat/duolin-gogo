package notifications

import (
	"fmt"
	"strings"

	"duolin-gogo/internal/cards"
	"duolin-gogo/internal/hooks"
)

const AppID = "duolin-gogo"

type Message struct {
	Title              string
	Body               string
	ActivationArgument string
}

type Sender interface {
	Send(Message) error
}

func BuildStudyMessage(card cards.Card) Message {
	return BuildStudyMessageForLanguage(card, "en", "playful", "prefer_manual")
}

func BuildStudyMessageForLanguage(card cards.Card, language string, style string, titleMode string) Message {
	manualTitle := card.ClickbaitEN
	if language == "zh-TW" {
		manualTitle = card.ClickbaitZH
	}

	generatedTitle, _ := hooks.Generate(card, language, style)
	title, body := resolveContent(manualTitle, generatedTitle, titleMode)
	if title == "" {
		title = fallbackTitle(card, language)
	}

	return Message{
		Title:              title,
		Body:               body,
		ActivationArgument: fmt.Sprintf("duolin-gogo://study/%s", card.ID),
	}
}

func resolveContent(manualTitle string, generatedTitle string, titleMode string) (string, string) {
	switch normalizeTitleMode(titleMode) {
	case "prefer_generated":
		return firstNonEmpty(generatedTitle, manualTitle), ""
	default:
		return firstNonEmpty(manualTitle, generatedTitle), ""
	}
}

func normalizeTitleMode(titleMode string) string {
	switch strings.TrimSpace(strings.ToLower(titleMode)) {
	case "prefer_generated":
		return "prefer_generated"
	default:
		return "prefer_manual"
	}
}

func fallbackTitle(card cards.Card, language string) string {
	if language == "zh-TW" {
		return firstNonEmpty(card.TitleZH, card.Title, card.TitleEN)
	}
	return firstNonEmpty(card.TitleEN, card.Title, card.TitleZH)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
