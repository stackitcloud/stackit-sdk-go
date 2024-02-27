package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/mod/semver"
	"golang.org/x/term"
)

const (
	sdkRepo     = "git@github.com:stackitcloud/stackit-sdk-go.git"
	patch       = "patch"
	minor       = "minor"
	allServices = "all-services"
	core        = "core"

	updateTypeFlag            = "update-type"
	sshPrivateKeyFilePathFlag = "ssh-private-key-file-path"
	targetFlag                = "target"
)

var (
	updateTypes = []string{minor, patch}
	targets     = []string{allServices, core}
	usage       = "go run automatic_tag.go --update-type [minor|patch] --ssh-private-key-file-path path/to/private-key --target [all-services|core]"
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
	var target string

	flag.StringVar(&updateType, updateTypeFlag, "", fmt.Sprintf("Update type, must be one of: %s (required)", strings.Join(updateTypes, ",")))
	flag.StringVar(&sshPrivateKeyFilePath, sshPrivateKeyFilePathFlag, "", "Path to the ssh private key (required)")
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
		return fmt.Errorf("the provided private key file path %s is not valid: %w\nUsage: %s", sshPrivateKeyFilePath, err, usage)
	}

	password, err := promptForPassword()
	if err != nil {
		return fmt.Errorf("prompt for password: %s", err.Error())
	}

	fmt.Print("Starting automatic tag update...\n\n")

	err = automaticTagUpdate(updateType, sshPrivateKeyFilePath, password, target)
	if err != nil {
		return fmt.Errorf("update tags: %s", err.Error())
	}
	return nil
}

// Prompts the user for the ssh key password.
func promptForPassword() (string, error) {
	fmt.Print("Enter SSH key passphrase (empty for no passphrase): ")
	defer fmt.Print("\n")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("read password: %w", err)
	}
	return string(bytePassword), nil
}

// automaticTagUpdate goes through all of the existing tags, gets the latest for the target, creates a new one according to the updateType and pushes them
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

	fmt.Printf("Cloned %s successfully\n\n", sdkRepo)

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

	var newTagsList []string
	for module, version := range latestTags {
		updatedVersion, err := computeUpdatedVersion(version, updateType)
		if err != nil {
			fmt.Printf("Error computing updated version for %s with version %s, this tag will be skipped: %s\n", module, version, err.Error())
			continue
		}

		var newTag string
		switch target {
		case core:
			if module != "core" {
				return fmt.Errorf("%s target was provided but there is a stored latest tag from another service: %s", target, module)
			}
			newTag = fmt.Sprintf("core/%s", updatedVersion)
		case allServices:
			newTag = fmt.Sprintf("services/%s/%s", module, updatedVersion)
		default:
			fmt.Printf("Error computing updated version for %s with version %s, this tag will be skipped: target %s not supported in version increment, fix the script\n", module, version, target)
			continue
		}

		newTagsList = append(newTagsList, newTag)
	}

	fmt.Printf("The following tags will be created:\n%s\n\n", strings.Join(newTagsList, "\n"))
	err = promptForConfirmation("Do you want to continue?")
	if err != nil {
		return fmt.Errorf("ask for confirmation: %w", err)
	}

	for _, newTag := range newTagsList {
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
	fmt.Print("\nTags were pushed successfully!\n")
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

// computeUpdatedVersion returns the updated version according to the update type
// example: for version v0.1.1 and updateType minor, it returns v0.2.0
func computeUpdatedVersion(version, updateType string) (string, error) {
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
	return updatedVersion, nil
}

// Prompts for confirmation.
//
// Returns nil only if the user (explicitly) answers positive.
// Returns error if the user answers negative.
func promptForConfirmation(prompt string) error {
	question := fmt.Sprintf("%s [y/N] ", prompt)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 3; i++ {
		fmt.Print(question)
		answer, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		answer = strings.ToLower(strings.TrimSpace(answer))
		if answer == "y" || answer == "yes" {
			return nil
		}
		if answer == "" || answer == "n" || answer == "no" {
			return errors.New("execution aborted")
		}
	}
	return fmt.Errorf("max number of wrong inputs")
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
