export GO_VERSION=1.21.1
export GOLANGCI_LINT_VERSION=v1.54.0

# configuration/aliases
devimage=metrica-dev
# To avoid downloading deps everytime it runs on containers
gopkg=$(devimage)-gopkg
gocache=$(devimage)-gocache
devrun=docker run $(devrunopts) --rm \
	-v `pwd`:/app \
	-v $(gopkg):/go/pkg \
	-v $(gocache):/root/.cache/go-build \
	$(devimage)


## run isolated tests
.PHONY: test
test:
	go test ./... -timeout 10s -race -shuffle on

## Run lint
.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./... -v
	go run golang.org/x/vuln/cmd/govulncheck ./...

## Create the dev container image
.PHONY: dev/image
dev/image:
	docker build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION) \
		-t $(devimage) \
		-f Dockerfile.dev \
		.

## Create a shell inside the dev container
.PHONY: dev
dev: devrunopts=-ti
dev: dev/image
	$(devrun)

## run a make target inside the dev container.
dev/%: dev/image
	$(devrun) make ${*}

## Display help for all targets
.PHONY: help
help:
	@awk '/^.PHONY: / { \
		msg = match(lastLine, /^## /); \
			if (msg) { \
				cmd = substr($$0, 9, 100); \
				msg = substr(lastLine, 4, 1000); \
				printf "  ${GREEN}%-30s${RESET} %s\n", cmd, msg; \
			} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
