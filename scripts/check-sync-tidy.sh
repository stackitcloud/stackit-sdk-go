#!/bin/bash
# This script is run in CI to check is sync-tidy produced any changes and it fails if it did
set -eo pipefail

git add -A
if git diff --cached --exit-code --quiet; then # --cached flag needed so new files are considered
    echo "All go.mod files are is synced and tidy."
else
    echo "Please run "make sync-tidy" locally to sync dependencies and tidy go.mod files. Commit the changes."
    exit 1
fi
