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
		return generateZHTitle(style, title, primary, metaphor, confusion, tags, index), body
	}

	return generateENTitle(style, title, primary, metaphor, confusion, tags, index), body
}

func generateZHTitle(style string, title string, primary string, metaphor string, confusion string, tags map[string]bool, index int) string {
	confusionLabel := prettyID(confusion)

	switch style {
	case "safe":
		options := []string{
			fmt.Sprintf("快速確認一下：你真的懂 %s 嗎？", title),
			fmt.Sprintf("先別急著滑走，%s 你能說清楚嗎？", title),
			fmt.Sprintf("你對 %s 的理解，真的夠穩嗎？", title),
		}
		return pick(options, index)
	case "aggressive":
		if confusionLabel != "" {
			options := []string{
				fmt.Sprintf("很多人把它跟 %s 搞混，你也是嗎？", confusionLabel),
				fmt.Sprintf("只要你還把它當成 %s，用 Git 就很容易翻車", confusionLabel),
				fmt.Sprintf("這題專抓還分不清 %s 的人", confusionLabel),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("很多人第一次就把 %s 用錯", title),
			fmt.Sprintf("%s 看起來簡單，出事前大家都這樣想", title),
			fmt.Sprintf("你以為自己懂 %s，錯誤通常就從這裡開始", title),
		}
		return pick(options, index)
	case "chaotic":
		if metaphor != "" {
			options := []string{
				fmt.Sprintf("看起來像「%s」的小動作，往往就是事故開頭", metaphor),
				fmt.Sprintf("熱門誤用榜又是它：把這題當成「%s」的人超多", metaphor),
				fmt.Sprintf("心測一下：你是會先「%s」的人，還是會直接做錯的人？", metaphor),
				fmt.Sprintf("別被「%s」這種錯覺騙了，很多人就是這樣翻車", metaphor),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("這個 %s 指令看起來無害，實際上很常害人收尾", primary),
			fmt.Sprintf("又一個看似普通的 %s 動作，正在悄悄搞亂工作流", primary),
			fmt.Sprintf("大家都以為這個 %s 很直覺，問題就出在這裡", title),
			fmt.Sprintf("這個看起來很日常的 %s 操作，混亂程度比你想得高", primary),
		}
		return pick(options, index)
	default:
		if metaphor != "" && tags["safer-first"] {
			options := []string{
				fmt.Sprintf("把它想成「%s」再做真正動作，你會更快懂這個 Git 概念", metaphor),
				fmt.Sprintf("如果先用「%s」來理解，這題會突然變簡單", metaphor),
				fmt.Sprintf("先把它當成「%s」，答案通常就浮出來了", metaphor),
			}
			return pick(options, index)
		}
		if confusionLabel != "" && tags["comparison"] {
			options := []string{
				fmt.Sprintf("你真的知道它和 %s 差在哪嗎？", confusionLabel),
				fmt.Sprintf("這題最常被誤認成 %s", confusionLabel),
				fmt.Sprintf("當 %s 出現時，你分得出這題在做什麼嗎？", confusionLabel),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("你真的理解 %s 嗎？", title),
			fmt.Sprintf("這個 %s 概念其實不難，但很多人還是會用錯", primary),
			fmt.Sprintf("你可能記得名字，但 %s 到底在做什麼？", title),
		}
		return pick(options, index)
	}
}

func generateENTitle(style string, title string, primary string, metaphor string, confusion string, tags map[string]bool, index int) string {
	confusionLabel := prettyID(confusion)

	switch style {
	case "safe":
		options := []string{
			fmt.Sprintf("Quick check: do you really understand %s?", title),
			fmt.Sprintf("Before you move on, can you explain %s?", title),
			fmt.Sprintf("Is your mental model for %s actually solid?", title),
		}
		return pick(options, index)
	case "aggressive":
		if confusionLabel != "" {
			options := []string{
				fmt.Sprintf("Most developers confuse this with %s. Do you?", confusionLabel),
				fmt.Sprintf("If you still treat this like %s, trouble is coming", confusionLabel),
				fmt.Sprintf("This one exposes people who still mix it up with %s", confusionLabel),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("A lot of developers get %s wrong on the first try", title),
			fmt.Sprintf("%s looks easy right before it goes wrong", title),
			fmt.Sprintf("You think you know %s. That is usually where mistakes start", title),
		}
		return pick(options, index)
	case "chaotic":
		if metaphor != "" {
			options := []string{
				fmt.Sprintf("The move that feels like \"%s\" is where trouble starts", metaphor),
				fmt.Sprintf("Trending mistake: treating this like \"%s\"", metaphor),
				fmt.Sprintf("Personality quiz: are you the kind of developer who \"%s\"s first?", metaphor),
				fmt.Sprintf("Do not let the \"%s\" vibe fool you. That is how people get burned", metaphor),
			}
			return pick(options, index)
		}
		options := []string{
			fmt.Sprintf("This %s command looks harmless until it isn't", primary),
			fmt.Sprintf("A suspiciously normal %s move is quietly ruining workflows", primary),
			fmt.Sprintf("Everyone thinks this %s step is obvious. That is the trap", title),
			fmt.Sprintf("This boring-looking %s action causes more chaos than people admit", primary),
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
				fmt.Sprintf("When %s shows up, can you still tell this apart?", confusionLabel),
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
