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
		"docker-build-image",
		"docker-compose-down-services",
		"docker-compose-up-services",
		"docker-exec-running-container",
		"docker-images-local-images",
		"docker-logs-container-output",
		"docker-ps-running-containers",
		"docker-run-start-container",
		"git-add-staging",
		"git-branch-list",
		"git-checkout-legacy",
		"git-cherry-pick-purpose",
		"git-clone-local-copy",
		"git-commit-snapshot",
		"git-diff-compare",
		"git-fetch-basic",
		"git-init-repository",
		"git-log-history",
		"git-merge-purpose",
		"git-merge-conflict-resolution",
		"git-pull-composition",
		"git-push-upstream",
		"git-rebase-vs-merge",
		"git-rebase-continue-flow",
		"git-remote-origin",
		"git-revert-safe-undo",
		"git-reset-head",
		"git-restore-discard",
		"git-stash-purpose",
		"git-status-purpose",
		"git-switch-branch",
		"git-tag-release-marker",
		"go-channel-communication",
		"go-defer-late-execution",
		"go-goroutine-concurrency",
		"go-interface-behavior-contract",
		"go-map-key-value",
		"go-pointer-memory-address",
		"go-slice-dynamic-view",
		"go-struct-field-grouping",
		"linux-cd-change-directory",
		"linux-chmod-permissions",
		"linux-find-search-files",
		"linux-grep-search-text",
		"linux-kill-stop-process",
		"linux-ls-list-files",
		"linux-ps-process-list",
		"linux-pwd-current-directory",
		"python-decorator-wrap-function",
		"python-dict-key-value",
		"python-exception-try-except",
		"python-generator-lazy-iteration",
		"python-lambda-anonymous-function",
		"python-list-vs-tuple",
		"python-set-unique-values",
		"python-venv-isolation",
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
