package notifications

import (
	"fmt"

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
	return BuildStudyMessageForLanguage(card, "en", "playful")
}

func BuildStudyMessageForLanguage(card cards.Card, language string, style string) Message {
	title := card.ClickbaitEN
	body := card.QuestionTextEN
	if language == "zh-TW" {
		title = card.ClickbaitZH
		body = card.QuestionTextZH
	}

	if title == "" {
		title, body = hooks.Generate(card, language, style)
	}

	if body == "" {
		if language == "zh-TW" {
			body = "新的學習卡已經準備好了。"
		} else {
			body = "A quick study card is ready."
		}
	}

	return Message{
		Title:              title,
		Body:               body,
		ActivationArgument: fmt.Sprintf("duolin-gogo://study/%s", card.ID),
	}
}
