# Project information
OWNER    = "tomohiro"
PACKAGE  = $(shell basename ${PWD})
VERSION  = $(shell git describe --abbrev=0 --tags)

# Build information
DIST_DIR   = $(shell pwd)/dist
ASSETS_DIR = $(DIST_DIR)/$(VERSION)
XC_OS	   = "linux darwin"
XC_ARCH	   = "386 amd64"

# Tasks
help:
	@echo "Please type: make [target]"
	@echo "  setup        Setup development environment"
	@echo "  deps         Install runtime dependencies"
	@echo "  updatedeps   Update runtime dependencies"
	@echo "  lint		  Lint codes"
	@echo "  test         Run tests"
	@echo "  dist         Ship packages as release assets"
	@echo "  release 	  Publish release assets to GitHub"
	@echo "  clean        Clean assets"
	@echo "  help         Show this help messages"

setup:
	@echo "===> Setup development tools..."

	# goxz - Just do cross building and archiving go tools conventionally
	GO111MODULE=off go get -u github.com/Songmu/goxz/cmd/goxz

	# ghr - Upload multiple artifacts to GitHub Release in parallel
	GO111MODULE=off go get -u github.com/tcnksm/ghr

deps:
	@echo "===> Installing runtime dependencies..."
	go mod download

updatedeps:
	@echo "===> Updating runtime dependencies..."
	go get -u

lint: deps
	@echo "===> Running lint..."
	go vet ./...
	golint -set_exit_status ./...

test: deps
	@echo "===> Running tests..."
	go test -v -cover ./...

dist: deps
	@echo "===> Shipping packages as release assets..."
	goxz -d $(ASSETS_DIR) -os $(XC_OS) -arch $(XC_ARCH)

release:
	@echo "===> Publishing release assets to GitHub..."
	ghr -u $(OWNER) -r $(PACKAGE) $(VERSION) $(ASSETS_DIR)

clean:
	@echo "===> Cleaning assets..."
	go clean
	rm -rf $(DIST_DIR)

.PHONY: help setup deps updatedeps test dist release clean
