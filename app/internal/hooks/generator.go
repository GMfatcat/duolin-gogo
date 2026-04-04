package hooks

import (
	"fmt"
	"strings"

	"duolin-gogo/internal/cards"
)

func Generate(card cards.Card, language string, style string) (string, string) {
	style = normalizeStyle(style)
	primary := primaryTag(card.Tags)
	metaphor := firstNonEmpty(card.MetaphorSeed...)
	confusion := firstNonEmpty(card.ConfusionWith...)

	if language == "zh-TW" {
		return generateZHTitle(card, style, primary, metaphor, confusion), localizedBody(card, language)
	}

	return generateENTitle(card, style, primary, metaphor, confusion), localizedBody(card, language)
}

func generateZHTitle(card cards.Card, style string, primary string, metaphor string, confusion string) string {
	switch style {
	case "safe":
		return fmt.Sprintf("快速確認一下：%s", localizedTitle(card, "zh-TW"))
	case "aggressive":
		if confusion != "" {
			return fmt.Sprintf("大多數人會把它跟 %s 搞混，你呢？", prettyID(confusion))
		}
		return fmt.Sprintf("很多人第一次就把 %s 用錯", localizedTitle(card, "zh-TW"))
	case "chaotic":
		if metaphor != "" {
			return fmt.Sprintf("看起來像「%s」的小動作，往往就是事故開頭", metaphor)
		}
		return fmt.Sprintf("這個 %s 指令看似無害，後果可能很大", primary)
	default:
		if metaphor != "" {
			return fmt.Sprintf("%s，不要急著下一步。你知道這是哪個 Git 動作嗎？", metaphor)
		}
		return fmt.Sprintf("你真的懂 %s 嗎？", localizedTitle(card, "zh-TW"))
	}
}

func generateENTitle(card cards.Card, style string, primary string, metaphor string, confusion string) string {
	switch style {
	case "safe":
		return fmt.Sprintf("Quick check: %s", localizedTitle(card, "en"))
	case "aggressive":
		if confusion != "" {
			return fmt.Sprintf("Most developers confuse this with %s. Do you?", prettyID(confusion))
		}
		return fmt.Sprintf("A lot of developers get %s wrong on the first try", localizedTitle(card, "en"))
	case "chaotic":
		if metaphor != "" {
			return fmt.Sprintf("The move that feels like \"%s\" is where trouble starts", metaphor)
		}
		return fmt.Sprintf("This %s command looks harmless until it isn't", primary)
	default:
		if metaphor != "" {
			return fmt.Sprintf("Think of it as %s before the real move. Know the command?", metaphor)
		}
		return fmt.Sprintf("Do you really understand %s?", localizedTitle(card, "en"))
	}
}

func localizedTitle(card cards.Card, language string) string {
	if language == "zh-TW" {
		return firstNonEmpty(card.TitleZH, card.Title, card.TitleEN)
	}
	return firstNonEmpty(card.TitleEN, card.Title, card.TitleZH)
}

func localizedBody(card cards.Card, language string) string {
	if language == "zh-TW" {
		return firstNonEmpty(card.QuestionTextZH, card.QuestionText, card.QuestionTextEN)
	}
	return firstNonEmpty(card.QuestionTextEN, card.QuestionText, card.QuestionTextZH)
}

func normalizeStyle(style string) string {
	switch strings.TrimSpace(strings.ToLower(style)) {
	case "safe", "playful", "aggressive", "chaotic":
		return strings.ToLower(style)
	default:
		return "playful"
	}
}

func primaryTag(tags []string) string {
	for _, tag := range tags {
		if tag != "" && tag != "git" {
			return tag
		}
	}
	if len(tags) > 0 {
		return tags[0]
	}
	return "git"
}

func prettyID(id string) string {
	id = strings.TrimSpace(id)
	if id == "" {
		return ""
	}
	id = strings.TrimPrefix(id, "git-")
	return strings.ReplaceAll(id, "-", " ")
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			return trimmed
		}
	}
	return ""
}
