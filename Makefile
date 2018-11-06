PKGS ?= $(shell glide novendor)
PKGS_NO_CMDS ?= $(shell glide novendor | grep -v ./cmd/)
BENCH_FLAGS ?= -benchmem

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
	go test -timeout 40s -race -v $(PKGS)

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
CMDS ?= $(shell ls cmd | grep -v internal)
cmd:
	go generate ./cmd/...
	$(foreach cmd,$(CMDS),go build -o ./bin/$(cmd) ./cmd/$(cmd);)