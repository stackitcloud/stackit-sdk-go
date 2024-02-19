package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"golang.org/x/mod/semver"
)

const (
	sdkRepo = "https://github.com/stackitcloud/stackit-sdk-go"
)

var (
	updateTypes = []string{"minor", "patch"}
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("wrong number of arguments. Usage: go run automatic-tag.go [minor|patch] --core-only")
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
		fmt.Printf("the provided argument `%s` is not valid, the valid values are: [%s]\n", updateType, strings.Join(updateTypes, ","))
		os.Exit(1)
	}

	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		fmt.Printf("create temporary directory: %s", err.Error())
		os.Exit(1)
	}

	r, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL: sdkRepo,
	})
	if err != nil {
		fmt.Printf("clone SDK repo: %s", err.Error())
		os.Exit(1)
	}

	tags, err := r.TagObjects()
	if err != nil {
		fmt.Printf("get tags: %s", err.Error())
		os.Exit(1)
	}

	var coreOnly bool
	if len(os.Args) == 3 {
		if os.Args[3] != "--core-only" {
			fmt.Println("wrong arguments. Usage: go run automatic-tag.go [minor|patch] --core-only")
			os.Exit(1)
		}
		coreOnly = true
	}

	existingTags := map[string]string{}
	err = tags.ForEach(func(t *object.Tag) error {
		tagName, _ := strings.CutSuffix(t.Name, "refs/tags/")
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

	for t := range existingTags {
		fmt.Println(t)
	}

	// TO-DO: Update tags
	// TO-DO: Push tags
}
