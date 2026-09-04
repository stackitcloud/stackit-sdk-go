#!/usr/bin/env bash
# This script runs go work sync and go mod tidy for SDK modules and examples
set -eo pipefail

ROOT_DIR=$(git rev-parse --show-toplevel)
CORE_PATH="${ROOT_DIR}/core"
SERVICES_PATH="${ROOT_DIR}/services"
EXAMPLES_PATH="${ROOT_DIR}/examples"

echo ">> Syncing go workspace"
cd ${ROOT_DIR}
go work sync

echo ">> Syncing core module"
cd ${CORE_PATH}
go mod tidy

for service_dir in ${SERVICES_PATH}/*; do
    echo ">> Syncing ${service_dir} module"
    cd ${service_dir}
    go mod tidy
done

for example_dir in ${EXAMPLES_PATH}/*; do
    echo ">> Syncing ${example_dir} module"
    cd ${example_dir}
    go mod tidy
done

# go.work.sum gets a lot of content that is useless after versions are synced
# So we delete it and sync again, to keep what's actually used
if [ -f "${ROOT_DIR}/go.work.sum" ]; then
    rm ${ROOT_DIR}/go.work.sum
fi
cd ${ROOT_DIR}
echo ">> Re-syncing go workspace"
go work sync
