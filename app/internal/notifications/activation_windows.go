//go:build windows

package notifications

import toast "git.sr.ht/~jackmordaunt/go-toast/v2"

func RegisterActivationHandler(onCard func(cardID string)) {
	toast.SetActivationCallback(func(args string, _ []toast.UserData) {
		if cardID, ok := CardIDFromActivationArgument(args); ok {
			onCard(cardID)
		}
	})
}

func ConfigureApp() error {
	return toast.SetAppData(toast.AppData{
		AppID: AppID,
	})
}
