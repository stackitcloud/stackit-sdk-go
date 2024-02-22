package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/mod/semver"
)

const (
	//sdkRepo = "git@github.com:stackitcloud/stackit-sdk-go.git"
	sdkRepo = "git@github.com:vicentepinto98/gh.git"
	patch   = "patch"
	minor   = "minor"
	usage   = "go run automatic_tag.go [minor|patch] private-key-file-path --core-only"
)

var (
	updateTypes = []string{minor, patch}
)

func main() {
	// Parse arguments
	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Printf("wrong number of arguments. Usage: %s\n", usage)
		os.Exit(1)
	}

	updateType := os.Args[1]
	valid := false
	for _, t := range updateTypes {
		if updateType == t {
			valid = true
			break
		}
	}
	if !valid {
		fmt.Printf("the provided argument for update type `%s` is not valid, the valid values are: [%s]\n", updateType, strings.Join(updateTypes, ","))
		os.Exit(1)
	}

	privateKeyFile := os.Args[2]
	_, err := os.Stat(privateKeyFile)
	if err != nil {
		fmt.Printf("The provided private key file path %s is not valid: %s\nUsage: %s\n", privateKeyFile, err, usage)
		os.Exit(1)
	}

	// Parse flag
	var coreOnly bool
	if len(os.Args) == 4 {
		if os.Args[3] != "--core-only" {
			fmt.Printf("Wrong arguments. Usage: %s\n", usage)
			os.Exit(1)
		}
		coreOnly = true
	}

	err = automaticTagUpdate(updateType, privateKeyFile, coreOnly)
	if err != nil {
		fmt.Printf("Error updating tags: %s\n", err.Error())
		os.Exit(1)
	}
}

func automaticTagUpdate(updateType, privateKeyFile string, coreOnly bool) error {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return fmt.Errorf("create temporary directory: %w", err)
	}

	defer func() {
		tempErr := os.RemoveAll(tempDir)
		if tempErr != nil && err == nil {
			err = fmt.Errorf("temporary directory %s could not be removed: %w", tempDir, tempErr)
		}
	}()

	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, "")
	if err != nil {
		return fmt.Errorf("get public keys from private key file: %w", err)
	}

	r, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		Auth:              publicKeys,
		URL:               sdkRepo,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		return fmt.Errorf("clone SDK repo: %w", err)
	}

	tagrefs, err := r.Tags()
	if err != nil {
		return fmt.Errorf("get tags: %w", err)
	}

	latestTags := map[string]string{}
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		latestTags = storeLatestTag(t, latestTags, coreOnly)
		return nil
	})
	if err != nil {
		return fmt.Errorf("iterate over existing tags: %w", err)
	}

	for service, version := range latestTags {
		newTag, err := computeUpdatedTag(service, version, updateType, coreOnly)
		if err != nil {
			fmt.Printf("Error for %s with version %s, this tag will be skipped. Error: %s\n", service, version, err.Error())
			continue
		}

		err = createTag(r, newTag)
		if err != nil {
			fmt.Printf("Create tag %s returned error: %s\n", newTag, err)
			continue
		}
		fmt.Printf("Created tag %s\n", newTag)
	}

	err = pushTags(r, publicKeys)
	if err != nil {
		return fmt.Errorf("push tags: %w", err)
	}
	return nil
}

func storeLatestTag(t *plumbing.Reference, latestTags map[string]string, coreOnly bool) map[string]string {
	tagName, _ := strings.CutPrefix(t.Name().String(), "refs/tags/")
	splitTag := strings.Split(tagName, "/")

	if coreOnly {
		if len(splitTag) != 2 || splitTag[0] != "core" {
			return latestTags
		}

		version := splitTag[1]
		if semver.Prerelease(version) != "" {
			return latestTags
		}

		// invalid (or empty) semantic version are considered less than a valid one
		if semver.Compare(latestTags["core"], version) == -1 {
			latestTags["core"] = version
		}
	} else {
		if len(splitTag) != 3 || splitTag[0] != "services" {
			return latestTags
		}

		service := splitTag[1]
		version := splitTag[2]
		if semver.Prerelease(version) != "" {
			return latestTags
		}

		// invalid (or empty) semantic version are considered less than a valid one
		if semver.Compare(latestTags[service], version) == -1 {
			latestTags[service] = version
		}
	}
	return latestTags
}

func computeUpdatedTag(service, version, updateType string, coreOnly bool) (string, error) {
	canonicalVersion := semver.Canonical(version)
	splitVersion := strings.Split(canonicalVersion, ".")
	if len(splitVersion) != 3 {
		return "", fmt.Errorf("invalid canonical version")
	}

	switch updateType {
	case patch:
		patchNumber, err := strconv.Atoi(splitVersion[2])
		if err != nil {
			return "", fmt.Errorf("couldnt convert patch number to int")
		}
		updatedPatchNumber := patchNumber + 1
		splitVersion[2] = fmt.Sprint(updatedPatchNumber)
	case minor:
		minorNumber, err := strconv.Atoi(splitVersion[1])
		if err != nil {
			return "", fmt.Errorf("couldnt convert minor number to int")
		}
		updatedPatchNumber := minorNumber + 1
		splitVersion[1] = fmt.Sprint(updatedPatchNumber)
		splitVersion[2] = "0"
	default:
		return "", fmt.Errorf("update type not supported in version increment, fix the script")
	}

	updatedVersion := strings.Join(splitVersion, ".")
	if coreOnly {
		if service != "core" {
			return "", fmt.Errorf("--core-only was provided but store latest tag from another service: %s", service)
		}
		return fmt.Sprintf("core/%s", updatedVersion), nil
	}

	return fmt.Sprintf("services/%s/%s", service, updatedVersion), nil
}

func createTag(r *git.Repository, tag string) error {
	h, err := r.Head()
	if err != nil {
		return fmt.Errorf("get HEAD: %w", err)
	}
	_, err = r.CreateTag(tag, h.Hash(), nil)
	if err != nil {
		return fmt.Errorf("create tag: %w", err)
	}
	return nil
}

func pushTags(r *git.Repository, publicKeys *ssh.PublicKeys) error {
	po := &git.PushOptions{
		Auth:       publicKeys,
		RemoteName: "origin",
		Progress:   os.Stdout,
		RefSpecs:   []config.RefSpec{config.RefSpec("refs/tags/*:refs/tags/*")},
	}
	err := r.Push(po)

	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			return fmt.Errorf("origin remote was up to date, no push done")
		}
		return fmt.Errorf("push to remote origin: %w", err)
	}

	return nil
}
