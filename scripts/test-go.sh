#!/bin/bash
# This script tests the SDK modules
# To skip manually maintained files, pass an extra "true" argument
# Pre-requisites: Go
set -eo pipefail

SKIP_NON_GENERATED_FILES="${1}"
if [ ! "${SKIP_NON_GENERATED_FILES}" = true ]; then
    SKIP_NON_GENERATED_FILES=false
fi

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

if [ "${SKIP_NON_GENERATED_FILES}" = false ]; then
    echo ">> Testing core"
    cd ${CORE_PATH}
    go test ./... ${GOTEST_ARGS}
fi

for service_dir in ${SERVICES_PATH}/*; do
    service=$(basename ${service_dir})

    # Our unit test template fails because it doesn't support fields with validations,
    # such as the UUID component used by IaaS. We introduce this hardcoded skip until we fix it
    if [ "${service}" = "iaas" ]; then
        echo ">> Skipping services/${service}"
    fi

    echo ">> Testing services/${service}"
    cd ${service_dir}
    if [ "${SKIP_NON_GENERATED_FILES}" = true ]; then
        go test ./ ${GOTEST_ARGS} # All manually maintained files are in subfolders
    else
        go test ./... ${GOTEST_ARGS}
    fi
done
