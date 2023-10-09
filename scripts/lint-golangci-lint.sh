#!/bin/bash
# This script lints the SDK modules and examples
# To skip manually maintained files, pass an extra "true" argument
# Pre-requisites: golangci-lint
set -eo pipefail

SKIP_NON_GENERATED_FILES="${1}"
if [ ! "${SKIP_NON_GENERATED_FILES}" = true ]; then
    SKIP_NON_GENERATED_FILES=false
fi

ROOT_DIR=$(git rev-parse --show-toplevel)
CORE_PATH="${ROOT_DIR}/core"
SERVICES_PATH="${ROOT_DIR}/services"
EXAMPLES_PATH="${ROOT_DIR}/examples"
GOLANG_CI_YAML_PATH="${ROOT_DIR}/golang-ci.yaml"
GOLANG_CI_ARGS="--allow-parallel-runners --timeout=5m --config=${GOLANG_CI_YAML_PATH}"

if type -p golangci-lint >/dev/null; then
    :
else
    echo "golangci-lint not installed, unable to proceed."
    exit 1
fi

if [ "${SKIP_NON_GENERATED_FILES}" = false ]; then
    echo ">> Linting core"
    cd ${CORE_PATH}
    golangci-lint run ${GOLANG_CI_ARGS}
fi

for service_dir in ${SERVICES_PATH}/*; do
    service=$(basename ${service_dir})
    echo ">> Linting service ${service}"
    cd ${service_dir}
    if [ "${SKIP_NON_GENERATED_FILES}" = true ]; then
        golangci-lint run ${GOLANG_CI_ARGS} --skip-dirs wait # All manually maintained files are in subfolders
    else
        golangci-lint run ${GOLANG_CI_ARGS}
    fi
done

if [ "${SKIP_NON_GENERATED_FILES}" = false ]; then
    for example_dir in ${EXAMPLES_PATH}/*; do
        example=$(basename ${example_dir})
        echo ">> Linting example ${example}"
        cd ${example_dir}
        golangci-lint run ${GOLANG_CI_ARGS}
    done
fi
