package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// incrementPatchVersion takes a version string (e.g., "v1.2.3") and returns the
// version with the patch number incremented (e.g., "v1.2.4"). It expects a "v" prefix.
func incrementPatchVersion(version string) (string, error) {
	// Ensure the version starts with 'v'
	if !strings.HasPrefix(version, "v") {
		return "", fmt.Errorf("version must start with 'v', got: %s", version)
	}
	// Remove the "v" prefix for easier parsing
	versionWithoutV := strings.TrimPrefix(version, "v")

	parts := strings.Split(versionWithoutV, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format: %s. Expected vX.Y.Z", version)
	}

	// Parse major, minor, and patch versions
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid major version '%s' in %s", parts[0], version)
	}
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid minor version '%s' in %s", parts[1], version)
	}
	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", fmt.Errorf("invalid patch version '%s' in %s", parts[2], version)
	}

	// Increment the patch version
	newPatch := patch + 1
	// Format the new version back with the "v" prefix
	return fmt.Sprintf("v%d.%d.%d", major, minor, newPatch), nil
}

// prependToFile reads the content of a file, prepends the new content, and writes it back.
// If the file does not exist, it will be created.
func prependToFile(filePath, contentToPrepend string) error {
	// Read existing content. If file doesn't exist, ReadFile returns an error
	// but we can proceed if it's an "is not exist" error.
	existingContent, err := ioutil.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// Combine new content with existing content, ensuring a newline between them
	// and a final newline at the end for good measure.
	newContent := []byte(contentToPrepend + "\n" + string(existingContent))

	// Write the combined content back to the file. 0644 sets read/write permissions for owner,
	// and read-only for others.
	return ioutil.WriteFile(filePath, newContent, 0644)
}

func main() {
	// Check if a changelog message is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generate.go \"<Your changelog message>\"")
		fmt.Println("Example: go run generate.go \"Fixed a minor bug in user authentication\"")
		os.Exit(1)
	}
	changelogMessage := os.Args[1] // The message for the changelog entry

	// Define paths relative to the script's execution directory (expected to be monorepo root)
	rootChangelogPath := "CHANGELOG.md"
	servicesDirPath := "services"

	// Get the current working directory to ensure all paths are absolute and correct
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	// Construct the full path to the services directory
	fullServicesDirPath := filepath.Join(cwd, servicesDirPath)

	// Read all entries in the services directory
	entries, err := ioutil.ReadDir(fullServicesDirPath)
	if err != nil {
		fmt.Printf("Error reading service directory %s: %v\n", fullServicesDirPath, err)
		os.Exit(1)
	}
	reverseSlice(entries)

	fmt.Println("Starting service version and changelog updates...")
	fmt.Println("-------------------------------------------------")

	// Iterate over each entry found in the services directory
	for _, entry := range entries {
		// Process only directories that are not hidden (e.g., .git)
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			serviceName := entry.Name()
			servicePath := filepath.Join(fullServicesDirPath, serviceName)
			serviceVersionFilePath := filepath.Join(servicePath, "VERSION")
			serviceChangelogFilePath := filepath.Join(servicePath, "CHANGELOG.md")

			fmt.Printf("Processing service: %s\n", serviceName)

			// 1. Read the current version from the service's VERSION file
			currentVersionBytes, err := ioutil.ReadFile(serviceVersionFilePath)
			if err != nil {
				fmt.Printf("  Error reading version file for %s (%s): %v. Skipping service.\n", serviceName, serviceVersionFilePath, err)
				continue // Move to the next service
			}
			currentVersion := strings.TrimSpace(string(currentVersionBytes))

			// 2. Increment the patch version
			newVersion, err := incrementPatchVersion(currentVersion)
			if err != nil {
				fmt.Printf("  Error incrementing version for %s (current: %s): %v. Skipping service.\n", serviceName, currentVersion, err)
				continue // Move to the next service
			}
			fmt.Printf("  Version bumped from %s to %s\n", currentVersion, newVersion)

			// 3. Write the new version back to the service's VERSION file
			err = ioutil.WriteFile(serviceVersionFilePath, []byte(newVersion), 0644)
			if err != nil {
				fmt.Printf("  Error writing new version to %s: %v. Skipping service.\n", serviceVersionFilePath, err)
				continue // Move to the next service
			}

			// 4. Add the new changelog entry to the service's CHANGELOG.md
			serviceChangelogEntry := fmt.Sprintf("## %s\n  - %s\n", newVersion, changelogMessage)
			err = prependToFile(serviceChangelogFilePath, serviceChangelogEntry)
			if err != nil {
				fmt.Printf("  Error adding changelog entry to %s: %v. Skipping service.\n", serviceChangelogFilePath, err)
				continue // Move to the next service
			}
			fmt.Printf("  Added changelog entry to %s\n", serviceChangelogFilePath)

			// 5. Add the new changelog entry to the root CHANGELOG.md
			rootChangelogEntry := fmt.Sprintf("- `%s`: [%s](services/%s/CHANGELOG.md#%s) \n  - %s", serviceName, newVersion, serviceName, strings.ReplaceAll(newVersion, ".", ""), changelogMessage)
			err = prependToFile(filepath.Join(cwd, rootChangelogPath), rootChangelogEntry)
			if err != nil {
				fmt.Printf("  Error adding changelog entry to root %s: %v. Skipping service.\n", filepath.Join(cwd, rootChangelogPath), err)
				continue // Move to the next service
			}
			fmt.Printf("  Added changelog entry to root %s\n", filepath.Join(cwd, rootChangelogPath))
			fmt.Println("-------------------------------------------------")
		}
	}

	fmt.Println("\nScript finished successfully!")
}

func reverseSlice[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

