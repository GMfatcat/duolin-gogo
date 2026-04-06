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

	first, err := Interact(path, "zh-TW", "all", now)
	if err != nil {
		t.Fatalf("first interaction failed: %v", err)
	}
	if first.State.BondXP != 1 {
		t.Fatalf("expected first interaction to add xp, got %d", first.State.BondXP)
	}

	second, err := Interact(path, "zh-TW", "all", now.Add(5*time.Second))
	if err != nil {
		t.Fatalf("second interaction failed: %v", err)
	}
	if second.State.BondXP != 1 {
		t.Fatalf("expected cooldown interaction not to add xp, got %d", second.State.BondXP)
	}
	if second.Reaction.Key != "cooldown" {
		t.Fatalf("expected cooldown reaction, got %s", second.Reaction.Key)
	}
	if second.Reaction.Pose != "rest" {
		t.Fatalf("expected cooldown pose rest, got %s", second.Reaction.Pose)
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

	result, err := Interact(path, "en", "all", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("interaction failed: %v", err)
	}

	if result.State.Stage < 2 {
		t.Fatalf("expected hidden stage at least 2, got %d", result.State.Stage)
	}
	if result.Reaction.Key != "stage_two_click_warm" &&
		result.Reaction.Key != "stage_two_click_sync" &&
		result.Reaction.Key != "stage_two_click_grin" &&
		result.Reaction.Key != "stage_two_click_anchor" &&
		result.Reaction.Key != "stage_two_click_companion" {
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

	result, err := ReactionForTrigger(path, TriggerCorrect, "en", "all", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("reaction for trigger failed: %v", err)
	}

	if result.Reaction.Key != "correct_stage_one_clean" && result.Reaction.Key != "correct_stage_one_locking" {
		t.Fatalf("expected stage-one correct reaction, got %s", result.Reaction.Key)
	}
	if result.Reaction.Pose != "nod" {
		t.Fatalf("expected correct pose nod, got %s", result.Reaction.Pose)
	}

	result, err = ReactionForTrigger(path, TriggerReviewComplete, "en", "all", base.Add(21*time.Minute))
	if err != nil {
		t.Fatalf("review complete reaction failed: %v", err)
	}

	if result.Reaction.Key != "review_complete_stage_one_closed" && result.Reaction.Key != "review_complete_stage_one_settle" {
		t.Fatalf("expected review-complete reaction, got %s", result.Reaction.Key)
	}
}

func TestReactionForTriggerCanStayQuietForLowNoisePrompts(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for index := 0; index < 4; index++ {
		if _, err := RecordStudyEvent(path, StudyEventAnsweredCorrect, base.Add(time.Duration(index)*time.Minute)); err != nil {
			t.Fatalf("seed study event %d failed: %v", index, err)
		}
	}

	for offset := 0; offset < 6; offset++ {
		result, err := ReactionForTrigger(path, TriggerCorrect, "en", "all", base.Add(time.Duration(17+offset)*time.Minute))
		if err != nil {
			t.Fatalf("reaction for quiet trigger failed: %v", err)
		}
		if result.Reaction.Body == "" {
			return
		}
	}

	t.Fatal("expected at least one quiet correct prompt within the sampled window")
}

func TestReviewCompleteAlwaysEmitsReactionEvenAfterRecentCue(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	if _, err := RecordStudyEvent(path, StudyEventReviewBatch, base); err != nil {
		t.Fatalf("seed review batch failed: %v", err)
	}

	first, err := ReactionForTrigger(path, TriggerLearnBreak, "en", "all", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("learn break reaction failed: %v", err)
	}
	if first.Reaction.Body == "" {
		t.Fatal("expected learn break reaction body")
	}

	second, err := ReactionForTrigger(path, TriggerReviewComplete, "en", "all", base.Add(21*time.Minute))
	if err != nil {
		t.Fatalf("review complete reaction failed: %v", err)
	}
	if second.Reaction.Body == "" {
		t.Fatal("expected review complete reaction body")
	}
	if second.Reaction.Pose != "spark" {
		t.Fatalf("expected review complete pose spark, got %s", second.Reaction.Pose)
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

	first, err := ReactionForTrigger(path, TriggerCorrect, "en", "all", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("first reaction failed: %v", err)
	}

	second, err := ReactionForTrigger(path, TriggerCorrect, "en", "all", base.Add(21*time.Minute))
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

	first, err := Interact(path, "en", "all", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("first interaction failed: %v", err)
	}

	second, err := Interact(path, "en", "all", base.Add(36*time.Minute))
	if err != nil {
		t.Fatalf("second interaction failed: %v", err)
	}

	if first.Reaction.Key == second.Reaction.Key {
		t.Fatalf("expected click pool rotation, got same key %s", first.Reaction.Key)
	}
}

func TestReactionForTriggerUsesTopicAwarePool(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 5, 10, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	result, err := ReactionForTrigger(path, TriggerReturn, "en", "docker", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("docker return reaction failed: %v", err)
	}

	if result.Reaction.Body == "" {
		t.Fatal("expected docker-themed reaction body")
	}
	if result.Reaction.Body != "Docker is back on deck. We can spin this up cleanly." &&
		result.Reaction.Body != "Back to docker. Let us keep the containers under control." {
		t.Fatalf("expected docker-themed reaction, got %q", result.Reaction.Body)
	}

	other, err := ReactionForTrigger(path, TriggerCorrect, "en", "languages", base.Add(21*time.Minute))
	if err != nil {
		t.Fatalf("languages correct reaction failed: %v", err)
	}

	if other.Reaction.Body == "" {
		t.Fatal("expected language-themed reaction body")
	}
	if other.Reaction.Body != "Nice catch. Your language instincts are settling in." &&
		other.Reaction.Body != "That was clean. The language side is starting to click." {
		t.Fatalf("expected language-themed reaction, got %q", other.Reaction.Body)
	}
}

func TestInteractTriggersRapidClickEasterEgg(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 6, 11, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	if _, err := Interact(path, "en", "all", base); err != nil {
		t.Fatalf("first interaction failed: %v", err)
	}
	if _, err := Interact(path, "en", "all", base.Add(4*time.Second)); err != nil {
		t.Fatalf("second interaction failed: %v", err)
	}
	third, err := Interact(path, "en", "all", base.Add(8*time.Second))
	if err != nil {
		t.Fatalf("third interaction failed: %v", err)
	}

	if third.Reaction.Key != "rapid_click_stage_zero_notice" && third.Reaction.Key != "rapid_click_stage_zero_tickle" {
		t.Fatalf("expected rapid click easter egg, got %s", third.Reaction.Key)
	}
	if third.Reaction.Pose != "spark" {
		t.Fatalf("expected rapid click pose spark, got %s", third.Reaction.Pose)
	}
}

func TestInteractTriggersWelcomeBackAfterLongGap(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 6, 11, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	first, err := Interact(path, "en", "git", base)
	if err != nil {
		t.Fatalf("first interaction failed: %v", err)
	}
	if first.Reaction.Key != "stage_zero_click_intro" &&
		first.Reaction.Key != "stage_zero_click_warmup" &&
		first.Reaction.Key != "topic_git_click_history" &&
		first.Reaction.Key != "topic_git_click_branch" {
		t.Fatalf("expected normal click reaction, got %s", first.Reaction.Key)
	}

	second, err := Interact(path, "en", "git", base.Add(9*time.Hour))
	if err != nil {
		t.Fatalf("second interaction failed: %v", err)
	}

	if second.Reaction.Key != "welcome_back_git" && second.Reaction.Key != "welcome_back_git_rebase" {
		t.Fatalf("expected welcome-back reaction, got %s", second.Reaction.Key)
	}
	if second.Reaction.Pose != "wave" {
		t.Fatalf("expected welcome-back pose wave, got %s", second.Reaction.Pose)
	}
}

func TestInteractRotatesClickLinesWithinSameMinute(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 6, 12, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for index := 0; index < 4; index++ {
		if _, err := RecordStudyEvent(path, StudyEventAnsweredCorrect, base.Add(time.Duration(index)*time.Minute)); err != nil {
			t.Fatalf("seed study event %d failed: %v", index, err)
		}
	}

	first, err := Interact(path, "en", "all", base.Add(20*time.Minute))
	if err != nil {
		t.Fatalf("first interaction failed: %v", err)
	}

	second, err := Interact(path, "en", "all", base.Add(20*time.Minute+20*time.Second))
	if err != nil {
		t.Fatalf("second interaction failed: %v", err)
	}

	if first.Reaction.Key == second.Reaction.Key {
		t.Fatalf("expected click lines to vary within the same minute, got %s twice", first.Reaction.Key)
	}
}

func TestReactionForTriggerCanUseTopicInsideJoke(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 6, 13, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for offset := 0; offset < 20; offset++ {
		result, err := ReactionForTrigger(path, TriggerReturn, "en", "http", base.Add(time.Duration(offset)*time.Minute))
		if err != nil {
			t.Fatalf("topic inside joke check failed: %v", err)
		}
		if result.Reaction.Key == "inside_joke_http_status" {
			return
		}
	}

	t.Fatal("expected at least one topic-inside-joke reaction in sampled window")
}

func TestReactionForTriggerCanUseRareCelebration(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 6, 14, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for index := 0; index < 6; index++ {
		if _, err := RecordStudyEvent(path, StudyEventReviewBatch, base.Add(time.Duration(index)*time.Minute)); err != nil {
			t.Fatalf("seed review event %d failed: %v", index, err)
		}
	}

	for offset := 0; offset < 20; offset++ {
		result, err := ReactionForTrigger(path, TriggerReviewComplete, "en", "backend-tools", base.Add(time.Duration(offset)*5*time.Minute))
		if err != nil {
			t.Fatalf("rare celebration check failed: %v", err)
		}
		if result.Reaction.Key == "rare_celebration_backend" || result.Reaction.Key == "rare_celebration_general" || result.Reaction.Key == "rare_celebration_general_soft" {
			return
		}
	}

	t.Fatal("expected at least one rare celebration reaction in sampled window")
}

func TestInteractCanTriggerTopicStreakEgg(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 6, 15, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	first, err := Interact(path, "en", "sql", base)
	if err != nil {
		t.Fatalf("first interaction failed: %v", err)
	}
	if first.Reaction.Body == "" {
		t.Fatal("expected first click reaction body")
	}

	if _, err := Interact(path, "en", "sql", base.Add(20*time.Second)); err != nil {
		t.Fatalf("second interaction failed: %v", err)
	}

	third, err := Interact(path, "en", "sql", base.Add(40*time.Second))
	if err != nil {
		t.Fatalf("third interaction failed: %v", err)
	}

	if third.Reaction.Key != "topic_streak_sql_where" {
		t.Fatalf("expected topic streak reaction, got %s", third.Reaction.Key)
	}
}

func TestReactionForTriggerCanUseAlmostThereEncouragement(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pet.json")
	base := time.Date(2026, 4, 6, 16, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))

	for index := 0; index < 3; index++ {
		if _, err := Interact(path, "en", "backend-tools", base.Add(time.Duration(index)*20*time.Second)); err != nil {
			t.Fatalf("seed interaction %d failed: %v", index, err)
		}
	}

	result, err := ReactionForTriggerWithContext(path, TriggerCorrect, "en", "backend-tools", Context{EncourageTopic: "go"}, base.Add(2*time.Minute))
	if err != nil {
		t.Fatalf("almost-there reaction failed: %v", err)
	}

	if result.Reaction.Key != "almost_there_go" {
		t.Fatalf("expected almost-there reaction, got %s", result.Reaction.Key)
	}
}
