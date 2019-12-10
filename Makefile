# Project information
OWNER    := tomohiro
PACKAGE  := $(shell basename $(PWD))
VERSION  := $(shell git describe --abbrev=0 --tags)

# Build information
DIST_DIR    := $(PWD)/dist
ASSETS_DIR  := $(DIST_DIR)/$(VERSION)
XC_OS       := "linux darwin"
XC_ARCH     := "386 amd64"
BUILD_FLAGS := "-w -s"

# Tasks
.PHONY: help
help:
	@echo "Please type: make [target]"
	@echo "  setup        Setup development environment"
	@echo "  deps         Install runtime dependencies"
	@echo "  updatedeps   Update runtime dependencies"
	@echo "  lint         Lint codes"
	@echo "  test         Run tests"
	@echo "  dist         Ship packages as release assets"
	@echo "  release      Publish release assets to GitHub"
	@echo "  clean        Clean assets"
	@echo "  help         Show this help messages"

.PHONY: setup
setup:
	@echo "===> Setup development tools..."
	# Install goxz - Just do cross building and archiving go tools conventionally
	GO111MODULE=off go get -u github.com/Songmu/goxz/cmd/goxz
	# Install ghr - Upload multiple artifacts to GitHub Release in parallel
	GO111MODULE=off go get -u github.com/tcnksm/ghr

.PHONY: deps
deps:
	@echo "===> Installing runtime dependencies..."
	go mod download

.PHONY: updatedeps
updatedeps:
	@echo "===> Updating runtime dependencies..."
	go get -u

.PHONY: lint
lint: deps
	@echo "===> Running lint..."
	go vet ./...
	golint -set_exit_status ./...

.PHONY: test
test: deps
	@echo "===> Running tests..."
	go test -v -cover ./...

.PHONY: dist
dist: deps
	@echo "===> Shipping packages as release assets..."
	goxz -z -d $(ASSETS_DIR) -os $(XC_OS) -arch $(XC_ARCH) --build-ldflags=$(BUILD_FLAGS)
	pushd $(ASSETS_DIR); \
	shasum -a 256 *.zip > ./$(VERSION)_SHA256SUMS; \
	popd

.PHONY: release
release:
	@echo "===> Publishing release assets to GitHub..."
	ghr -u $(OWNER) -r $(PACKAGE) $(VERSION) $(ASSETS_DIR)

.PHONY: clean
clean:
	@echo "===> Cleaning assets..."
	go clean
	rm -rf $(DIST_DIR)
