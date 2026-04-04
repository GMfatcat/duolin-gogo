package diagnostics

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadReturnsEmptyDiagnosticsWhenFileMissing(t *testing.T) {
	dir := t.TempDir()
	file, err := Load(filepath.Join(dir, "import-errors.json"))
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(file.Errors) != 0 {
		t.Fatalf("expected no diagnostics, got %d", len(file.Errors))
	}
}

func TestLoadReadsExistingDiagnostics(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "import-errors.json")
	content := `{
  "version": 1,
  "errors": [
    {
      "source_path": "D:\\duolin-gogo\\knowledge\\git\\bad.md",
      "code": "missing_language_section",
      "field": "body",
      "message": "Body must contain both ## zh-TW and ## en sections."
    }
  ]
}`

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write file failed: %v", err)
	}

	file, err := Load(path)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	if len(file.Errors) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(file.Errors))
	}
}
