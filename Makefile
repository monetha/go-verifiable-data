SHELL := bash
PACKAGE_NAME := github.com/monetha/reputation-go-sdk
ARTIFACTS_DIR := $(if $(ARTIFACTS_DIR),$(ARTIFACTS_DIR),bin)

PKGS ?= $(shell glide novendor | grep -v ./internal/)
PKGS_NO_CMDS ?= $(shell glide novendor | grep -Ev './internal/|./cmd/')
BENCH_FLAGS ?= -benchmem

VERSION := $(if $(TRAVIS_TAG),$(TRAVIS_TAG),$(if $(TRAVIS_BRANCH),$(TRAVIS_BRANCH),development_in_$(shell git rev-parse --abbrev-ref HEAD)))
COMMIT := $(if $(TRAVIS_COMMIT),$(TRAVIS_COMMIT),$(shell git rev-parse HEAD))
BUILD_TIME := $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')

CMD_GO_LDFLAGS := '-X "$(PACKAGE_NAME)/cmd.Version=$(VERSION)" -X "$(PACKAGE_NAME)/cmd.BuildTime=$(BUILD_TIME)" -X "$(PACKAGE_NAME)/cmd.GitHash=$(COMMIT)"'

export GO111MODULE := on

.PHONY: all
all: lint test

.PHONY: dependencies
dependencies:
	@echo "Installing dependencies..."
	go mod download
	@echo "Installing goimports..."
	go install golang.org/x/tools/cmd/goimports
	@echo "Installing golint..."
	go install golang.org/x/lint/golint
	@echo "Installing staticcheck..."
	go install honnef.co/go/tools/cmd/staticcheck
	@echo "Installing enumer..."
	go install github.com/alvaroloes/enumer
	@echo "Installing abigen..."
	go install github.com/ethereum/go-ethereum/cmd/abigen

.PHONY: lint
lint:
	@echo "Checking formatting..."
	@gofiles=$$(go list -f {{.Dir}} $(PKGS) | grep -v mock) && [ -z "$$gofiles" ] || unformatted=$$(for d in $$gofiles; do goimports -l $$d/*.go; done) && [ -z "$$unformatted" ] || (echo >&2 "Go files must be formatted with goimports. Following files has problem:\n$$unformatted" && false)
	@echo "Checking vet..."
	@go vet $(PKG_FILES)
	@echo "Checking staticcheck..."
	@staticcheck $(PKG_FILES)
	@echo "Checking lint..."
	@$(foreach dir,$(PKGS),golint $(dir);)

.PHONY: test
test:
	go test -count=1 -tags=dev -timeout 60s -race -v $(PKGS)

.PHONY: bench
BENCH ?= .
bench:
	$(foreach pkg,$(PKGS),go test -bench=$(BENCH) -run="^$$" $(BENCH_FLAGS) $(pkg);)

.PHONY: cover
cover:
	mkdir -p ./${ARTIFACTS_DIR}/.cover
	go test -count=1 -race -coverprofile=./${ARTIFACTS_DIR}/.cover/cover.out -covermode=atomic -coverpkg=./... $(PKGS_NO_CMDS)
	go tool cover -func=./${ARTIFACTS_DIR}/.cover/cover.out
	go tool cover -html=./${ARTIFACTS_DIR}/.cover/cover.out -o ./${ARTIFACTS_DIR}/cover.html

.PHONY: fmt
fmt:
	@echo "Formatting files..."
	@gofiles=$$(go list -f {{.Dir}} $(PKGS) | grep -v mock) && [ -z "$$gofiles" ] || for d in $$gofiles; do goimports -l -w $$d/*.go; done

.PHONY: cmd
CMDS ?= $(shell ls -d ./cmd/*/ | xargs -L1 basename | grep -v internal)
cmd: cmd-gen
	$(foreach cmd,$(CMDS),go build --ldflags=$(CMD_GO_LDFLAGS) -o $(ARTIFACTS_DIR)/$(cmd) ./cmd/$(cmd);)

.PHONY: cmdx
CMDX_PLATFORMS = "windows/amd64" "darwin/amd64" "linux/amd64"
CMDX_CMDS = "passport-scanner"
cmdx: cmd-gen
	for platform in $(CMDX_PLATFORMS); do \
		platform_split=($${platform//\// }); \
		GOOS=$${platform_split[0]}; \
		GOARCH=$${platform_split[1]}; \
		HUMAN_OS=$${GOOS}; \
		if [ "$$HUMAN_OS" = "darwin" ]; then \
			HUMAN_OS='macos'; \
		fi; \
		for cmd in $(CMDX_CMDS); do \
			output_name=$(ARTIFACTS_DIR)/$${cmd}; \
			if [ "$$GOOS" = "windows" ]; then \
				output_name+='.exe'; \
			fi; \
			env GOOS=$$GOOS GOARCH=$$GOARCH go build --ldflags=$(CMD_GO_LDFLAGS) -o $${output_name} ./cmd/$${cmd}; \
			if [ "$$GOOS" = "windows" ]; then \
				pushd ${ARTIFACTS_DIR}; zip $${cmd}-$${HUMAN_OS}-$${GOARCH}-$(VERSION).zip $${cmd}.exe; popd; \
			else \
				pushd ${ARTIFACTS_DIR}; tar cvzf $${cmd}-$${HUMAN_OS}-$${GOARCH}-$(VERSION).tgz $${cmd}; popd; \
			fi; \
			rm $${output_name}; \
		done; \
	done

.PHONY: cmd-gen
cmd-gen:
	go generate ./cmd/...
