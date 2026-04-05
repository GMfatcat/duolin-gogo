package pet

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const (
	clickCooldownWindow   = 15 * time.Second
	ambientCooldownWindow = 6 * time.Second
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
	LastInteractionAt *string `json:"last_interaction_at,omitempty"`
	LastReactionAt    *string `json:"last_reaction_at,omitempty"`
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

func Interact(path string, language string, now time.Time) (InteractionResult, error) {
	state, err := Load(path)
	if err != nil {
		return InteractionResult{}, err
	}

	if lastInteractionWithinCooldown(state, now) {
		return InteractionResult{
			State:    state,
			Reaction: cooldownReaction(language),
		}, nil
	}

	state.BondXP += 1
	state.Stage = stageForXP(state.BondXP)
	stamp := now.Format(time.RFC3339)
	state.LastInteractionAt = &stamp
	state.LastReactionAt = &stamp

	if err := Save(path, state); err != nil {
		return InteractionResult{}, err
	}

	return InteractionResult{
		State:    state,
		Reaction: pickReaction(clickReactions(language, state.Stage), state, TriggerClicked, now),
	}, nil
}

func ReactionForTrigger(path string, trigger string, language string, now time.Time) (InteractionResult, error) {
	state, err := Load(path)
	if err != nil {
		return InteractionResult{}, err
	}

	if !shouldEmitAmbientReaction(state, trigger, now) {
		return InteractionResult{State: state}, nil
	}

	state.LastReactionAt = stringPtr(now.Format(time.RFC3339))
	if err := Save(path, state); err != nil {
		return InteractionResult{}, err
	}

	return InteractionResult{
		State:    state,
		Reaction: reactionForTrigger(language, trigger, state, now),
	}, nil
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

	seed := int(now.Unix()/60) + state.BondXP + len(trigger)
	switch trigger {
	case TriggerCorrect:
		return seed%3 != 0
	case TriggerWrong:
		return seed%2 == 0
	default:
		return true
	}
}

func reactionForTrigger(language string, trigger string, state State, now time.Time) Reaction {
	switch trigger {
	case TriggerCorrect:
		return pickReaction(correctReactions(language, state.Stage), state, trigger, now)
	case TriggerWrong:
		return pickReaction(wrongReactions(language, state.Stage), state, trigger, now)
	case TriggerLearnBreak:
		return pickReaction(learnBreakReactions(language, state.Stage), state, trigger, now)
	case TriggerReviewComplete:
		return pickReaction(reviewCompleteReactions(language, state.Stage), state, trigger, now)
	case TriggerReturn:
		return pickReaction(returnReactions(language, state.Stage), state, trigger, now)
	default:
		return pickReaction(clickReactions(language, state.Stage), state, trigger, now)
	}
}

func pickReaction(pool []Reaction, state State, trigger string, now time.Time) Reaction {
	if len(pool) == 0 {
		return Reaction{Key: "fallback", Variant: "neutral", Pose: "idle", Title: "DG", Body: "..."}
	}

	index := int((now.Unix()/60)+int64(state.BondXP)+int64(len(trigger))) % len(pool)
	return pool[index]
}

func cooldownReaction(language string) Reaction {
	if language == "zh-TW" {
		return Reaction{Key: "cooldown", Variant: "focus", Pose: "rest", Title: "DG", Body: "我有聽到，先讓我喘一口氣。"}
	}
	return Reaction{Key: "cooldown", Variant: "focus", Pose: "rest", Title: "DG", Body: "I heard you. Give me a beat."}
}

func clickReactions(language string, stage int) []Reaction {
	switch {
	case stage >= 2:
		if language == "zh-TW" {
			return []Reaction{
				{Key: "stage_two_click_warm", Variant: "celebration", Pose: "spark", Title: "DG", Body: "你回來了，我開始抓到你的節奏了。"},
				{Key: "stage_two_click_sync", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪我跟得上，你只管繼續。"},
			}
		}
		return []Reaction{
			{Key: "stage_two_click_warm", Variant: "celebration", Pose: "spark", Title: "DG", Body: "You are back. I am starting to learn your rhythm."},
			{Key: "stage_two_click_sync", Variant: "celebration", Pose: "spark", Title: "DG", Body: "I am in sync now. You can keep the pace up."},
		}
	case stage >= 1:
		if language == "zh-TW" {
			return []Reaction{
				{Key: "stage_one_click_focus", Variant: "focus", Pose: "wave", Title: "DG", Body: "好，這一輪我們一起走完。"},
				{Key: "stage_one_click_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "我準備好了，你先開題。"},
			}
		}
		return []Reaction{
			{Key: "stage_one_click_focus", Variant: "focus", Pose: "wave", Title: "DG", Body: "Alright, let us work through this batch together."},
			{Key: "stage_one_click_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "I am ready. You take the first step."},
		}
	default:
		if language == "zh-TW" {
			return []Reaction{
				{Key: "stage_zero_click_intro", Variant: "neutral", Pose: "idle", Title: "DG", Body: "我在這裡，慢慢來就好。"},
				{Key: "stage_zero_click_warmup", Variant: "neutral", Pose: "idle", Title: "DG", Body: "多點我幾次，我會更快進入狀態。"},
			}
		}
		return []Reaction{
			{Key: "stage_zero_click_intro", Variant: "neutral", Pose: "idle", Title: "DG", Body: "I am here. Keep tapping in and I will warm up."},
			{Key: "stage_zero_click_warmup", Variant: "neutral", Pose: "idle", Title: "DG", Body: "Tap back in a little more and I will wake up faster."},
		}
	}
}

func correctReactions(language string, stage int) []Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "correct_stage_one_clean", Variant: "celebration", Pose: "nod", Title: "DG", Body: "這題很乾淨，感覺開始黏住了。"},
				{Key: "correct_stage_one_locking", Variant: "celebration", Pose: "nod", Title: "DG", Body: "對，就是這種手感，先記住。"},
			}
		}
		return []Reaction{
			{Key: "correct_stage_one_clean", Variant: "celebration", Pose: "nod", Title: "DG", Body: "That was clean. I can tell this is starting to stick."},
			{Key: "correct_stage_one_locking", Variant: "celebration", Pose: "nod", Title: "DG", Body: "Yes, that is the feeling. Keep it for the next card."},
		}
	}
	if language == "zh-TW" {
		return []Reaction{
			{Key: "correct_stage_zero_nice", Variant: "celebration", Pose: "nod", Title: "DG", Body: "漂亮，這一題先收下來。"},
			{Key: "correct_stage_zero_hold", Variant: "celebration", Pose: "nod", Title: "DG", Body: "很好，把這個感覺帶到下一題。"},
		}
	}
	return []Reaction{
		{Key: "correct_stage_zero_nice", Variant: "celebration", Pose: "nod", Title: "DG", Body: "Nice hit. Hold on to that feeling for the next one."},
		{Key: "correct_stage_zero_hold", Variant: "celebration", Pose: "nod", Title: "DG", Body: "Good catch. Bring that same energy into the next card."},
	}
}

