.PHONY: help dpjournal clean

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## Lint
	# go vet ./...

dpreport: lint ## Build dpreport
	mkdir -p dist
	go build -o dist cmd/dpreport/main.go

clean: ## Clean intermediate files
	rm -rf dist
