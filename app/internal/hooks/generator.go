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
	title := localizedTitle(card, language)
	body := localizedBody(card, language)
	tags := toSet(card.HookStyleTags)
	index := stableIndex(card.ID + "|" + language + "|" + style)

	if language == "zh-TW" {
		return generateZHTitle(style, title, body, primary, metaphor, confusion, tags, index), body
	}

	return generateENTitle(style, title, body, primary, metaphor, confusion, tags, index), body
}

func generateZHTitle(style string, title string, body string, primary string, metaphor string, confusion string, tags map[string]bool, index int) string {
	confusionLabel := prettyID(confusion)

	switch style {
	case "safe":
		options := []string{
			fmt.Sprintf("快速確認一下：%s", title),
			fmt.Sprintf("這題先別急，先看懂 %s", title),
			fmt.Sprintf("你現在能清楚解釋 %s 嗎？", title),
		}
		return pick(options, index)
	case "aggressive":
		if confusionLabel != "" {
			options := []string{
				fmt.Sprintf("大多數人會把它跟 %s 搞混，你也可能中招", confusionLabel),
				fmt.Sprintf("如果你把它當成 %s，接下來很容易出事", confusionLabel),
				fmt.Sprintf("這題專門抓會把它和 %s 混為一談的人", confusionLabel),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("很多人第一次就把 %s 用錯", title),
			fmt.Sprintf("這個 %s 動作，看起來簡單卻最常翻車", primary),
			fmt.Sprintf("你以為自己會 %s，但多數人其實第一步就錯", title),
		}
		return pick(options, index)
	case "chaotic":
		if metaphor != "" {
			options := []string{
				fmt.Sprintf("看起來像「%s」的小動作，往往就是事故開頭", metaphor),
				fmt.Sprintf("今天最容易讓人手滑的，不是新聞，是這個「%s」操作", metaphor),
				fmt.Sprintf("心測一下：你是會先 %s，還是直接做錯的人？", metaphor),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("這個 %s 指令看似無害，後果可能很大", primary),
			fmt.Sprintf("熱門錯誤操作前幾名，這個 %s 通常都上榜", primary),
			fmt.Sprintf("你以為只是普通的 %s？很多問題都從這裡開始", title),
		}
		return pick(options, index)
	default:
		if metaphor != "" && tags["safer-first"] {
			options := []string{
				fmt.Sprintf("%s，不要急著下一步。你知道這是哪個 Git 動作嗎？", metaphor),
				fmt.Sprintf("先想成「%s」，你就比較不容易把這題做錯", metaphor),
				fmt.Sprintf("把它當成「%s」來理解，突然就會了", metaphor),
			}
			return pick(options, index)
		}
		if confusionLabel != "" && tags["comparison"] {
			options := []string{
				fmt.Sprintf("你真的分得清它和 %s 嗎？", confusionLabel),
				fmt.Sprintf("這題最常被拿來跟 %s 搞混", confusionLabel),
				fmt.Sprintf("看到 %s 時，你知道自己不是在做它嗎？", confusionLabel),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("你真的懂 %s 嗎？", title),
			fmt.Sprintf("這個 %s 觀念，多數人其實沒講清楚", primary),
			fmt.Sprintf("這題不難，難的是你以為自己早就會了：%s", title),
		}
		return pick(options, index)
	}
}

func generateENTitle(style string, title string, body string, primary string, metaphor string, confusion string, tags map[string]bool, index int) string {
	confusionLabel := prettyID(confusion)

	switch style {
	case "safe":
		options := []string{
			fmt.Sprintf("Quick check: %s", title),
			fmt.Sprintf("Before you click ahead, can you explain %s?", title),
			fmt.Sprintf("Do you have a clean mental model for %s?", title),
		}
		return pick(options, index)
	case "aggressive":
		if confusionLabel != "" {
			options := []string{
				fmt.Sprintf("Most developers confuse this with %s. Do you?", confusionLabel),
				fmt.Sprintf("If you treat this like %s, you are setting yourself up for trouble", confusionLabel),
				fmt.Sprintf("This one catches people who still mix it up with %s", confusionLabel),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("A lot of developers get %s wrong on the first try", title),
			fmt.Sprintf("%s looks easy right before it goes wrong", title),
			fmt.Sprintf("You probably think you know %s. That is where mistakes start", title),
		}
		return pick(options, index)
	case "chaotic":
		if metaphor != "" {
			options := []string{
				fmt.Sprintf("The move that feels like \"%s\" is where trouble starts", metaphor),
				fmt.Sprintf("Trending mistake of the day: treating this like \"%s\"", metaphor),
				fmt.Sprintf("Personality quiz: are you the kind of developer who \"%s\"s first?", metaphor),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("This %s command looks harmless until it isn't", primary),
			fmt.Sprintf("A suspiciously normal %s move is ruining a lot of workflows", primary),
			fmt.Sprintf("Everyone thinks this %s step is obvious. That is the trap", title),
		}
		return pick(options, index)
	default:
		if metaphor != "" && tags["safer-first"] {
			options := []string{
				fmt.Sprintf("Think of it as %s before the real move. Know the command?", metaphor),
				fmt.Sprintf("If you frame it as %s, this Git concept gets easier fast", metaphor),
				fmt.Sprintf("Treat this like %s first and the answer becomes obvious", metaphor),
			}
			return pick(options, index)
		}
		if confusionLabel != "" && tags["comparison"] {
			options := []string{
				fmt.Sprintf("Do you actually know how this differs from %s?", confusionLabel),
				fmt.Sprintf("This gets confused with %s more often than it should", confusionLabel),
				fmt.Sprintf("When %s shows up, can you tell this apart from it?", confusionLabel),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("Do you really understand %s?", title),
			fmt.Sprintf("This %s concept is simpler than it looks, but easier to misuse than people admit", primary),
			fmt.Sprintf("You may know the command name. Do you know what %s actually does?", title),
		}
		return pick(options, index)
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

func toSet(values []string) map[string]bool {
	out := map[string]bool{}
	for _, value := range values {
		value = strings.TrimSpace(strings.ToLower(value))
		if value != "" {
			out[value] = true
		}
	}
	return out
}

func stableIndex(seed string) int {
	sum := 0
	for _, r := range seed {
		sum += int(r)
	}
	if sum < 0 {
		sum = -sum
	}
	return sum
}

func pick(options []string, index int) string {
	if len(options) == 0 {
		return ""
	}
	return options[index%len(options)]
}
