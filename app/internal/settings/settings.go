package settings

import (
	"encoding/json"
	"os"
)

type File struct {
	Version                     int      `json:"version"`
	KnowledgeDirectories        []string `json:"knowledge_directories"`
	SelectedTopic               string   `json:"selected_topic"`
	NotificationIntervalMinutes int      `json:"notification_interval_minutes"`
	ActiveHours                 struct {
		Enabled bool   `json:"enabled"`
		Start   string `json:"start"`
		End     string `json:"end"`
	} `json:"active_hours"`
	ReviewSchedule struct {
		Mode      string  `json:"mode"`
		Weekday   *string `json:"weekday"`
		Time      string  `json:"time"`
		BatchSize int     `json:"batch_size"`
	} `json:"review_schedule"`
	Language struct {
		Default     string `json:"default"`
		AllowToggle bool   `json:"allow_toggle"`
	} `json:"language"`
	Notifications struct {
		Style     string `json:"style"`
		TitleMode string `json:"title_mode"`
	} `json:"notifications"`
	StudyRules struct {
		MaxNewCardsPerDay          int `json:"max_new_cards_per_day"`
		SnoozeMinutes              int `json:"snooze_minutes"`
		CooldownAfterAnswerMinutes int `json:"cooldown_after_answer_minutes"`
		RevealSpeed                string `json:"reveal_speed"`
	} `json:"study_rules"`
	Onboarding struct {
		Seen bool `json:"seen"`
	} `json:"onboarding"`
}

func Load(path string) (File, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return File{}, err
	}

	var file File
	if err := json.Unmarshal(bytes, &file); err != nil {
		return File{}, err
	}

	return file, nil
}
