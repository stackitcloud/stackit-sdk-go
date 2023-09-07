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
	@$(SCRIPTS_BASE)/lint-golangci-lint.sh

sync-tidy:
	@echo "Syncing and tidying dependencies"
	@$(SCRIPTS_BASE)/sync-tidy.sh

lint: sync-tidy lint-golangci-lint

# TEST
test-go:
	@echo "Running Go tests"
	@$(SCRIPTS_BASE)/test-go.sh

test: test-go