func wrongReactions(language string, stage int) []Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "wrong_stage_one_almost", Variant: "warning", Pose: "think", Title: "DG", Body: "沒關係，這種差一點最值得記。"},
				{Key: "wrong_stage_one_keep", Variant: "warning", Pose: "think", Title: "DG", Body: "先記住差異，下次會更穩。"},
			}
		}
		return []Reaction{
			{Key: "wrong_stage_one_almost", Variant: "warning", Pose: "think", Title: "DG", Body: "That is okay. These almost-right misses are worth keeping."},
			{Key: "wrong_stage_one_keep", Variant: "warning", Pose: "think", Title: "DG", Body: "Keep the difference in view. The next pass will feel steadier."},
		}
	}
	if language == "zh-TW" {
		return []Reaction{
			{Key: "wrong_stage_zero_difference", Variant: "warning", Pose: "think", Title: "DG", Body: "先抓住差異，下一輪就會好很多。"},
			{Key: "wrong_stage_zero_retry", Variant: "warning", Pose: "think", Title: "DG", Body: "這題先別怕，等下再看一次。"},
		}
	}
	return []Reaction{
		{Key: "wrong_stage_zero_difference", Variant: "warning", Pose: "think", Title: "DG", Body: "Just hold on to the difference. The next pass will feel steadier."},
		{Key: "wrong_stage_zero_retry", Variant: "warning", Pose: "think", Title: "DG", Body: "Do not worry about this one yet. We can loop back cleanly."},
	}
}

