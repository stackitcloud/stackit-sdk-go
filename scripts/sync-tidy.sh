#!/bin/bash
# This script runs go work sync and go mod tidy for SDK modules and examples
set -eo pipefail

ROOT_DIR=$(git rev-parse --show-toplevel)
CORE_PATH="${ROOT_DIR}/core"
SERVICES_PATH="${ROOT_DIR}/services"
EXAMPLES_PATH="${ROOT_DIR}/examples"

cd ${ROOT_DIR}
go work sync
# go.work.sum gets a lot of content that is useless after versions are synced
# So we delete it and sync again, to keep what's actually used
if [ -f "${ROOT_DIR}/go.work.sum" ]; then
    rm ${ROOT_DIR}/go.work.sum
fi
go work sync

cd ${CORE_PATH}
go mod tidy

for service_dir in ${SERVICES_PATH}/*; do
    cd ${service_dir}
    go mod tidy
done

for example_dir in ${EXAMPLES_PATH}/*; do
    cd ${example_dir}
    go mod tidy
done

# Fail when changes are produced
git add -A
if git diff --cached --exit-code --quiet; then # --cached flag needed so new files are considered
    echo "All go.mod files are is synced and tidy."
else
    echo "Please run "make sync-tidy" locally to sync dependencies and tidy go.mod files."
    exit 1
fi
