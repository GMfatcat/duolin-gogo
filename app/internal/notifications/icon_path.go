package notifications

import (
	"os"
	"path/filepath"
)

func resolveIconPath(executablePathFunc func() (string, error), workingDirFunc func() (string, error)) string {
	candidates := []string{}

	if executablePath, err := executablePathFunc(); err == nil && executablePath != "" {
		candidates = append(candidates, filepath.Dir(executablePath))
	}

	if workingDir, err := workingDirFunc(); err == nil && workingDir != "" {
		candidates = append(candidates, workingDir)
	}

	for _, base := range candidates {
		for _, root := range walkUp(base) {
			for _, relative := range []string{
				filepath.Join("build", "windows", "icon.ico"),
				filepath.Join("app", "build", "windows", "icon.ico"),
			} {
				iconPath := filepath.Join(root, relative)
				if fileExists(iconPath) {
					return iconPath
				}
			}
		}
	}

	return ""
}

func defaultIconPath() string {
	return resolveIconPath(os.Executable, os.Getwd)
}

func walkUp(start string) []string {
	current := filepath.Clean(start)
	var paths []string

	for {
		paths = append(paths, current)
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	return paths
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
