.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## Lint
	go vet ./...

dpjournal: ## Build dpjournal
	mkdir -p dist
	go build -o dist cmd/dpjournal/dpjournal.go
