PACKAGE_NAME := github.com/monetha/reputation-go-sdk
ARTIFACTS_DIR := $(if $(ARTIFACTS_DIR),$(ARTIFACTS_DIR),bin)

PKGS ?= $(shell glide novendor)
PKGS_NO_CMDS ?= $(shell glide novendor | grep -v ./cmd/)
BENCH_FLAGS ?= -benchmem

VERSION := $(if $(TRAVIS_TAG),$(TRAVIS_TAG),$(if $(TRAVIS_BRANCH),$(TRAVIS_BRANCH),development_in_$(shell git rev-parse --abbrev-ref HEAD)))
COMMIT := $(if $(TRAVIS_COMMIT),$(TRAVIS_COMMIT),$(shell git rev-parse HEAD))
BUILD_TIME := $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')

CMD_GO_LDFLAGS := '-X "$(PACKAGE_NAME)/cmd.Version=$(VERSION)" -X "$(PACKAGE_NAME)/cmd.BuildTime=$(BUILD_TIME)" -X "$(PACKAGE_NAME)/cmd.GitHash=$(COMMIT)"'

.PHONY: all
all: lint test

.PHONY: dependencies
dependencies:
	@echo "Installing Glide and locked dependencies..."
	glide --version || go get -u -f github.com/Masterminds/glide
	glide install
	@echo "Installing goimports..."
	go install ./vendor/golang.org/x/tools/cmd/goimports
	@echo "Installing golint..."
	go install ./vendor/golang.org/x/lint/golint
	@echo "Installing gosimple..."
	go install ./vendor/honnef.co/go/tools/cmd/gosimple
	@echo "Installing unused..."
	go install ./vendor/honnef.co/go/tools/cmd/unused
	@echo "Installing staticcheck..."
	go install ./vendor/honnef.co/go/tools/cmd/staticcheck

.PHONY: lint
lint:
	@echo "Checking formatting..."
	@gofiles=$$(go list -f {{.Dir}} $(PKGS) | grep -v mock) && [ -z "$$gofiles" ] || unformatted=$$(for d in $$gofiles; do goimports -l $$d/*.go; done) && [ -z "$$unformatted" ] || (echo >&2 "Go files must be formatted with goimports. Following files has problem:\n$$unformatted" && false)
	@echo "Checking vet..."
	@go vet $(PKG_FILES)
	@echo "Checking simple..."
	@gosimple $(PKG_FILES)
	@echo "Checking unused..."
	@unused $(PKG_FILES)
	@echo "Checking staticcheck..."
	@staticcheck $(PKG_FILES)
	@echo "Checking lint..."
	@$(foreach dir,$(PKGS),golint $(dir);)

.PHONY: test
test:
	go test -tags=dev -timeout 40s -race -v $(PKGS)

.PHONY: bench
BENCH ?= .
bench:
	$(foreach pkg,$(PKGS),go test -bench=$(BENCH) -run="^$$" $(BENCH_FLAGS) $(pkg);)

.PHONY: cover
cover:
	mkdir -p ./.cover
	go test -race -coverprofile=./.cover/cover.out -covermode=atomic -coverpkg=./... $(PKGS_NO_CMDS)
	go tool cover -func=./.cover/cover.out
	go tool cover -html=./.cover/cover.out -o ./.cover/cover.html

.PHONY: fmt
fmt:
	@echo "Formatting files..."
	@gofiles=$$(go list -f {{.Dir}} $(PKGS) | grep -v mock) && [ -z "$$gofiles" ] || for d in $$gofiles; do goimports -l -w $$d/*.go; done

.PHONY: cmd
CMDS ?= $(shell ls -d ./cmd/*/ | xargs -L1 basename | grep -v internal)
cmd: cmd-gen cmd-clean
	$(foreach cmd,$(CMDS),go build --ldflags=$(CMD_GO_LDFLAGS) -o $(ARTIFACTS_DIR)/$(cmd) ./cmd/$(cmd);)

.PHONY: cmdx
CMDX_PLATFORMS = "windows/amd64" "darwin/amd64" "linux/amd64"
CMDX_CMDS = "passport-scanner"
cmdx: cmd-gen cmd-clean
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

.PHONY: cmd-clean
cmd-clean:
	rm -rf $(ARTIFACTS_DIR)