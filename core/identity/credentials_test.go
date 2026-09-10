package identity

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// captureLogs installs a logger for the duration of the test and returns its buffer.
func captureLogs(t *testing.T) *bytes.Buffer {
	t.Helper()
	buf := &bytes.Buffer{}
	SetLogger(slog.New(slog.NewTextHandler(buf, &slog.HandlerOptions{Level: slog.LevelWarn})))
	t.Cleanup(func() { SetLogger(nil) })
	return buf
}

func writeCredentialsFile(t *testing.T, perm os.FileMode) string {
	t.Helper()
	//nolint:gosec // G101 False positive: This is a test JWT token, not a credential
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjk5OTk5OTk5OTl9.test"
	content, err := json.Marshal(Credentials{ServiceAccountToken: token})
	if err != nil {
		t.Fatalf("marshal credentials: %v", err)
	}
	path := filepath.Join(t.TempDir(), "credentials.json")
	if err := os.WriteFile(path, content, perm); err != nil {
		t.Fatalf("write credentials file: %v", err)
	}
	// WriteFile applies the umask, so set the mode explicitly.
	if err := os.Chmod(path, perm); err != nil {
		t.Fatalf("chmod credentials file: %v", err)
	}
	return path
}

func TestReadCredentialsFileWarnsOnLoosePermissions(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("permission bits are synthesised on Windows")
	}
	logs := captureLogs(t)

	path := writeCredentialsFile(t, 0o644)
	if _, err := ReadCredentialsFile(path); err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	if !strings.Contains(logs.String(), "readable by other users") {
		t.Fatalf("expected a warning for a world-readable credentials file, got: %q", logs.String())
	}
}

func TestReadCredentialsFileSilentOnTightPermissions(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("permission bits are synthesised on Windows")
	}
	logs := captureLogs(t)

	path := writeCredentialsFile(t, 0o600)
	if _, err := ReadCredentialsFile(path); err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	if logs.Len() != 0 {
		t.Fatalf("expected no warning for a 0600 credentials file, got: %q", logs.String())
	}
}

// TestReadCredentialsFileWarningDoesNotLeakToken guards the security property that matters:
// the warning must name the file, never its contents.
func TestReadCredentialsFileWarningDoesNotLeakToken(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("permission bits are synthesised on Windows")
	}
	logs := captureLogs(t)

	path := writeCredentialsFile(t, 0o644)
	credentials, err := ReadCredentialsFile(path)
	if err != nil {
		t.Fatalf("expected no error: %v", err)
	}

	if strings.Contains(logs.String(), credentials.ServiceAccountToken) {
		t.Fatalf("the warning leaked the token")
	}
}
