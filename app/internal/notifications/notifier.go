package notifications

import (
	"fmt"

	"duolin-gogo/internal/cards"
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
	return BuildStudyMessageForLanguage(card, "en")
}

func BuildStudyMessageForLanguage(card cards.Card, language string) Message {
	title := card.ClickbaitEN
	body := card.QuestionTextEN
	if language == "zh-TW" {
		title = card.ClickbaitZH
		body = card.QuestionTextZH
	}

	if title == "" {
		if language == "zh-TW" {
			title = card.TitleZH
		} else {
			title = card.TitleEN
		}
	}

	if body == "" {
		if language == "zh-TW" {
			body = "新的 Git 小測驗已準備好。"
		} else {
			body = "Quick Git check ready."
		}
	}

	return Message{
		Title:              title,
		Body:               body,
		ActivationArgument: fmt.Sprintf("duolin-gogo://study/%s", card.ID),
	}
}
