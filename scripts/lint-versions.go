package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	targetFileName = "VERSION"
)

var (
	targetDirs = []string{"services", "core"}
)

func main() {
	fmt.Println(">> Linting SDK module versions...")

	// e.g., versions["services/iaas"] = "v1.2.3"
	versions := make(map[string]string)

	for _, rootDir := range targetDirs {
		err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("error accessing path %s: %w", path, err)
			}

			if d.IsDir() && d.Name() == ".openapi-generator" {
				return filepath.SkipDir
			}

			if !d.IsDir() && d.Name() == targetFileName {
				content, err := os.ReadFile(path)
				if err != nil {
					return fmt.Errorf("could not read file %s: %v", path, err)
				}

				dirPath := filepath.Dir(path)
				version := strings.TrimSpace(string(content))

				versions[dirPath] = version
			}

			return nil
		})

		if err != nil {
			log.Fatalf("failed to walk directory %s: %v", rootDir, err)
		}
	}

	re := regexp.MustCompile(`^v[0-1]\.\d+\.\d+$`)
	for dir, version := range versions {
		versionOk := re.MatchString(version)

		if !versionOk {
			log.Fatalf("Invalid version for %s: %s\n", dir, version)
		}
	}
}
