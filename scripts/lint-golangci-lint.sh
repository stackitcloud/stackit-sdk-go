#!/bin/bash
# This script lints the SDK modules and examples
# Pre-requisites: golangci-lint
set -eo pipefail

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

echo ">> Linting core"
cd ${CORE_PATH}
golangci-lint run ${GOLANG_CI_ARGS}

for service_dir in ${SERVICES_PATH}/*; do
    service=$(basename ${service_dir})
    echo ">> Linting service ${service}"
    cd ${service_dir}
    golangci-lint run ${GOLANG_CI_ARGS}
done

for example_dir in ${EXAMPLES_PATH}/*; do
    example=$(basename ${example_dir})
    echo ">> Linting example ${example}"
    cd ${example_dir}
    golangci-lint run ${GOLANG_CI_ARGS}
done
