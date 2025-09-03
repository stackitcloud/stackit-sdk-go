#!/usr/bin/env bash

# Immediate exit on failure
set -e

# VERSION contains only one single line which contains the version of the service
# in the following format e.g. v0.3.0

# Check all version files which have changed
for file in $(git diff --name-only HEAD~1..HEAD | grep -E "(^services/[^/]+/VERSION$|^core/VERSION$)"); do

    # Extract the current version and build the expected tag
    dirpath=$(dirname $file)
    version_path="$dirpath/VERSION"
    version=$(<"$version_path")
    expected_tag="$dirpath/$version"
    # Check if the tag already exists
    if git rev-parse --verify $expected_tag &> /dev/null; then
        echo "Tag '$expected_tag' already exists."
    else
        # Tag doesn't exist. Create a tag and push it.
        echo "Tag '$expected_tag' does not exist."
        git tag -a $expected_tag -m "Release $version"
        git push origin tag $expected_tag
    fi
done
