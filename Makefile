ROOT_DIR              ?= $(shell git rev-parse --show-toplevel)
SCRIPTS_BASE          ?= $(ROOT_DIR)/scripts
GOLANG_CI_YAML_PATH   ?= ${ROOT_DIR}/golang-ci.yaml
GOLANG_CI_ARGS        ?= --allow-parallel-runners --timeout=5m --config=${GOLANG_CI_YAML_PATH}

# SETUP AND TOOL INITIALIZATION TASKS
project-help:
	@$(SCRIPTS_BASE)/project.sh help

project-tools:
	@$(SCRIPTS_BASE)/project.sh tools

# LINT
lint-golangci-lint:
	@echo "Linting with golangci-lint"
	@$(SCRIPTS_BASE)/lint-golangci-lint.sh ${skip-non-generated-files}

lint-scripts:
	@echo "Linting scripts"
	@cd ${ROOT_DIR}/scripts && golangci-lint run ${GOLANG_CI_ARGS}

sync-tidy:
	@echo "Syncing and tidying dependencies"
	@$(SCRIPTS_BASE)/sync-tidy.sh

lint: sync-tidy
	@$(MAKE) --no-print-directory lint-golangci-lint skip-non-generated-files=${skip-non-generated-files}

# TEST
test-go:
	@echo "Running Go tests"
	@$(SCRIPTS_BASE)/test-go.sh ${skip-non-generated-files}

test-scripts:
	@echo "Running Go tests for scripts"
	@go test $(ROOT_DIR)/scripts/... ${GOTEST_ARGS}

test:
	@$(MAKE) --no-print-directory test-go skip-non-generated-files=${skip-non-generated-files}

# AUTOMATIC TAG
sdk-tag: 
	@go run $(SCRIPTS_BASE)/automatic_tag.go ${update-type} ${ssh-private-key-file-path} 

sdk-tag-core: 
	@go run $(SCRIPTS_BASE)/automatic_tag.go ${update-type} ${ssh-private-key-file-path} --core-only
