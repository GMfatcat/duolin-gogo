//go:build windows

package notifications

import toast "git.sr.ht/~jackmordaunt/go-toast/v2"

type WindowsSender struct{}

func (WindowsSender) Send(message Message) error {
	notification := toast.Notification{
		AppID:               AppID,
		Title:               message.Title,
		Body:                message.Body,
		Icon:                defaultIconPath(),
		ActivationArguments: message.ActivationArgument,
		ActivationType:      toast.Foreground,
	}

	return notification.Push()
}
