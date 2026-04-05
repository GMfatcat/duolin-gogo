package pet

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const cooldownWindow = 15 * time.Second

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
	timestamp := now.Format(time.RFC3339)
	state.LastReactionAt = &timestamp

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
			Reaction: reactionFor(language, TriggerClicked, state.Stage, true),
		}, nil
	}

	state.BondXP += 1
	state.Stage = stageForXP(state.BondXP)
	timestamp := now.Format(time.RFC3339)
	state.LastInteractionAt = &timestamp
	state.LastReactionAt = &timestamp

	if err := Save(path, state); err != nil {
		return InteractionResult{}, err
	}

	return InteractionResult{
		State:    state,
		Reaction: reactionFor(language, TriggerClicked, state.Stage, false),
	}, nil
}

func ReactionForTrigger(path string, trigger string, language string, now time.Time) (InteractionResult, error) {
	state, err := Load(path)
	if err != nil {
		return InteractionResult{}, err
	}

	timestamp := now.Format(time.RFC3339)
	state.LastReactionAt = &timestamp

	if err := Save(path, state); err != nil {
		return InteractionResult{}, err
	}

	return InteractionResult{
		State:    state,
		Reaction: reactionFor(language, trigger, state.Stage, false),
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

func lastInteractionWithinCooldown(state State, now time.Time) bool {
	if state.LastInteractionAt == nil || *state.LastInteractionAt == "" {
		return false
	}

	lastInteractionAt, err := time.Parse(time.RFC3339, *state.LastInteractionAt)
	if err != nil {
		return false
	}

	return now.Sub(lastInteractionAt) < cooldownWindow
}

func reactionFor(language string, trigger string, stage int, cooldown bool) Reaction {
	if cooldown {
		if language == "zh-TW" {
			return Reaction{Key: "cooldown", Variant: "focus", Title: "DG", Body: "我有在聽，再一下就回你。"}
		}
		return Reaction{Key: "cooldown", Variant: "focus", Title: "DG", Body: "I heard you. Give me a beat."}
	}

	switch trigger {
	case TriggerCorrect:
		return correctReaction(language, stage)
	case TriggerWrong:
		return wrongReaction(language, stage)
	case TriggerLearnBreak:
		return learnBreakReaction(language, stage)
	case TriggerReviewComplete:
		return reviewCompleteReaction(language, stage)
	case TriggerReturn:
		return returnReaction(language, stage)
	default:
		return clickReaction(language, stage)
	}
}

func clickReaction(language string, stage int) Reaction {
	switch {
	case stage >= 2:
		if language == "zh-TW" {
			return Reaction{Key: "stage_two_click", Variant: "celebration", Title: "DG", Body: "你又回來了，我開始抓到你的節奏了。"}
		}
		return Reaction{Key: "stage_two_click", Variant: "celebration", Title: "DG", Body: "You are back. I am starting to learn your rhythm."}
	case stage >= 1:
		if language == "zh-TW" {
			return Reaction{Key: "stage_one_click", Variant: "focus", Title: "DG", Body: "好，我陪你把這輪慢慢走完。"}
		}
		return Reaction{Key: "stage_one_click", Variant: "focus", Title: "DG", Body: "Alright, let us work through this batch together."}
	default:
		if language == "zh-TW" {
			return Reaction{Key: "stage_zero_click", Variant: "neutral", Title: "DG", Body: "我在這裡，點我我會慢慢變熟。"}
		}
		return Reaction{Key: "stage_zero_click", Variant: "neutral", Title: "DG", Body: "I am here. Keep tapping in and I will warm up."}
	}
}

func correctReaction(language string, stage int) Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return Reaction{Key: "correct_stage_one", Variant: "celebration", Title: "DG", Body: "這題很穩，我有看到你越來越熟。"}
		}
		return Reaction{Key: "correct_stage_one", Variant: "celebration", Title: "DG", Body: "That was clean. I can tell this is starting to stick."}
	}
	if language == "zh-TW" {
		return Reaction{Key: "correct_stage_zero", Variant: "celebration", Title: "DG", Body: "答得不錯，先把這個手感記住。"}
	}
	return Reaction{Key: "correct_stage_zero", Variant: "celebration", Title: "DG", Body: "Nice hit. Hold on to that feeling for the next one."}
}

func wrongReaction(language string, stage int) Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return Reaction{Key: "wrong_stage_one", Variant: "warning", Title: "DG", Body: "沒關係，這種差一點的題最值得撿回來。"}
		}
		return Reaction{Key: "wrong_stage_one", Variant: "warning", Title: "DG", Body: "That is okay. These almost-right misses are worth keeping."}
	}
	if language == "zh-TW" {
		return Reaction{Key: "wrong_stage_zero", Variant: "warning", Title: "DG", Body: "先記住差異就好，下一次會更穩。"}
	}
	return Reaction{Key: "wrong_stage_zero", Variant: "warning", Title: "DG", Body: "Just hold on to the difference. The next pass will feel steadier."}
}

func learnBreakReaction(language string, stage int) Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return Reaction{Key: "learn_break_stage_one", Variant: "focus", Title: "DG", Body: "這輪收得很好，先讓腦袋留一點空間。"}
		}
		return Reaction{Key: "learn_break_stage_one", Variant: "focus", Title: "DG", Body: "That batch landed well. Give your brain a little room now."}
	}
	if language == "zh-TW" {
		return Reaction{Key: "learn_break_stage_zero", Variant: "focus", Title: "DG", Body: "先休一下，等下一輪來就好。"}
	}
	return Reaction{Key: "learn_break_stage_zero", Variant: "focus", Title: "DG", Body: "Take a short beat. The next batch can wait."}
}

func reviewCompleteReaction(language string, stage int) Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return Reaction{Key: "review_complete_stage_one", Variant: "celebration", Title: "DG", Body: "這輪複習收得很完整，我記得你有把它撿回來。"}
		}
		return Reaction{Key: "review_complete_stage_one", Variant: "celebration", Title: "DG", Body: "That review batch closed out nicely. I can feel the loop settling in."}
	}
	if language == "zh-TW" {
		return Reaction{Key: "review_complete_stage_zero", Variant: "celebration", Title: "DG", Body: "這輪複習做完了，先收一下成果。"}
	}
	return Reaction{Key: "review_complete_stage_zero", Variant: "celebration", Title: "DG", Body: "That review batch is done. Take a moment and let it settle."}
}

func returnReaction(language string, stage int) Reaction {
	if stage >= 1 {
		if language == "zh-TW" {
			return Reaction{Key: "return_stage_one", Variant: "focus", Title: "DG", Body: "你回來了，我們可以接著往下推。"}
		}
		return Reaction{Key: "return_stage_one", Variant: "focus", Title: "DG", Body: "You are back. We can pick up the thread from here."}
	}
	if language == "zh-TW" {
		return Reaction{Key: "return_stage_zero", Variant: "focus", Title: "DG", Body: "好，下一輪可以開始了。"}
	}
	return Reaction{Key: "return_stage_zero", Variant: "focus", Title: "DG", Body: "Alright, the next round is ready."}
}
