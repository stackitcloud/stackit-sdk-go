package credentials

import (
	"os"
	"sync"
)

// credentialsFileSource reads the credentials file lazily and at most once, so that
// resolving several values (and several providers in a chain) does not read the file, or
// warn about its permissions, repeatedly. A file that cannot be read yields no values.
type credentialsFileSource struct {
	path  string
	once  sync.Once
	creds *credentialsFile
}

func newCredentialsFileSource(path string) *credentialsFileSource {
	return &credentialsFileSource{path: path}
}

// value returns the field selected by pick, or "" when the file is unavailable.
func (s *credentialsFileSource) value(pick func(*credentialsFile) string) string {
	s.once.Do(func() {
		if creds, err := readCredentialsFile(s.path); err == nil {
			s.creds = creds
		}
	})
	if s.creds == nil {
		return ""
	}
	return pick(s.creds)
}

// fileSourceFor returns shared, or a new source for path when the provider is built on its own.
func fileSourceFor(shared *credentialsFileSource, path string) *credentialsFileSource {
	if shared != nil {
		return shared
	}
	return newCredentialsFileSource(path)
}

// resolveValue returns the first non-empty value among, in order: explicit, the contents of
// the file at explicitPath, the environment variable envVar, the contents of the file named
// by envPathVar, and the fallbacks. Empty explicitPath and envPathVar are skipped. It returns
// "" when nothing is set. An error is returned only when a path that is set cannot be read.
func resolveValue(explicit, explicitPath, envVar, envPathVar string, fallbacks ...func() string) (string, error) {
	if explicit != "" {
		return explicit, nil
	}
	if explicitPath != "" {
		data, err := os.ReadFile(explicitPath)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	if value := os.Getenv(envVar); value != "" {
		return value, nil
	}
	if envPathVar != "" {
		if path := os.Getenv(envPathVar); path != "" {
			data, err := os.ReadFile(path) //nolint:gosec // G703: the path is configuration chosen by the caller
			if err != nil {
				return "", err
			}
			return string(data), nil
		}
	}
	for _, fallback := range fallbacks {
		if value := fallback(); value != "" {
			return value, nil
		}
	}
	return "", nil
}

func fromFile(s *credentialsFileSource, pick func(*credentialsFile) string) func() string {
	return func() string { return s.value(pick) }
}
