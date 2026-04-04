//go:build !windows

package notifications

func RegisterActivationHandler(onCard func(cardID string)) {}

func ConfigureApp() error {
	return nil
}
