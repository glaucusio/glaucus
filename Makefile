all: help

.PHONY: cmd
cmd: DATE := $(shell date '+%FT%T.%N%:z')
cmd: GIT_COMMIT := $(shell git rev-list -1 HEAD)
cmd: CGO_ENABLED := 0
cmd: ## Build all commands
	@echo "# Building bin/gctl ..."
	@go build -mod=vendor -installsuffix cgo -ldflags '-X "command.buildDate=$(DATE)" -X "main.gitCommit=$(GIT_COMMIT)"' -o ./bin/gctl ./cmd/gctl
	@echo "# Building bin/glaucus ..."
	@go build -mod=vendor -installsuffix cgo -ldflags '-s -w -X "command.buildDate=$(DATE)" -X "main.gitCommit=$(GIT_COMMIT)"' -o ./bin/glaucus ./cmd/glaucus

.PHONY: init
init: ## Installs tools

.PHONY: help
help: ## This help command
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {printf "\033[36m%-25s\033[0m %s\n", $$1, $$NF}' $(MAKEFILE_LIST)
