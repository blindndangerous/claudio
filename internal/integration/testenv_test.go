package integration

import (
	"path/filepath"
	"strings"
	"testing"

	clitestenv "claudio.click/internal/cli/testenv"
	"claudio.click/internal/tracking"
)

func isolateIntegrationXDG(t *testing.T) string {
	t.Helper()
	return clitestenv.IsolateXDG(t)
}

func TestIntegrationEnvironmentUsesTemporaryXDG(t *testing.T) {
	root := isolateIntegrationXDG(t)

	databasePath, err := tracking.GetDatabasePath()
	if err != nil {
		t.Fatalf("GetDatabasePath: %v", err)
	}
	wantPrefix := filepath.Join(root, ".cache") + string(filepath.Separator)
	if !strings.HasPrefix(databasePath, wantPrefix) {
		t.Fatalf("tracking database path %q is outside isolated cache %q", databasePath, wantPrefix)
	}
}
