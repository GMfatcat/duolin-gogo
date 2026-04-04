package notifications

import "strings"

func CardIDFromActivationArgument(argument string) (string, bool) {
	const prefix = "duolin-gogo://study/"
	if !strings.HasPrefix(argument, prefix) {
		return "", false
	}

	cardID := strings.TrimPrefix(argument, prefix)
	if cardID == "" {
		return "", false
	}

	return cardID, true
}