func learnBreakReactions(language string, stage int) []Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "learn_break_stage_one_land", Variant: "focus", Pose: "rest", Title: "DG", Body: "這輪收得不錯，先讓腦袋留點空間。"},
				{Key: "learn_break_stage_one_room", Variant: "focus", Pose: "rest", Title: "DG", Body: "停一下剛剛好，讓記憶沉一沉。"},
			}
		}
		return []Reaction{
			{Key: "learn_break_stage_one_land", Variant: "focus", Pose: "rest", Title: "DG", Body: "That batch landed well. Give your brain a little room now."},
			{Key: "learn_break_stage_one_room", Variant: "focus", Pose: "rest", Title: "DG", Body: "A short pause is right. Let the last few cards settle."},
		}
	}
	if language == "zh-TW" {
		return []Reaction{
			{Key: "learn_break_stage_zero_wait", Variant: "focus", Pose: "rest", Title: "DG", Body: "先休息一下，下一輪不急。"},
			{Key: "learn_break_stage_zero_pause", Variant: "focus", Pose: "rest", Title: "DG", Body: "這裡先停一下，等等再回來。"},
		}
	}
	return []Reaction{
		{Key: "learn_break_stage_zero_wait", Variant: "focus", Pose: "rest", Title: "DG", Body: "Take a short beat. The next batch can wait."},
		{Key: "learn_break_stage_zero_pause", Variant: "focus", Pose: "rest", Title: "DG", Body: "Pause here for a moment. The next round is fine waiting."},
	}
}

func reviewCompleteReactions(language string, stage int) []Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "review_complete_stage_one_closed", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪複習收得很漂亮，節奏在成形了。"},
				{Key: "review_complete_stage_one_settle", Variant: "celebration", Pose: "spark", Title: "DG", Body: "很好，讓這一輪在腦中沉下去。"},
			}
		}
		return []Reaction{
			{Key: "review_complete_stage_one_closed", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That review batch closed out nicely. I can feel the loop settling in."},
			{Key: "review_complete_stage_one_settle", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Nice finish. Let that review loop settle in a bit."},
		}
	}
	if language == "zh-TW" {
		return []Reaction{
			{Key: "review_complete_stage_zero_done", Variant: "celebration", Pose: "spark", Title: "DG", Body: "這輪複習完成了，先讓它停一下。"},
			{Key: "review_complete_stage_zero_breathe", Variant: "celebration", Pose: "spark", Title: "DG", Body: "複習做完了，現在先喘一口氣。"},
		}
	}
	return []Reaction{
		{Key: "review_complete_stage_zero_done", Variant: "celebration", Pose: "spark", Title: "DG", Body: "That review batch is done. Take a moment and let it settle."},
		{Key: "review_complete_stage_zero_breathe", Variant: "celebration", Pose: "spark", Title: "DG", Body: "Review complete. Take a breath before you move on."},
	}
}

func returnReactions(language string, stage int) []Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return []Reaction{
				{Key: "return_stage_one_pickup", Variant: "focus", Pose: "wave", Title: "DG", Body: "你回來了，我們從這裡接著走。"},
				{Key: "return_stage_one_thread", Variant: "focus", Pose: "wave", Title: "DG", Body: "剛剛那條線還在，現在可以繼續。"},
			}
		}
		return []Reaction{
			{Key: "return_stage_one_pickup", Variant: "focus", Pose: "wave", Title: "DG", Body: "You are back. We can pick up the thread from here."},
			{Key: "return_stage_one_thread", Variant: "focus", Pose: "wave", Title: "DG", Body: "That thread is still here. We can keep going now."},
		}
	}
	if language == "zh-TW" {
		return []Reaction{
			{Key: "return_stage_zero_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "下一輪已經準備好了。"},
			{Key: "return_stage_zero_resume", Variant: "focus", Pose: "wave", Title: "DG", Body: "好，現在可以重新開始。"},
		}
	}
	return []Reaction{
		{Key: "return_stage_zero_ready", Variant: "focus", Pose: "wave", Title: "DG", Body: "Alright, the next round is ready."},
		{Key: "return_stage_zero_resume", Variant: "focus", Pose: "wave", Title: "DG", Body: "Okay, we can start fresh from here."},
	}
}

func lastInteractionWithinCooldown(state State, now time.Time) bool {
	return lastTimestampWithin(state.LastInteractionAt, clickCooldownWindow, now)
}

func lastReactionWithinCooldown(state State, now time.Time) bool {
	return lastTimestampWithin(state.LastReactionAt, ambientCooldownWindow, now)
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

func stringPtr(value string) *string {
	return &value
}
