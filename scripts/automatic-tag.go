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
	sdkRepo = "git@github.com:stackitcloud/stackit-sdk-go.git"
	patch   = "patch"
	minor   = "minor"
)

var (
	updateTypes = []string{minor, patch}
)

func main() {
	// Parse arguments
	if len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Println("wrong number of arguments. Usage: go run automatic-tag.go [minor|patch] private-key-file-path --core-only")
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
		fmt.Printf("the provided private key file path %s is not valid: %s\nUsage: go run automatic-tag.go [minor|patch] private-key-file-path --core-only\n", privateKeyFile, err)
		os.Exit(1)
	}

	// Parse flag
	var coreOnly bool
	if len(os.Args) == 4 {
		if os.Args[3] != "--core-only" {
			fmt.Println("wrong arguments. Usage: go run automatic-tag.go [minor|patch] --core-only")
			os.Exit(1)
		}
		coreOnly = true
	}

	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		fmt.Printf("create temporary directory: %s", err.Error())
		os.Exit(1)
	}

	defer func() {
		tempErr := os.RemoveAll(tempDir)
		if tempErr != nil {
			fmt.Printf("temporary directory %s could not be removed: %s", tempDir, tempErr.Error())
		}
	}()

	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, "")
	if err != nil {
		fmt.Printf("get public keys from private key file: %s\n", err.Error())
		os.Exit(1)
	}

	r, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		Auth:              publicKeys,
		URL:               sdkRepo,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		fmt.Printf("clone SDK repo: %s\n", err.Error())
		os.Exit(1)
	}

	tagrefs, err := r.Tags()
	if err != nil {
		fmt.Printf("get tags: %s", err.Error())
		os.Exit(1)
	}

	existingTags := map[string]string{}
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		tagName, _ := strings.CutPrefix(t.Name().String(), "refs/tags/")
		splitTag := strings.Split(tagName, "/")

		if coreOnly {
			if len(splitTag) != 2 || splitTag[0] != "core" {
				return nil
			}

			version := splitTag[1]
			if semver.Prerelease(version) != "" {
				return nil
			}

			// invalid (or empty) semantic version are considered less than a valid one
			if semver.Compare(existingTags["core"], version) == -1 {
				existingTags["core"] = version
			}
		} else {
			if len(splitTag) != 3 || splitTag[0] != "services" {
				return nil
			}

			service := splitTag[1]
			version := splitTag[2]
			if semver.Prerelease(version) != "" {
				return nil
			}

			// invalid (or empty) semantic version are considered less than a valid one
			if semver.Compare(existingTags[service], version) == -1 {
				existingTags[service] = version
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("iterate over existing tags: %s\n", err.Error())
		os.Exit(1)
	}

	for service, version := range existingTags {
		newTag, err := incrementTag(service, version, updateType, coreOnly)
		if err != nil {
			fmt.Printf("Error for %s with version %s, this tag will be skipped\n", service, version)
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
		fmt.Printf("push tags: %s\n", err)
		os.Exit(1)
	}
}

func incrementTag(service, version, updateType string, coreOnly bool) (string, error) {
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
		fmt.Println("Update type not supported in version increment, fix the script")
		os.Exit(1)
	}

	updatedVersion := strings.Join(splitVersion, ".")
	var newTag string
	if coreOnly {
		newTag = fmt.Sprintf("core/%s", updatedVersion)
	} else {
		newTag = fmt.Sprintf("services/%s/%s", service, updatedVersion)
	}
	return newTag, nil
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
