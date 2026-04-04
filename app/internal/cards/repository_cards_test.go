package cards

import (
	"path/filepath"
	"slices"
	"testing"
)

func TestRepositoryKnowledgeCardsImportWithoutErrors(t *testing.T) {
	knowledgeDir := filepath.Clean(filepath.Join("..", "..", "..", "knowledge"))

	result, err := ScanDirectories([]string{knowledgeDir})
	if err != nil {
		t.Fatalf("scan failed: %v", err)
	}

	if len(result.Errors) != 0 {
		t.Fatalf("expected no import errors, got %d", len(result.Errors))
	}

	expectedIDs := []string{
		"git-add-staging",
		"git-checkout-legacy",
		"git-cherry-pick-purpose",
		"git-commit-snapshot",
		"git-fetch-basic",
		"git-merge-purpose",
		"git-pull-composition",
		"git-rebase-vs-merge",
		"git-restore-discard",
		"git-stash-purpose",
		"git-status-purpose",
		"git-switch-branch",
	}

	importedIDs := make([]string, 0, len(result.Cards))
	for _, card := range result.Cards {
		importedIDs = append(importedIDs, card.ID)
	}

	for _, expectedID := range expectedIDs {
		if !slices.Contains(importedIDs, expectedID) {
			t.Fatalf("expected repository card %q to exist", expectedID)
		}
	}
}
