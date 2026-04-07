package diagnostics

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type File struct {
	Version     int     `json:"version"`
	GeneratedAt *string `json:"generated_at"`
	Errors      []Error `json:"errors"`
}

type Error struct {
	SourcePath   string `json:"source_path"`
	Severity     string `json:"severity,omitempty"`
	Code         string `json:"code"`
	Field        string `json:"field,omitempty"`
	Message      string `json:"message"`
	SuggestionZH string `json:"suggestion_zh,omitempty"`
	SuggestionEN string `json:"suggestion_en,omitempty"`
}

func Load(path string) (File, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return File{
			Version: 1,
			Errors:  []Error{},
		}, nil
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return File{}, err
	}
	bytes = []byte(strings.TrimPrefix(string(bytes), "\uFEFF"))

	var file File
	if err := json.Unmarshal(bytes, &file); err != nil {
		return File{}, err
	}

	if file.Errors == nil {
		file.Errors = []Error{}
	}
	if file.Version == 0 {
		file.Version = 1
	}

	return file, nil
}
