package pet

import (
	"path/filepath"
	"testing"
	"time"
)

func TestRecordStudyEventAdvancesHiddenGrowthStage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	state, err := RecordStudyEvent(path, StudyEventAnsweredCorrect, now)
	if err != nil {
		t.Fatalf("record first study event failed: %v", err)
	}

	if state.BondXP != 2 {
		t.Fatalf("expected bond xp 2, got %d", state.BondXP)
	}
	if state.Stage != 0 {
		t.Fatalf("expected stage 0, got %d", state.Stage)
	}

	state, err = RecordStudyEvent(path, StudyEventAnsweredCorrect, now.Add(1*time.Minute))
	if err != nil {
		t.Fatalf("record second study event failed: %v", err)
	}
	state, err = RecordStudyEvent(path, StudyEventAnsweredCorrect, now.Add(2*time.Minute))
	if err != nil {
		t.Fatalf("record third study event failed: %v", err)
	}

	if state.BondXP != 6 {
		t.Fatalf("expected bond xp 6, got %d", state.BondXP)
	}
	if state.Stage != 1 {
		t.Fatalf("expected stage 1 after enough study activity, got %d", state.Stage)
	}
}

func TestInteractUsesCooldownWithoutAddingMoreGrowth(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	now := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	first, err := Interact(path, "zh-TW", now)
	if err != nil {
		t.Fatalf("first interaction failed: %v", err)
	}
	if first.State.BondXP != 1 {
		t.Fatalf("expected first interaction to add xp, got %d", first.State.BondXP)
	}

	second, err := Interact(path, "zh-TW", now.Add(5*time.Second))
	if err != nil {
		t.Fatalf("second interaction failed: %v", err)
	}
	if second.State.BondXP != 1 {
		t.Fatalf("expected cooldown interaction not to add xp, got %d", second.State.BondXP)
	}
	if second.Reaction.Key != "cooldown" {
		t.Fatalf("expected cooldown reaction, got %s", second.Reaction.Key)
	}
}

func TestInteractUnlocksRicherReactionPoolAtHigherStage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for index := 0; index < 8; index++ {
		if _, err := RecordStudyEvent(path, StudyEventAnsweredCorrect, base.Add(time.Duration(index)*time.Minute)); err != nil {
			t.Fatalf("seed study event %d failed: %v", index, err)
		}
	}

	result, err := Interact(path, "en", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("interaction failed: %v", err)
	}

	if result.State.Stage < 2 {
		t.Fatalf("expected hidden stage at least 2, got %d", result.State.Stage)
	}
	if result.Reaction.Key != "stage_two_click_warm" && result.Reaction.Key != "stage_two_click_sync" {
		t.Fatalf("expected stage two click pool reaction, got %s", result.Reaction.Key)
	}
}

func TestReactionForTriggerUsesContextSpecificPool(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for index := 0; index < 4; index++ {
		if _, err := RecordStudyEvent(path, StudyEventAnsweredCorrect, base.Add(time.Duration(index)*time.Minute)); err != nil {
			t.Fatalf("seed study event %d failed: %v", index, err)
		}
	}

	result, err := ReactionForTrigger(path, TriggerCorrect, "en", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("reaction for trigger failed: %v", err)
	}

	if result.Reaction.Key != "correct_stage_one_clean" && result.Reaction.Key != "correct_stage_one_locking" {
		t.Fatalf("expected stage-one correct reaction, got %s", result.Reaction.Key)
	}

	result, err = ReactionForTrigger(path, TriggerReviewComplete, "en", base.Add(21*time.Minute))
	if err != nil {
		t.Fatalf("review complete reaction failed: %v", err)
	}

	if result.Reaction.Key != "review_complete_stage_one_closed" && result.Reaction.Key != "review_complete_stage_one_settle" {
		t.Fatalf("expected review-complete reaction, got %s", result.Reaction.Key)
	}
}

func TestReactionForTriggerRotatesWithinStagePool(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for index := 0; index < 4; index++ {
		if _, err := RecordStudyEvent(path, StudyEventAnsweredCorrect, base.Add(time.Duration(index)*time.Minute)); err != nil {
			t.Fatalf("seed study event %d failed: %v", index, err)
		}
	}

	first, err := ReactionForTrigger(path, TriggerCorrect, "en", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("first reaction failed: %v", err)
	}

	second, err := ReactionForTrigger(path, TriggerCorrect, "en", base.Add(21*time.Minute))
	if err != nil {
		t.Fatalf("second reaction failed: %v", err)
	}

	if first.Reaction.Key == second.Reaction.Key {
		t.Fatalf("expected reaction pool rotation, got same key %s", first.Reaction.Key)
	}
}

func TestInteractUsesDifferentClickVariantsAtHigherStage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for index := 0; index < 8; index++ {
		if _, err := RecordStudyEvent(path, StudyEventAnsweredCorrect, base.Add(time.Duration(index)*time.Minute)); err != nil {
			t.Fatalf("seed study event %d failed: %v", index, err)
		}
	}

	first, err := Interact(path, "en", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("first interaction failed: %v", err)
	}

	second, err := Interact(path, "en", base.Add(36*time.Minute))
	if err != nil {
		t.Fatalf("second interaction failed: %v", err)
	}

	if first.Reaction.Key == second.Reaction.Key {
		t.Fatalf("expected click pool rotation, got same key %s", first.Reaction.Key)
	}
}
