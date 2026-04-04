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
	title := card.Clickbait
	if title == "" {
		title = card.Title
	}

	body := card.QuestionText
	if body == "" {
		body = "Quick Git check ready."
	}

	return Message{
		Title:              title,
		Body:               body,
		ActivationArgument: fmt.Sprintf("duolin-gogo://study/%s", card.ID),
	}
}
