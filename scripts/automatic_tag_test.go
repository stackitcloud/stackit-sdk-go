package main

import (
	"reflect"
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
)

func TestStoreLatestTag(t *testing.T) {
	for _, test := range []struct {
		desc               string
		tag                *plumbing.Reference
		latestTags         map[string]string
		target             string
		expectedLatestTags map[string]string
		isValid            bool
	}{
		{
			desc: "valid service latest version",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.2.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			target: "all-services",
			expectedLatestTags: map[string]string{
				"foo": "v0.2.0",
				"bar": "v0.1.0",
			},
			isValid: true,
		},
		{
			desc: "valid service no update",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.1.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			target: "all-services",
			expectedLatestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			isValid: true,
		},
		{
			desc: "valid core latest version",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/core/v0.2.0", ""),
			latestTags: map[string]string{
				"foo":  "v0.1.0",
				"bar":  "v0.1.0",
				"core": "v0.1.0",
			},
			target: "core",
			expectedLatestTags: map[string]string{
				"foo":  "v0.1.0",
				"bar":  "v0.1.0",
				"core": "v0.2.0",
			},
			isValid: true,
		},
		{
			desc: "invalid service tag",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/foo/v0.2.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			target: "all-services",
			expectedLatestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			isValid: true,
		},
		{
			desc: "valid service tag but coreOnly",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.2.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			target: "core",
			expectedLatestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			isValid: true,
		},
		{
			desc: "dont update to pre-release",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.2.0-pre-release", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			target: "core",
			expectedLatestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			isValid: true,
		},
		{
			desc: "unsupported target",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.2.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			target:             "unsupported",
			expectedLatestTags: map[string]string{},
			isValid:            false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			updatedLatestTags, err := storeLatestTag(test.tag, test.latestTags, test.target)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && !reflect.DeepEqual(test.expectedLatestTags, updatedLatestTags) {
				t.Fatalf("Updated tags are wrong, expected %v, got %v", test.expectedLatestTags, updatedLatestTags)
			}
		})
	}
}

func TestComputeUpdatedVersion(t *testing.T) {
	for _, test := range []struct {
		desc            string
		version         string
		updateType      string
		expectedVersion string
		isValid         bool
	}{
		{
			desc:            "valid minor update",
			version:         "v0.1.0",
			updateType:      "minor",
			expectedVersion: "v0.2.0",
			isValid:         true,
		},
		{
			desc:            "valid minor update with previous patch version not zero",
			version:         "v0.1.1",
			updateType:      "minor",
			expectedVersion: "v0.2.0",
			isValid:         true,
		},
		{
			desc:            "valid patch update",
			version:         "v0.1.0",
			updateType:      "patch",
			expectedVersion: "v0.1.1",
			isValid:         true,
		},
		{
			desc:       "invalid update type",
			version:    "v0.1.0",
			updateType: "major",
			isValid:    false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			updatedVersion, err := computeUpdatedVersion(test.version, test.updateType)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && test.expectedVersion != updatedVersion {
				t.Fatalf("updated tag is wrong: expected %s, got %s", test.expectedVersion, updatedVersion)
			}
		})
	}
}
