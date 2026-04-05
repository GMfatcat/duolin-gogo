package pet

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

const (
	clickCooldownWindow   = 15 * time.Second
	ambientCooldownWindow = 6 * time.Second
	rapidClickWindow      = 12 * time.Second
	easterEggCooldown     = 30 * time.Minute
	welcomeBackWindow     = 8 * time.Hour
)

const (
	StudyEventAnsweredCorrect = "answered_correct"
	StudyEventAnsweredWrong   = "answered_wrong"
	StudyEventLearnBatch      = "learn_batch_complete"
	StudyEventReviewBatch     = "review_batch_complete"
	StudyEventStreak          = "streak_continued"
)

const (
	TriggerClicked        = "clicked"
	TriggerCorrect        = "correct"
	TriggerWrong          = "wrong"
	TriggerLearnBreak     = "learn_break"
	TriggerReviewComplete = "review_complete"
	TriggerReturn         = "return"
)

type State struct {
	BondXP            int     `json:"bond_xp"`
	Stage             int     `json:"stage"`
	ReactionStep      int     `json:"reaction_step,omitempty"`
	LastInteractionAt *string `json:"last_interaction_at,omitempty"`
	LastReactionAt    *string `json:"last_reaction_at,omitempty"`
	LastRapidClickAt  *string `json:"last_rapid_click_at,omitempty"`
	RapidClickCount   int     `json:"rapid_click_count,omitempty"`
	LastWelcomeAt     *string `json:"last_welcome_at,omitempty"`
	LastEasterEggAt   *string `json:"last_easter_egg_at,omitempty"`
}

