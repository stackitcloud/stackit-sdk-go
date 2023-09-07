#!/bin/bash
# This script tests the SDK modules
# Pre-requisites: Go
set -eo pipefail

ROOT_DIR=$(git rev-parse --show-toplevel)
GOTEST_ARGS="-timeout=5m -cover -count=1"
CORE_PATH="${ROOT_DIR}/core"
SERVICES_PATH="${ROOT_DIR}/services"

if type -p go >/dev/null; then
    :
else
    echo "Go not installed, unable to proceed."
    exit 1
fi

echo ">> Testing core"
cd ${CORE_PATH}
go test ./... ${GOTEST_ARGS}

for service_dir in ${SERVICES_PATH}/*; do
    service=$(basename ${service_dir})
    echo ">> Testing services/${service}"
    cd ${service_dir}
    go test ./... ${GOTEST_ARGS}
done
