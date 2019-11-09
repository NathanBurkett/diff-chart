.PHONY: help

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test-race: ## Run Tests with Race Detector
	@go test -race -count=1 ./...

test: ## Run Tests
	@go test -count=1 ./...

test-full: test test-race

coverage-html: ## Build coverage report with html output
	@./coverage/run.sh --html

coverage: ## Build coverage report
	@./coverage/run.sh

lint:
	@$(GOPATH)/bin/golint -set_exit_status
