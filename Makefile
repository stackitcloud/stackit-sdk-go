ROOT_DIR              ?= $(shell git rev-parse --show-toplevel)
SCRIPTS_BASE          ?= $(ROOT_DIR)/scripts

# SETUP AND TOOL INITIALIZATION TASKS
project-help:
	@$(SCRIPTS_BASE)/project.sh help

project-tools:
	@$(SCRIPTS_BASE)/project.sh tools

# LINT
lint-golangci-lint:
	@echo "Linting with golangci-lint"
	@$(SCRIPTS_BASE)/lint-golangci-lint.sh ${skip-non-generated-files}

sync-tidy:
	@echo "Syncing and tidying dependencies"
	@$(SCRIPTS_BASE)/sync-tidy.sh

lint: sync-tidy
	@$(MAKE) --no-print-directory lint-golangci-lint skip-non-generated-files=${skip-non-generated-files}

# TEST
test-go:
	@echo "Running Go tests"
	@$(SCRIPTS_BASE)/test-go.sh ${skip-non-generated-files}

test:
	@$(MAKE) --no-print-directory test-go skip-non-generated-files=${skip-non-generated-files}
