package main

import (
	"errors"
	"flag"
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
	sdkRepo     = "git@github.com:vicentepinto98/gh.git"
	patch       = "patch"
	minor       = "minor"
	allServices = "all-services"
	core        = "core"

	updateTypeFlag            = "update-type"
	sshPrivateKeyFilePathFlag = "ssh-private-key-file-path"
	passwordFlag              = "password"
	targetFlag                = "target"
)

var (
	updateTypes = []string{minor, patch}
	targets     = []string{allServices, core}
	usage       = "go run automatic_tag.go --update-type [minor|patch] --ssh-private-key-file-path path/to/private-key --password password --target [all-services|core]"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var updateType string
	var sshPrivateKeyFilePath string
	var password string
	var target string

	flag.StringVar(&updateType, updateTypeFlag, "", fmt.Sprintf("Update type, must be one of: %s (required)", strings.Join(updateTypes, ",")))
	flag.StringVar(&sshPrivateKeyFilePath, sshPrivateKeyFilePathFlag, "", "Path to the ssh private key (required)")
	flag.StringVar(&password, passwordFlag, "", "Password of the ssh private key (optional)")
	flag.StringVar(&target, targetFlag, allServices, fmt.Sprintf("Create tags for this target, must be one of %s (optional, default is %s)", strings.Join(targets, ","), allServices))

	flag.Parse()

	validUpdateType := false
	for _, t := range updateTypes {
		if updateType == t {
			validUpdateType = true
			break
		}
	}
	if !validUpdateType {
		return fmt.Errorf("the provided update type `%s` is not valid, the valid values are: [%s]", updateType, strings.Join(updateTypes, ","))
	}

	validTarget := false
	for _, t := range targets {
		if target == t {
			validTarget = true
			break
		}
	}
	if !validTarget {
		return fmt.Errorf("the provided target `%s` is not valid, the valid values are: [%s]", target, strings.Join(targets, ","))
	}

	_, err := os.Stat(sshPrivateKeyFilePath)
	if err != nil {
		return fmt.Errorf("the provided private key file path %s is not valid: %s\nUsage: %s", sshPrivateKeyFilePath, err, usage)
	}

	err = automaticTagUpdate(updateType, sshPrivateKeyFilePath, password, target)
	if err != nil {
		return fmt.Errorf("updating tags: %s", err.Error())
	}
	return nil
}

// automaticTagUpdate goes through all of the exising tags, gets the latest for the targer, creates a new one according to the updateType and pushes them
func automaticTagUpdate(updateType, sshPrivateKeyFilePath, password, target string) error {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return fmt.Errorf("create temporary directory: %w", err)
	}

	defer func() {
		tempErr := os.RemoveAll(tempDir)
		if tempErr != nil {
			fmt.Printf("Warning: temporary directory %s could not be removed: %s", tempDir, tempErr.Error())
		}
	}()

	publicKeys, err := ssh.NewPublicKeysFromFile("git", sshPrivateKeyFilePath, password)
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
		latestTags, err = storeLatestTag(t, latestTags, target)
		if err != nil {
			return fmt.Errorf("store latest tag: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("iterate over existing tags: %w", err)
	}

	for service, version := range latestTags {
		newTag, err := computeUpdatedTag(service, version, updateType, target)
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

// storeLatestTag receives a tag in the form of a plumbing.Reference and a map with the latest tag per service
// It checks if the tag is part of the current target (if it is belonging to a service or to core),
// checks if it is newer than the current latest tag stored in the map and if it is, updates latestTags and returns it
func storeLatestTag(t *plumbing.Reference, latestTags map[string]string, target string) (map[string]string, error) {
	tagName, _ := strings.CutPrefix(t.Name().String(), "refs/tags/")
	splitTag := strings.Split(tagName, "/")

	switch target {
	case core:
		if len(splitTag) != 2 || splitTag[0] != "core" {
			return latestTags, nil
		}

		version := splitTag[1]
		if semver.Prerelease(version) != "" {
			return latestTags, nil
		}

		// invalid (or empty) semantic version are considered less than a valid one
		if semver.Compare(latestTags["core"], version) == -1 {
			latestTags["core"] = version
		}
	case allServices:
		if len(splitTag) != 3 || splitTag[0] != "services" {
			return latestTags, nil
		}

		service := splitTag[1]
		version := splitTag[2]
		if semver.Prerelease(version) != "" {
			return latestTags, nil
		}

		// invalid (or empty) semantic version are considered less than a valid one
		if semver.Compare(latestTags[service], version) == -1 {
			latestTags[service] = version
		}
	default:
		return nil, fmt.Errorf("target not supported in storeLatestTag, fix the script")
	}
	return latestTags, nil
}

// computeUpdatedTag returns the a new tag with the updated version for a specific service according to the update type and target
// example: for service argus with version v0.1.1, target all-services and updateType minor, it returns services/argus/v0.2.0
func computeUpdatedTag(service, version, updateType, target string) (string, error) {
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

	switch target {
	case core:
		if service != "core" {
			return "", fmt.Errorf("%s target was provided but store latest tag from another service: %s", target, service)
		}
		return fmt.Sprintf("core/%s", updatedVersion), nil
	case allServices:
		return fmt.Sprintf("services/%s/%s", service, updatedVersion), nil
	default:
		return "", fmt.Errorf("target not supported in version increment, fix the script")
	}
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
