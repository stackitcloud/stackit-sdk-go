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
		coreOnly           bool
		expectedLatestTags map[string]string
	}{
		{
			desc: "valid service latest version",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.2.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			coreOnly: false,
			expectedLatestTags: map[string]string{
				"foo": "v0.2.0",
				"bar": "v0.1.0",
			},
		},
		{
			desc: "valid service no update",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.1.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			coreOnly: false,
			expectedLatestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
		},
		{
			desc: "valid core latest version",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/core/v0.2.0", ""),
			latestTags: map[string]string{
				"foo":  "v0.1.0",
				"bar":  "v0.1.0",
				"core": "v0.1.0",
			},
			coreOnly: true,
			expectedLatestTags: map[string]string{
				"foo":  "v0.1.0",
				"bar":  "v0.1.0",
				"core": "v0.2.0",
			},
		},
		{
			desc: "valid core latest version",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/core/v0.2.0", ""),
			latestTags: map[string]string{
				"foo":  "v0.1.0",
				"bar":  "v0.1.0",
				"core": "v0.1.0",
			},
			coreOnly: true,
			expectedLatestTags: map[string]string{
				"foo":  "v0.1.0",
				"bar":  "v0.1.0",
				"core": "v0.2.0",
			},
		},
		{
			desc: "invalid service tag",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/foo/v0.2.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			coreOnly: false,
			expectedLatestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
		},
		{
			desc: "valid service tag but coreOnly",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.2.0", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			coreOnly: true,
			expectedLatestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
		},
		{
			desc: "dont update to pre-release",
			tag:  plumbing.NewReferenceFromStrings("refs/tags/services/foo/v0.2.0-pre-release", ""),
			latestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
			coreOnly: true,
			expectedLatestTags: map[string]string{
				"foo": "v0.1.0",
				"bar": "v0.1.0",
			},
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			updatedLatestTags := storeLatestTag(test.tag, test.latestTags, test.coreOnly)
			if !reflect.DeepEqual(test.expectedLatestTags, updatedLatestTags) {
				t.Fatalf("Updated tags are wrong, expected %v, got %v", test.expectedLatestTags, updatedLatestTags)
			}
		})
	}
}

func TestComputeUpdatedTag(t *testing.T) {
	for _, test := range []struct {
		desc        string
		service     string
		version     string
		updateType  string
		coreOnly    bool
		expectedTag string
		isValid     bool
	}{
		{
			desc:        "valid service minor update",
			service:     "foo",
			version:     "v0.1.0",
			updateType:  "minor",
			expectedTag: "services/foo/v0.2.0",
			isValid:     true,
		},
		{
			desc:        "valid service patch update",
			service:     "foo",
			version:     "v0.1.0",
			updateType:  "patch",
			expectedTag: "services/foo/v0.1.1",
			isValid:     true,
		},
		{
			desc:        "valid core minor update",
			service:     "core",
			version:     "v0.1.0",
			updateType:  "minor",
			coreOnly:    true,
			expectedTag: "core/v0.2.0",
			isValid:     true,
		},
		{
			desc:        "valid core patch update",
			service:     "core",
			version:     "v0.1.0",
			updateType:  "patch",
			coreOnly:    true,
			expectedTag: "core/v0.1.1",
			isValid:     true,
		},
		{
			desc:       "invalid update type",
			service:    "foo",
			version:    "v0.1.0",
			updateType: "major",
			isValid:    false,
		},
		{
			desc:       "invalid core update with wrong name - this should never happen",
			service:    "foo",
			version:    "v0.1.0",
			updateType: "patch",
			coreOnly:   true,
			isValid:    false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			updatedTag, err := computeUpdatedTag(test.service, test.version, test.updateType, test.coreOnly)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && test.expectedTag != updatedTag {
				t.Fatalf("updated tag is wrong: expected %s, got %s", test.expectedTag, updatedTag)
			}
		})
	}
}
