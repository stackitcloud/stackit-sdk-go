ROOT_DIR              ?= $(shell git rev-parse --show-toplevel)
SCRIPTS_BASE          ?= $(ROOT_DIR)/scripts
GOLANG_CI_YAML_PATH   ?= ${ROOT_DIR}/golang-ci.yaml
GOLANG_CI_ARGS        ?= --allow-parallel-runners --timeout=5m --config=${GOLANG_CI_YAML_PATH}

##
# Console Colors
##
GREEN  := $(shell printf "\033[0;32m")
YELLOW := $(shell printf "\033[0;33m")
WHITE  := $(shell printf "\033[0;37m")
CYAN   := $(shell printf "\033[0;36m")
RESET  := $(shell printf "\033[0m")


##
# Targets
##
.PHONY: help
help: ## Show this help
	@echo "Project: stackit-sdk-go"
	@echo 'Usage:'
	@echo "  ${GREEN}make${RESET} ${YELLOW}<target>${RESET}"
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "  ${GREEN}%-21s${YELLOW}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST) | sort

# SETUP AND TOOL INITIALIZATION TASKS
project-help: ## Show help for the project
	@$(SCRIPTS_BASE)/project.sh help

project-tools: ## Install project tools
	@$(SCRIPTS_BASE)/project.sh tools

# LINT
lint-golangci-lint: ## Lint Go code
	@echo ">> Linting with golangci-lint"
	@$(SCRIPTS_BASE)/lint-golangci-lint.sh "${skip-non-generated-files}" "${service}"

lint-scripts: ## Lint scripts
	@echo ">> Linting scripts"
	@cd ${ROOT_DIR}/scripts && golangci-lint run ${GOLANG_CI_ARGS}

sync-tidy: ## Sync and tidy dependencies
	@echo ">> Syncing and tidying dependencies"
	@$(SCRIPTS_BASE)/sync-tidy.sh

lint: sync-tidy ## Lint all code
	@$(MAKE) --no-print-directory lint-golangci-lint skip-non-generated-files=${skip-non-generated-files} service=${service}

# TEST
test-go: ## Run Go tests
	@echo ">> Running Go tests"
	@$(SCRIPTS_BASE)/test-go.sh "${skip-non-generated-files}" "${service}"

test-scripts: ## Run tests for scripts
	@echo ">> Running Go tests for scripts"
	@go test $(ROOT_DIR)/scripts/... ${GOTEST_ARGS}

test: ## Run all tests
	@$(MAKE) --no-print-directory test-go skip-non-generated-files=${skip-non-generated-files} service=${service}

# AUTOMATIC TAG
sdk-tag-services:
	@go run $(SCRIPTS_BASE)/automatic_tag.go --update-type ${update-type} --ssh-private-key-file-path ${ssh-private-key-file-path};


sdk-tag-core: 
	@go run $(SCRIPTS_BASE)/automatic_tag.go --update-type ${update-type} --ssh-private-key-file-path ${ssh-private-key-file-path} --target core;
