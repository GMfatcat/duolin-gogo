//go:build !windows

package notifications

type WindowsSender struct{}

func (WindowsSender) Send(message Message) error {
	return nil
}
