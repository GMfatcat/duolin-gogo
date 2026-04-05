package notifications

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveIconPathFindsWindowsIconAlongExecutableTree(t *testing.T) {
	root := t.TempDir()
	iconPath := filepath.Join(root, "app", "build", "windows", "icon.ico")
	if err := os.MkdirAll(filepath.Dir(iconPath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(iconPath, []byte("icon"), 0o644); err != nil {
		t.Fatalf("write icon failed: %v", err)
	}

	resolved := resolveIconPath(func() (string, error) {
		return filepath.Join(root, "app", "build", "bin", "app.exe"), nil
	}, func() (string, error) {
		return "", os.ErrNotExist
	})

	if resolved != iconPath {
		t.Fatalf("expected %q, got %q", iconPath, resolved)
	}
}

func TestResolveIconPathFallsBackToWorkingDirectory(t *testing.T) {
	root := t.TempDir()
	iconPath := filepath.Join(root, "build", "windows", "icon.ico")
	if err := os.MkdirAll(filepath.Dir(iconPath), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(iconPath, []byte("icon"), 0o644); err != nil {
		t.Fatalf("write icon failed: %v", err)
	}

	resolved := resolveIconPath(func() (string, error) {
		return "", os.ErrNotExist
	}, func() (string, error) {
		return root, nil
	})

	if resolved != iconPath {
		t.Fatalf("expected %q, got %q", iconPath, resolved)
	}
}