type Reaction struct {
	Key     string `json:"key"`
	Variant string `json:"variant"`
	Pose    string `json:"pose"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type InteractionResult struct {
	State    State    `json:"state"`
	Reaction Reaction `json:"reaction"`
}

func Load(path string) (State, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return State{}, nil
		}
		return State{}, err
	}

	var state State
	if err := json.Unmarshal(bytes, &state); err != nil {
		return State{}, err
	}
	state.Stage = stageForXP(state.BondXP)
	return state, nil
}

func Save(path string, state State) error {
	state.Stage = stageForXP(state.BondXP)
	bytes, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(bytes, '\n'), 0o644)
}

func RecordStudyEvent(path string, event string, now time.Time) (State, error) {
	state, err := Load(path)
	if err != nil {
		return State{}, err
	}

	state.BondXP += xpForStudyEvent(event)
	state.Stage = stageForXP(state.BondXP)

	if err := Save(path, state); err != nil {
		return State{}, err
	}

	return state, nil
}

func Interact(path string, language string, topic string, now time.Time) (InteractionResult, error) {
	state, err := Load(path)
	if err != nil {
		return InteractionResult{}, err
	}

	state = trackRapidClicks(state, now)

	if shouldTriggerRapidClickEgg(state, now) {
		stamp := now.Format(time.RFC3339)
		state.LastReactionAt = &stamp
		state.LastEasterEggAt = &stamp
		reaction := pickReaction(rapidClickReactions(language, topic, state.Stage), state, "rapid_click")
		state.ReactionStep++
		if err := Save(path, state); err != nil {
			return InteractionResult{}, err
		}
		return InteractionResult{State: state, Reaction: reaction}, nil
	}

	if lastInteractionWithinCooldown(state, now) {
		if err := Save(path, state); err != nil {
			return InteractionResult{}, err
		}
		return InteractionResult{
			State:    state,
			Reaction: cooldownReaction(language),
		}, nil
	}

	state.BondXP += 1
	state.Stage = stageForXP(state.BondXP)
	stamp := now.Format(time.RFC3339)
	wasAwayLongEnough := shouldTriggerWelcomeBack(state, now)
	state.LastInteractionAt = &stamp
	state.LastReactionAt = &stamp
	if wasAwayLongEnough {
		state.LastWelcomeAt = &stamp
		state.LastEasterEggAt = &stamp
	}

	normalizedTopic := normalizeTopic(topic)
	reaction := pickReaction(clickReactions(language, normalizedTopic, state.Stage), state, TriggerClicked)
	if wasAwayLongEnough {
		reaction = pickReaction(welcomeBackReactions(language, normalizedTopic, state.Stage), state, "welcome_back")
	} else if isGeneralTopic(normalizedTopic) && shouldTriggerTimeOfDayEgg(state, now) {
		state.LastEasterEggAt = &stamp
		reaction = pickReaction(timeOfDayReactions(language, now), state, "time_of_day")
	} else if shouldTriggerTopicInsideJoke(state, TriggerClicked, normalizedTopic, now) {
		state.LastEasterEggAt = &stamp
		reaction = pickReaction(topicInsideJokeReactions(language, normalizedTopic, TriggerClicked), state, "topic_inside_joke")
	}

	state.ReactionStep++
	if err := Save(path, state); err != nil {
		return InteractionResult{}, err
	}

	return InteractionResult{State: state, Reaction: reaction}, nil
}

func ReactionForTrigger(path string, trigger string, language string, topic string, now time.Time) (InteractionResult, error) {
	state, err := Load(path)
	if err != nil {
		return InteractionResult{}, err
	}

	if !shouldEmitAmbientReaction(state, trigger, now) {
		return InteractionResult{State: state}, nil
	}

	stamp := now.Format(time.RFC3339)
	state.LastReactionAt = &stamp

	normalizedTopic := normalizeTopic(topic)
	reaction := reactionForTrigger(language, normalizedTopic, trigger, state)
	if shouldTriggerRareCelebration(state, trigger, now) {
		state.LastEasterEggAt = &stamp
		reaction = pickReaction(rareCelebrationReactions(language, normalizedTopic), state, "rare_celebration")
	} else if shouldTriggerTopicInsideJoke(state, trigger, normalizedTopic, now) {
		state.LastEasterEggAt = &stamp
		reaction = pickReaction(topicInsideJokeReactions(language, normalizedTopic, trigger), state, "topic_inside_joke")
	}

	state.ReactionStep++
	if err := Save(path, state); err != nil {
		return InteractionResult{}, err
	}

	return InteractionResult{State: state, Reaction: reaction}, nil
}

func xpForStudyEvent(event string) int {
	switch event {
	case StudyEventAnsweredCorrect:
		return 2
	case StudyEventAnsweredWrong:
		return 1
	case StudyEventLearnBatch:
		return 3
	case StudyEventReviewBatch:
		return 4
	case StudyEventStreak:
		return 2
	default:
		return 0
	}
}

func stageForXP(bondXP int) int {
	switch {
	case bondXP >= 16:
		return 2
	case bondXP >= 6:
		return 1
	default:
		return 0
	}
}

func shouldEmitAmbientReaction(state State, trigger string, now time.Time) bool {
	switch trigger {
	case TriggerLearnBreak, TriggerReviewComplete, TriggerReturn:
		return true
	}

	if lastReactionWithinCooldown(state, now) {
		return false
	}

	seed := state.ReactionStep + state.BondXP + len(trigger) + now.Minute()
	switch trigger {
	case TriggerCorrect:
		return seed%4 != 0
	case TriggerWrong:
		return seed%3 == 0 || seed%3 == 2
	default:
		return true
	}
}

func reactionForTrigger(language string, topic string, trigger string, state State) Reaction {
	switch trigger {
	case TriggerCorrect:
		return pickReaction(correctReactions(language, topic, state.Stage), state, trigger)
	case TriggerWrong:
		return pickReaction(wrongReactions(language, topic, state.Stage), state, trigger)
	case TriggerLearnBreak:
		return pickReaction(learnBreakReactions(language, topic, state.Stage), state, trigger)
	case TriggerReviewComplete:
		return pickReaction(reviewCompleteReactions(language, topic, state.Stage), state, trigger)
	case TriggerReturn:
		return pickReaction(returnReactions(language, topic, state.Stage), state, trigger)
	default:
		return pickReaction(clickReactions(language, topic, state.Stage), state, trigger)
	}
}

func pickReaction(pool []Reaction, state State, trigger string) Reaction {
	if len(pool) == 0 {
		return Reaction{Key: "fallback", Variant: "neutral", Pose: "idle", Title: "DG", Body: "..."}
	}
	index := (state.ReactionStep + len(trigger)) % len(pool)
	return pool[index]
}

func cooldownReaction(language string) Reaction {
	if language == "zh-TW" {
		return Reaction{Key: "cooldown", Variant: "focus", Pose: "rest", Title: "DG", Body: "我有聽到，先讓我緩一下。"}
	}
	return Reaction{Key: "cooldown", Variant: "focus", Pose: "rest", Title: "DG", Body: "I heard you. Give me a beat."}
}

func clickReactions(language string, topic string, stage int) []Reaction {
	if topicPool := topicClickReactions(language, topic); len(topicPool) > 0 {
		return topicPool
	}

	switch {
	case stage >= 2:
		if language == "zh-TW" {
			return []Reaction{
				{Key: "stage_two_click_warm", Variant: "celebration", Pose: "spark", Title: "DG", Body: "你回來了，我開始記住你的學習節奏。"},
				{Key: "stage_two_click_sync", Variant: "celebration", Pose: "spark", Title: "DG", Body: "我們現在滿同步的，這輪可以走得很順。"},
				{Key: "stage_two_click_grin", Variant: "celebration", Pose: "spark", Title: "DG", Body: "又來找我了，這種回鍋速度我很喜歡。"},
				{Key: "stage_two_click_anchor", Variant: "celebration", Pose: "spark", Title: "DG", Body: "好，這輪我幫你把專注感先定住。"},
				{Key: "stage_two_click_companion", Variant: "celebration", Pose: "spark", Title: "DG", Body: "行，我跟上了，這次陪你把這輪收漂亮。"},
			}
		}
		return []Reaction{
			{Key: "stage_two_click_warm", Variant: "celebration", Pose: "spark", Title: "DG", Body: "You are back. I am starting to learn your rhythm."},
			{Key: "stage_two_click_sync", Variant: "celebration", Pose: "spark", Title: "DG", Body: "I am in sync now. You can keep the pace up."},
			{Key: "stage_two_click_grin", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Back again already? I like this kind of return speed."},
			{Key: "stage_two_click_anchor", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Alright, let me pin the focus down before we move."},
			{Key: "stage_two_click_companion", Variant: "celebration", Pose: "spark", Title: "DG", Body: "I am with you. Let us make this batch look clean."},
		}
	case stage >= 1:
		if language == "zh-TW" {
			return []Reaction{
				{Key: "stage_one_click_focus", Variant: "focus", Pose: "wave", Title: "DG", Body: "好，這輪我們一起走完。"},
				{Key: "stage_one_click_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "我準備好了，你先踏出第一步。"},
				{Key: "stage_one_click_steady", Variant: "focus", Pose: "wave", Title: "DG", Body: "慢慢來就好，這輪不用衝。"},
				{Key: "stage_one_click_companion", Variant: "focus", Pose: "wave", Title: "DG", Body: "你負責選答案，我負責把節奏顧好。"},
				{Key: "stage_one_click_reset", Variant: "focus", Pose: "wave", Title: "DG", Body: "來，先把注意力拉回這裡。"},
			}
		}
		return []Reaction{
			{Key: "stage_one_click_focus", Variant: "focus", Pose: "wave", Title: "DG", Body: "Alright, let us work through this batch together."},
			{Key: "stage_one_click_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "I am ready. You take the first step."},
			{Key: "stage_one_click_steady", Variant: "focus", Pose: "wave", Title: "DG", Body: "No rush. This round can stay steady."},
			{Key: "stage_one_click_companion", Variant: "focus", Pose: "wave", Title: "DG", Body: "You pick the answers. I will hold the rhythm."},
			{Key: "stage_one_click_reset", Variant: "focus", Pose: "wave", Title: "DG", Body: "Come on, let us pull the focus back in."},
		}
	default:
		if language == "zh-TW" {
			return []Reaction{
				{Key: "stage_zero_click_intro", Variant: "neutral", Pose: "idle", Title: "DG", Body: "我在這裡，多點我幾次我就會更熟。"},
				{Key: "stage_zero_click_warmup", Variant: "neutral", Pose: "idle", Title: "DG", Body: "再互動一點點，我會醒得更快。"},
				{Key: "stage_zero_click_small", Variant: "neutral", Pose: "idle", Title: "DG", Body: "好，先從一小步開始就行。"},
				{Key: "stage_zero_click_glad", Variant: "neutral", Pose: "idle", Title: "DG", Body: "你有回來就很好，今天先慢慢熱開。"},
				{Key: "stage_zero_click_soft", Variant: "neutral", Pose: "idle", Title: "DG", Body: "不用急著變強，先把今天這一張看懂。"},
			}
		}
		return []Reaction{
			{Key: "stage_zero_click_intro", Variant: "neutral", Pose: "idle", Title: "DG", Body: "I am here. Keep tapping in and I will warm up."},
			{Key: "stage_zero_click_warmup", Variant: "neutral", Pose: "idle", Title: "DG", Body: "Tap back in a little more and I will wake up faster."},
			{Key: "stage_zero_click_small", Variant: "neutral", Pose: "idle", Title: "DG", Body: "Alright, one small step is enough to start."},
			{Key: "stage_zero_click_glad", Variant: "neutral", Pose: "idle", Title: "DG", Body: "You showed up. That is already a good start."},
			{Key: "stage_zero_click_soft", Variant: "neutral", Pose: "idle", Title: "DG", Body: "No need to become brilliant all at once. One clean card is fine."},
		}
	}
}

func correctReactions(language string, topic string, stage int) []Reaction {
	if topicPool := topicCorrectReactions(language, topic); len(topicPool) > 0 {
		return topicPool
	}

	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "correct_stage_one_clean", Variant: "celebration", Pose: "nod", Title: "DG", Body: "這題抓得很乾淨，感覺開始穩了。"},
				{Key: "correct_stage_one_locking", Variant: "celebration", Pose: "nod", Title: "DG", Body: "對，就是這種感覺，把它帶去下一張。"},
				{Key: "correct_stage_one_steady", Variant: "celebration", Pose: "nod", Title: "DG", Body: "這個點你有抓住，今天的手感有起來。"},
			}
		}
		return []Reaction{
			{Key: "correct_stage_one_clean", Variant: "celebration", Pose: "nod", Title: "DG", Body: "That was clean. I can tell this is starting to stick."},
			{Key: "correct_stage_one_locking", Variant: "celebration", Pose: "nod", Title: "DG", Body: "Yes, that is the feeling. Keep it for the next card."},
			{Key: "correct_stage_one_steady", Variant: "celebration", Pose: "nod", Title: "DG", Body: "You really caught that one. The rhythm is getting steadier."},
		}
	}

	if language == "zh-TW" {
		return []Reaction{
			{Key: "correct_stage_zero_nice", Variant: "celebration", Pose: "nod", Title: "DG", Body: "不錯，先把這個手感留住。"},
			{Key: "correct_stage_zero_hold", Variant: "celebration", Pose: "nod", Title: "DG", Body: "好球，把同樣的節奏帶進下一題。"},
			{Key: "correct_stage_zero_small", Variant: "celebration", Pose: "nod", Title: "DG", Body: "這種小勝利很重要，先收好。"},
		}
	}
	return []Reaction{
		{Key: "correct_stage_zero_nice", Variant: "celebration", Pose: "nod", Title: "DG", Body: "Nice hit. Hold on to that feeling for the next one."},
		{Key: "correct_stage_zero_hold", Variant: "celebration", Pose: "nod", Title: "DG", Body: "Good catch. Bring that same energy into the next card."},
		{Key: "correct_stage_zero_small", Variant: "celebration", Pose: "nod", Title: "DG", Body: "Small wins count. Keep this one."},
	}
}

func wrongReactions(language string, topic string, stage int) []Reaction {
	if topicPool := topicWrongReactions(language, topic); len(topicPool) > 0 {
		return topicPool
	}

	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "wrong_stage_one_almost", Variant: "warning", Pose: "think", Title: "DG", Body: "沒關係，這種差一點的錯很值得記住。"},
				{Key: "wrong_stage_one_keep", Variant: "warning", Pose: "think", Title: "DG", Body: "先盯住那個差異，下次會更穩。"},
				{Key: "wrong_stage_one_reset", Variant: "warning", Pose: "think", Title: "DG", Body: "這題不用急著扳回來，先把關鍵差異看清楚。"},
			}
		}
		return []Reaction{
			{Key: "wrong_stage_one_almost", Variant: "warning", Pose: "think", Title: "DG", Body: "That is okay. These almost-right misses are worth keeping."},
			{Key: "wrong_stage_one_keep", Variant: "warning", Pose: "think", Title: "DG", Body: "Keep the difference in view. The next pass will feel steadier."},
			{Key: "wrong_stage_one_reset", Variant: "warning", Pose: "think", Title: "DG", Body: "No need to win this back instantly. Just keep the key difference visible."},
		}
	}

	if language == "zh-TW" {
		return []Reaction{
			{Key: "wrong_stage_zero_difference", Variant: "warning", Pose: "think", Title: "DG", Body: "先記住差異就好，下一輪會更穩。"},
			{Key: "wrong_stage_zero_retry", Variant: "warning", Pose: "think", Title: "DG", Body: "這題先不用急，我們等一下再繞回來。"},
			{Key: "wrong_stage_zero_soft", Variant: "warning", Pose: "think", Title: "DG", Body: "今天先把這種差異看熟就很夠了。"},
		}
	}
	return []Reaction{
		{Key: "wrong_stage_zero_difference", Variant: "warning", Pose: "think", Title: "DG", Body: "Just hold on to the difference. The next pass will feel steadier."},
		{Key: "wrong_stage_zero_retry", Variant: "warning", Pose: "think", Title: "DG", Body: "Do not worry about this one yet. We can loop back cleanly."},
		{Key: "wrong_stage_zero_soft", Variant: "warning", Pose: "think", Title: "DG", Body: "Getting used to this difference is already enough for today."},
	}
}

func learnBreakReactions(language string, topic string, stage int) []Reaction {
	if topicPool := topicLearnBreakReactions(language, topic); len(topicPool) > 0 {
		return topicPool
	}

	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "learn_break_stage_one_land", Variant: "focus", Pose: "rest", Title: "DG", Body: "這輪落得不錯，先讓腦袋有點空間。"},
				{Key: "learn_break_stage_one_room", Variant: "focus", Pose: "rest", Title: "DG", Body: "短暫停一下剛好，讓剛剛那幾張沉一下。"},
				{Key: "learn_break_stage_one_soft", Variant: "focus", Pose: "rest", Title: "DG", Body: "先別急著補下一張，這幾個點正在落地。"},
			}
		}
		return []Reaction{
			{Key: "learn_break_stage_one_land", Variant: "focus", Pose: "rest", Title: "DG", Body: "That batch landed well. Give your brain a little room now."},
			{Key: "learn_break_stage_one_room", Variant: "focus", Pose: "rest", Title: "DG", Body: "A short pause is right. Let the last few cards settle."},
			{Key: "learn_break_stage_one_soft", Variant: "focus", Pose: "rest", Title: "DG", Body: "No need to rush the next card. These ideas are still settling in."},
		}
	}

	if language == "zh-TW" {
		return []Reaction{
			{Key: "learn_break_stage_zero_wait", Variant: "focus", Pose: "rest", Title: "DG", Body: "先休息一下，下一輪可以等你。"},
			{Key: "learn_break_stage_zero_pause", Variant: "focus", Pose: "rest", Title: "DG", Body: "在這裡停一下剛好，下一輪不會跑掉。"},
			{Key: "learn_break_stage_zero_breathe", Variant: "focus", Pose: "rest", Title: "DG", Body: "先喘口氣，這輪已經有收到了。"},
		}
	}
	return []Reaction{
		{Key: "learn_break_stage_zero_wait", Variant: "focus", Pose: "rest", Title: "DG", Body: "Take a short beat. The next batch can wait."},
		{Key: "learn_break_stage_zero_pause", Variant: "focus", Pose: "rest", Title: "DG", Body: "Pause here for a moment. The next round is fine waiting."},
		{Key: "learn_break_stage_zero_breathe", Variant: "focus", Pose: "rest", Title: "DG", Body: "Take a breath. This round already landed."},
	}
}

func reviewCompleteReactions(language string, topic string, stage int) []Reaction {
	if topicPool := topicReviewCompleteReactions(language, topic); len(topicPool) > 0 {
		return topicPool
	}

	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "review_complete_stage_one_closed", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪複習收得很好，我感覺整個循環開始穩了。"},
				{Key: "review_complete_stage_one_settle", Variant: "celebration", Pose: "spark", Title: "DG", Body: "收尾漂亮，先讓這輪複習沉一下。"},
				{Key: "review_complete_stage_one_glow", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這種收尾值得亮一下，今天這輪真的不錯。"},
			}
		}
		return []Reaction{
			{Key: "review_complete_stage_one_closed", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That review batch closed out nicely. I can feel the loop settling in."},
			{Key: "review_complete_stage_one_settle", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Nice finish. Let that review loop settle in a bit."},
			{Key: "review_complete_stage_one_glow", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That finish deserves a little glow. This round was good."},
		}
	}

	if language == "zh-TW" {
		return []Reaction{
			{Key: "review_complete_stage_zero_done", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪複習完成了，先讓成果沉一下。"},
			{Key: "review_complete_stage_zero_breathe", Variant: "celebration", Pose: "spark", Title: "DG", Body: "複習告一段落，先喘口氣再往下。"},
			{Key: "review_complete_stage_zero_soft", Variant: "celebration", Pose: "spark", Title: "DG", Body: "好，這輪收好了，等等再決定要不要繼續。"},
		}
	}
	return []Reaction{
		{Key: "review_complete_stage_zero_done", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That review batch is done. Take a moment and let it settle."},
		{Key: "review_complete_stage_zero_breathe", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Review complete. Take a breath before you move on."},
		{Key: "review_complete_stage_zero_soft", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Alright, that round is tucked away. Decide on the next move after a beat."},
	}
}

func returnReactions(language string, topic string, stage int) []Reaction {
	if topicPool := topicReturnReactions(language, topic); len(topicPool) > 0 {
		return topicPool
	}

	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "return_stage_one_pickup", Variant: "focus", Pose: "wave", Title: "DG", Body: "你回來了，我們可以從這裡把線接回去。"},
				{Key: "return_stage_one_thread", Variant: "focus", Pose: "wave", Title: "DG", Body: "剛剛那條線還在，我們可以從這裡繼續。"},
				{Key: "return_stage_one_resume", Variant: "focus", Pose: "wave", Title: "DG", Body: "好，重新接回來，不需要從頭暖機。"},
			}
		}
		return []Reaction{
			{Key: "return_stage_one_pickup", Variant: "focus", Pose: "wave", Title: "DG", Body: "You are back. We can pick up the thread from here."},
			{Key: "return_stage_one_thread", Variant: "focus", Pose: "wave", Title: "DG", Body: "That thread is still here. We can keep going now."},
			{Key: "return_stage_one_resume", Variant: "focus", Pose: "wave", Title: "DG", Body: "Good. We can reconnect without warming up from scratch."},
		}
	}

	if language == "zh-TW" {
		return []Reaction{
			{Key: "return_stage_zero_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "好，下一輪已經準備好了。"},
			{Key: "return_stage_zero_resume", Variant: "focus", Pose: "wave", Title: "DG", Body: "可以了，我們從這裡重新開始。"},
			{Key: "return_stage_zero_soft", Variant: "focus", Pose: "wave", Title: "DG", Body: "你回來了，慢慢接上就好。"},
		}
	}
	return []Reaction{
		{Key: "return_stage_zero_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "Alright, the next round is ready."},
		{Key: "return_stage_zero_resume", Variant: "focus", Pose: "wave", Title: "DG", Body: "Okay, we can start fresh from here."},
		{Key: "return_stage_zero_soft", Variant: "focus", Pose: "wave", Title: "DG", Body: "You are back. We can ease into this again."},
	}
}

func rapidClickReactions(language string, topic string, stage int) []Reaction {
	switch normalizeTopic(topic) {
	case "docker":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "rapid_click_docker_ports", Variant: "celebration", Pose: "spark", Title: "DG", Body: "別狂戳啦，我還沒幫你把 port 全都 expose 出來。"},
				{Key: "rapid_click_docker_stack", Variant: "celebration", Pose: "spark", Title: "DG", Body: "你這個節奏像 docker compose 剛一起全部拉起來。"},
			}
		}
		return []Reaction{
			{Key: "rapid_click_docker_ports", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Easy there. I have not exposed every port yet."},
			{Key: "rapid_click_docker_stack", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That pace feels like a whole compose stack starting at once."},
		}
	case "sql":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "rapid_click_sql_where", Variant: "celebration", Pose: "spark", Title: "DG", Body: "你這連點像忘了加 where，一下就掃過整張表。"},
				{Key: "rapid_click_sql_limit", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這個節奏該加個 limit，不然我都要被掃完整輪了。"},
			}
		}
		return []Reaction{
			{Key: "rapid_click_sql_where", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That tapping pace feels like a query with no where clause."},
			{Key: "rapid_click_sql_limit", Variant: "celebration", Pose: "spark", Title: "DG", Body: "This pace could use a limit clause before it scans the whole room."},
		}
	}

	if stage >= 2 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "rapid_click_stage_two_grin", Variant: "celebration", Pose: "spark", Title: "DG", Body: "好啦好啦，我知道你在戳我。這已經算熟客待遇了。"},
				{Key: "rapid_click_stage_two_playful", Variant: "celebration", Pose: "spark", Title: "DG", Body: "再戳下去，我要懷疑你是來摸魚不是來學習。"},
				{Key: "rapid_click_stage_two_boop", Variant: "celebration", Pose: "spark", Title: "DG", Body: "你這個 boop 的頻率，已經很像在叫我加班。"},
			}
		}
		return []Reaction{
			{Key: "rapid_click_stage_two_grin", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Alright, alright. I know you are poking me on purpose now."},
			{Key: "rapid_click_stage_two_playful", Variant: "celebration", Pose: "spark", Title: "DG", Body: "At this point I have to ask if you are studying or just booping me."},
			{Key: "rapid_click_stage_two_boop", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That boop rate is starting to feel like unpaid overtime."},
		}
	}

	if language == "zh-TW" {
		return []Reaction{
			{Key: "rapid_click_stage_zero_notice", Variant: "celebration", Pose: "spark", Title: "DG", Body: "欸，我有在這裡，不用一口氣戳這麼多下。"},
			{Key: "rapid_click_stage_zero_tickle", Variant: "celebration", Pose: "spark", Title: "DG", Body: "好啦我醒了，這種連點很像在搔癢。"},
			{Key: "rapid_click_stage_zero_panic", Variant: "celebration", Pose: "spark", Title: "DG", Body: "你這個節奏有點像 fire drill，我先跟上。"},
		}
	}
	return []Reaction{
		{Key: "rapid_click_stage_zero_notice", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Hey, I am here. You do not need to tap that many times."},
		{Key: "rapid_click_stage_zero_tickle", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Alright, I am awake. That kind of rapid tapping tickles."},
		{Key: "rapid_click_stage_zero_panic", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That pace feels a little like a fire drill, but I am with you."},
	}
}

func welcomeBackReactions(language string, topic string, stage int) []Reaction {
	switch normalizeTopic(topic) {
	case "git":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "welcome_back_git", Variant: "focus", Pose: "wave", Title: "DG", Body: "你回來了。Git 歷史還沒亂掉，我們可以從這裡接回去。"},
				{Key: "welcome_back_git_rebase", Variant: "focus", Pose: "wave", Title: "DG", Body: "歡迎回來，這次我們應該不用把整段歷史 rebase 重來吧。"},
			}
		}
		return []Reaction{
			{Key: "welcome_back_git", Variant: "focus", Pose: "wave", Title: "DG", Body: "You are back. The git history is still clean enough for us to resume."},
			{Key: "welcome_back_git_rebase", Variant: "focus", Pose: "wave", Title: "DG", Body: "Welcome back. Hopefully we do not need to rebase the whole mood today."},
		}
	case "http":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "welcome_back_http", Variant: "focus", Pose: "wave", Title: "DG", Body: "歡迎回來，這次我們先別急著怪 server。"},
				{Key: "welcome_back_http_status", Variant: "focus", Pose: "wave", Title: "DG", Body: "你回來了，先看 status code，再決定誰要背鍋。"},
			}
		}
		return []Reaction{
			{Key: "welcome_back_http", Variant: "focus", Pose: "wave", Title: "DG", Body: "Welcome back. This time we can avoid blaming the server too early."},
			{Key: "welcome_back_http_status", Variant: "focus", Pose: "wave", Title: "DG", Body: "You are back. Let us check the status code before assigning blame."},
		}
	}

	if stage >= 2 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "welcome_back_stage_two_warm", Variant: "focus", Pose: "wave", Title: "DG", Body: "你回來了。我還記得我們剛剛停在哪裡。"},
				{Key: "welcome_back_stage_two_soft", Variant: "focus", Pose: "wave", Title: "DG", Body: "回來得剛好，我把節奏幫你留著了。"},
				{Key: "welcome_back_stage_two_familiar", Variant: "focus", Pose: "wave", Title: "DG", Body: "你這種回來的時間點，我開始有點熟了。"},
			}
		}
		return []Reaction{
			{Key: "welcome_back_stage_two_warm", Variant: "focus", Pose: "wave", Title: "DG", Body: "You are back. I still remember where we left off."},
			{Key: "welcome_back_stage_two_soft", Variant: "focus", Pose: "wave", Title: "DG", Body: "Good timing. I kept the rhythm warm for you."},
			{Key: "welcome_back_stage_two_familiar", Variant: "focus", Pose: "wave", Title: "DG", Body: "I am starting to recognize this exact kind of return."},
		}
	}

	if language == "zh-TW" {
		return []Reaction{
			{Key: "welcome_back_stage_zero_simple", Variant: "focus", Pose: "wave", Title: "DG", Body: "歡迎回來，我們可以慢慢把這輪接上。"},
			{Key: "welcome_back_stage_zero_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "你回來了，下一張卡已經準備好了。"},
			{Key: "welcome_back_stage_zero_soft", Variant: "focus", Pose: "wave", Title: "DG", Body: "沒關係，慢慢接回來就好。"},
		}
	}
	return []Reaction{
		{Key: "welcome_back_stage_zero_simple", Variant: "focus", Pose: "wave", Title: "DG", Body: "Welcome back. We can ease into the next round."},
		{Key: "welcome_back_stage_zero_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "You are back. The next card is ready when you are."},
		{Key: "welcome_back_stage_zero_soft", Variant: "focus", Pose: "wave", Title: "DG", Body: "No problem. We can reconnect gently."},
	}
}

func timeOfDayReactions(language string, now time.Time) []Reaction {
	hour := now.Hour()
	switch {
	case hour < 6:
		if language == "zh-TW" {
			return []Reaction{{Key: "time_of_day_night_owl", Variant: "focus", Pose: "rest", Title: "DG", Body: "這時間你還在，今天真的是夜貓子模式。"}}
		}
		return []Reaction{{Key: "time_of_day_night_owl", Variant: "focus", Pose: "rest", Title: "DG", Body: "You are really in night-owl mode right now."}}
	case hour < 11:
		if language == "zh-TW" {
			return []Reaction{{Key: "time_of_day_morning", Variant: "focus", Pose: "wave", Title: "DG", Body: "早上的腦袋比較新鮮，這時候抓觀念很划算。"}}
		}
		return []Reaction{{Key: "time_of_day_morning", Variant: "focus", Pose: "wave", Title: "DG", Body: "Morning brain is sharp. This is a good time to lock concepts in."}}
	case hour < 17:
		if language == "zh-TW" {
			return []Reaction{{Key: "time_of_day_daylight", Variant: "focus", Pose: "wave", Title: "DG", Body: "白天這輪很適合把卡片一張張收乾淨。"}}
		}
		return []Reaction{{Key: "time_of_day_daylight", Variant: "focus", Pose: "wave", Title: "DG", Body: "Daylight rounds are good for keeping things crisp and clean."}}
	default:
		if language == "zh-TW" {
			return []Reaction{{Key: "time_of_day_evening", Variant: "focus", Pose: "rest", Title: "DG", Body: "晚上就用穩穩的節奏來，不用硬衝。"}}
		}
		return []Reaction{{Key: "time_of_day_evening", Variant: "focus", Pose: "rest", Title: "DG", Body: "Evening pace is fine. No need to force speed now."}}
	}
}

func topicInsideJokeReactions(language string, topic string, trigger string) []Reaction {
	switch normalizeTopic(topic) {
	case "git":
		if language == "zh-TW" {
			return []Reaction{{Key: "inside_joke_git_force_push", Variant: "celebration", Pose: "think", Title: "DG", Body: "先別急著 force push，我們還沒走到那一步。"}}
		}
		return []Reaction{{Key: "inside_joke_git_force_push", Variant: "celebration", Pose: "think", Title: "DG", Body: "Easy there. We are nowhere near needing a force push yet."}}
	case "docker":
		if language == "zh-TW" {
			return []Reaction{{Key: "inside_joke_docker_port", Variant: "celebration", Pose: "think", Title: "DG", Body: "這次真的先確認一下，不要又是 port 沒開。"}}
		}
		return []Reaction{{Key: "inside_joke_docker_port", Variant: "celebration", Pose: "think", Title: "DG", Body: "This might not be broken. It might still just be the port mapping again."}}
	case "sql":
		if language == "zh-TW" {
			return []Reaction{{Key: "inside_joke_sql_where", Variant: "celebration", Pose: "think", Title: "DG", Body: "我先幫你問一句：這次 where 有寫對吧？"}}
		}
		return []Reaction{{Key: "inside_joke_sql_where", Variant: "celebration", Pose: "think", Title: "DG", Body: "Quick check: this time we definitely remembered the where clause, right?"}}
	case "http":
		if language == "zh-TW" {
			return []Reaction{{Key: "inside_joke_http_status", Variant: "celebration", Pose: "think", Title: "DG", Body: "先看 status code，再決定要不要怪 server。"}}
		}
		return []Reaction{{Key: "inside_joke_http_status", Variant: "celebration", Pose: "think", Title: "DG", Body: "Before blaming the server, let us at least check the status code."}}
	}
	return nil
}

func rareCelebrationReactions(language string, topic string) []Reaction {
	switch normalizeTopic(topic) {
	case "languages":
		if language == "zh-TW" {
			return []Reaction{{Key: "rare_celebration_languages", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪語感長得很漂亮，今天的節奏真的不錯。"}}
		}
		return []Reaction{{Key: "rare_celebration_languages", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That one landed beautifully. The language instinct is really showing today."}}
	case "backend-tools":
		if language == "zh-TW" {
			return []Reaction{{Key: "rare_celebration_backend", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪後端工具題收得很順，像把整個工作台整理好了。"}}
		}
		return []Reaction{{Key: "rare_celebration_backend", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That backend-tools finish was crisp. It feels like the whole bench just got organized."}}
	}

	if language == "zh-TW" {
		return []Reaction{
			{Key: "rare_celebration_general", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪收得特別漂亮，我偷偷幫你記一筆。"},
			{Key: "rare_celebration_general_soft", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這種收尾值得亮一下，今天的手感很可以。"},
		}
	}
	return []Reaction{
		{Key: "rare_celebration_general", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That finish was unusually clean. I am quietly filing this one away."},
		{Key: "rare_celebration_general_soft", Variant: "celebration", Pose: "spark", Title: "DG", Body: "This one deserves a little extra glow. The rhythm is good today."},
	}
}

func topicClickReactions(language string, topic string) []Reaction {
	switch normalizeTopic(topic) {
	case "docker":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_docker_click_stack", Variant: "focus", Pose: "wave", Title: "DG", Body: "Docker 模式已開，我們把這堆容器顧整齊。"},
				{Key: "topic_docker_click_watch", Variant: "focus", Pose: "wave", Title: "DG", Body: "回到 docker 地帶了，我會幫你盯著這些移動零件。"},
				{Key: "topic_docker_click_ports", Variant: "focus", Pose: "wave", Title: "DG", Body: "這輪 docker 題先盯住 port、volume 跟 exec 這幾個常出事的點。"},
			}
		}
		return []Reaction{
			{Key: "topic_docker_click_stack", Variant: "focus", Pose: "wave", Title: "DG", Body: "Docker mode is on. Let us keep this stack tidy."},
			{Key: "topic_docker_click_watch", Variant: "focus", Pose: "wave", Title: "DG", Body: "Back in docker land. I am watching the moving parts with you."},
			{Key: "topic_docker_click_ports", Variant: "focus", Pose: "wave", Title: "DG", Body: "Docker round. Let us keep an eye on ports, volumes, and exec flows."},
		}
	case "languages":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_languages_click_tune", Variant: "focus", Pose: "wave", Title: "DG", Body: "這輪偏語言題，我們先把細節音準調好。"},
				{Key: "topic_languages_click_rhythm", Variant: "focus", Pose: "wave", Title: "DG", Body: "語言模式開了，這裡比的是語感跟細節。"},
				{Key: "topic_languages_click_shape", Variant: "focus", Pose: "wave", Title: "DG", Body: "語言題常差在細節形狀，這輪慢一點看會更準。"},
			}
		}
		return []Reaction{
			{Key: "topic_languages_click_tune", Variant: "focus", Pose: "wave", Title: "DG", Body: "This round leans language-heavy. Let us tune the details."},
			{Key: "topic_languages_click_rhythm", Variant: "focus", Pose: "wave", Title: "DG", Body: "Language mode is on. The rhythm here is more about nuance."},
			{Key: "topic_languages_click_shape", Variant: "focus", Pose: "wave", Title: "DG", Body: "Language rounds often hinge on tiny differences. Slow eyes help here."},
		}
	case "git":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_git_click_history", Variant: "focus", Pose: "wave", Title: "DG", Body: "回到 git 了，我們先把歷史線維持乾淨。"},
				{Key: "topic_git_click_branch", Variant: "focus", Pose: "wave", Title: "DG", Body: "這輪 git 題先別急，branch 跟 rebase 的小差異很容易滑掉。"},
			}
		}
		return []Reaction{
			{Key: "topic_git_click_history", Variant: "focus", Pose: "wave", Title: "DG", Body: "Back to git. Let us keep the history line clean."},
			{Key: "topic_git_click_branch", Variant: "focus", Pose: "wave", Title: "DG", Body: "Git round. Tiny branch and rebase differences are where people slip."},
		}
	default:
		return nil
	}
}

func topicCorrectReactions(language string, topic string) []Reaction {
	switch normalizeTopic(topic) {
	case "languages":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_languages_correct_instinct", Variant: "celebration", Pose: "nod", Title: "DG", Body: "這題抓得漂亮，你的語感開始站穩了。"},
				{Key: "topic_languages_correct_click", Variant: "celebration", Pose: "nod", Title: "DG", Body: "很乾淨，語言這側開始有手感了。"},
			}
		}
		return []Reaction{
			{Key: "topic_languages_correct_instinct", Variant: "celebration", Pose: "nod", Title: "DG", Body: "Nice catch. Your language instincts are settling in."},
			{Key: "topic_languages_correct_click", Variant: "celebration", Pose: "nod", Title: "DG", Body: "That was clean. The language side is starting to click."},
		}
	default:
		return nil
	}
}

func topicWrongReactions(language string, topic string) []Reaction {
	switch normalizeTopic(topic) {
	case "docker":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_docker_wrong_layers", Variant: "warning", Pose: "think", Title: "DG", Body: "這題像 container layer 一樣滑掉了，一個細節就整個讀偏。"},
			}
		}
		return []Reaction{
			{Key: "topic_docker_wrong_layers", Variant: "warning", Pose: "think", Title: "DG", Body: "That one slipped like a container layer. One detail changed the whole read."},
		}
	case "git":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_git_wrong_branch", Variant: "warning", Pose: "think", Title: "DG", Body: "這感覺像切到錯的 branch，小差異但後果很大。"},
			}
		}
		return []Reaction{
			{Key: "topic_git_wrong_branch", Variant: "warning", Pose: "think", Title: "DG", Body: "That felt like switching to the wrong branch. Small difference, big consequence."},
		}
	default:
		return nil
	}
}

func topicLearnBreakReactions(language string, topic string) []Reaction {
	switch normalizeTopic(topic) {
	case "backend-tools":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_backend_break", Variant: "focus", Pose: "rest", Title: "DG", Body: "工具這側的手感先讓它沉一下。"},
			}
		}
		return []Reaction{
			{Key: "topic_backend_break", Variant: "focus", Pose: "rest", Title: "DG", Body: "Let the tool-side muscle memory settle for a moment."},
		}
	default:
		return nil
	}
}

func topicReviewCompleteReactions(language string, topic string) []Reaction {
	switch normalizeTopic(topic) {
	case "languages":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_languages_review_complete", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪語言複習收得不錯，語感正在長出來。"},
			}
		}
		return []Reaction{
			{Key: "topic_languages_review_complete", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That language review closed well. The instincts are growing."},
		}
	case "backend-tools":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_backend_review_complete", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪 backend-tools 收得漂亮，操作感越來越熟。"},
			}
		}
		return []Reaction{
			{Key: "topic_backend_review_complete", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Nice backend-tools finish. The operational feel is sticking."},
		}
	default:
		return nil
	}
}

func topicReturnReactions(language string, topic string) []Reaction {
	switch normalizeTopic(topic) {
	case "docker":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_docker_return_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "Docker 又上場了，我們可以把這輪乾淨地跑起來。"},
				{Key: "topic_docker_return_control", Variant: "focus", Pose: "wave", Title: "DG", Body: "回到 docker 了，這次把容器們看緊一點。"},
			}
		}
		return []Reaction{
			{Key: "topic_docker_return_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "Docker is back on deck. We can spin this up cleanly."},
			{Key: "topic_docker_return_control", Variant: "focus", Pose: "wave", Title: "DG", Body: "Back to docker. Let us keep the containers under control."},
		}
	case "languages":
		if language == "zh-TW" {
			return []Reaction{
				{Key: "topic_languages_return", Variant: "focus", Pose: "wave", Title: "DG", Body: "回到語言這側了，我們把那個語感接回來。"},
			}
		}
		return []Reaction{
			{Key: "topic_languages_return", Variant: "focus", Pose: "wave", Title: "DG", Body: "Back to the language side. Let us pick that intuition back up."},
		}
	default:
		return nil
	}
}

func normalizeTopic(topic string) string {
	normalized := strings.ToLower(strings.TrimSpace(topic))
	if normalized == "" {
		return "all"
	}
	return normalized
}

func isGeneralTopic(topic string) bool {
	switch topic {
	case "all", "backend-tools", "languages":
		return true
	default:
		return false
	}
}

func lastInteractionWithinCooldown(state State, now time.Time) bool {
	return lastTimestampWithin(state.LastInteractionAt, clickCooldownWindow, now)
}

func lastReactionWithinCooldown(state State, now time.Time) bool {
	return lastTimestampWithin(state.LastReactionAt, ambientCooldownWindow, now)
}

func trackRapidClicks(state State, now time.Time) State {
	if lastTimestampWithin(state.LastRapidClickAt, rapidClickWindow, now) {
		state.RapidClickCount++
	} else {
		state.RapidClickCount = 1
	}
	state.LastRapidClickAt = stringPtr(now.Format(time.RFC3339))
	return state
}

func shouldTriggerRapidClickEgg(state State, now time.Time) bool {
	if state.RapidClickCount < 3 {
		return false
	}
	if state.LastEasterEggAt != nil && lastTimestampWithin(state.LastEasterEggAt, easterEggCooldown, now) {
		return false
	}
	return true
}

func shouldTriggerWelcomeBack(state State, now time.Time) bool {
	if state.LastInteractionAt == nil || *state.LastInteractionAt == "" {
		return false
	}
	if !lastTimestampOlderThan(state.LastInteractionAt, welcomeBackWindow, now) {
		return false
	}
	if state.LastWelcomeAt != nil && !lastTimestampOlderThan(state.LastWelcomeAt, welcomeBackWindow, now) {
		return false
	}
	if state.LastEasterEggAt != nil && !lastTimestampOlderThan(state.LastEasterEggAt, easterEggCooldown, now) {
		return false
	}
	return true
}

func shouldTriggerTimeOfDayEgg(state State, now time.Time) bool {
	if state.LastEasterEggAt != nil && lastTimestampWithin(state.LastEasterEggAt, easterEggCooldown, now) {
		return false
	}
	seed := state.ReactionStep + state.BondXP + now.Hour() + now.Minute()
	return seed%11 == 0
}

func shouldTriggerTopicInsideJoke(state State, trigger string, topic string, now time.Time) bool {
	if len(topicInsideJokeReactions("en", topic, trigger)) == 0 {
		return false
	}
	if state.LastEasterEggAt != nil && lastTimestampWithin(state.LastEasterEggAt, easterEggCooldown, now) {
		return false
	}
	seed := state.ReactionStep + state.BondXP + len(trigger) + len(normalizeTopic(topic)) + now.Minute()
	switch trigger {
	case TriggerClicked, TriggerWrong, TriggerCorrect, TriggerReturn:
		return seed%9 == 0
	default:
		return false
	}
}

func shouldTriggerRareCelebration(state State, trigger string, now time.Time) bool {
	if trigger != TriggerReviewComplete && trigger != TriggerLearnBreak {
		return false
	}
	if state.LastEasterEggAt != nil && lastTimestampWithin(state.LastEasterEggAt, easterEggCooldown, now) {
		return false
	}
	seed := state.ReactionStep + state.BondXP + len(trigger) + now.Minute()
	return seed%7 == 0
}

func lastTimestampWithin(value *string, window time.Duration, now time.Time) bool {
	if value == nil || *value == "" {
		return false
	}
	stamp, err := time.Parse(time.RFC3339, *value)
	if err != nil {
		return false
	}
	return now.Sub(stamp) < window
}

func lastTimestampOlderThan(value *string, window time.Duration, now time.Time) bool {
	if value == nil || *value == "" {
		return false
	}
	stamp, err := time.Parse(time.RFC3339, *value)
	if err != nil {
		return false
	}
	return now.Sub(stamp) >= window
}

func stringPtr(value string) *string {
	return &value
}
