.PHONY: help report clean

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## Lint
	# go vet ./...

report: lint ## Build dpreport
	mkdir -p dist
	go build -o dist/dpreport cmd/dpreport/dpreport.go

clean: ## Clean intermediate files
	rm -rf dist
