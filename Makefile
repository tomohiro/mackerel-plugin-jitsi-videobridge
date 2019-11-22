# Project information
PACKAGE  = $(shell basename ${PWD})
VERSION  = $(shell git describe --abbrev=0 --tags)

# Build information
DIST_DIR = $(shell pwd)/dist
XC_OS	 = "linux darwin"
XC_ARCH	 = "386 amd64"

# Tasks
help:
	@echo "Please type: make [target]"
	@echo "  setup        Setup development environment"
	@echo "  deps         Install runtime dependencies"
	@echo "  updatedeps   Update runtime dependencies"
	@echo "  test         Run tests"
	@echo "  dist         Ship packages to release"
	@echo "  clean        Clean output binary"
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

test: deps
	@echo "===> Running tests..."
	go test -v -cover ./lib

dist:
	@echo "===> Build and shipping packages..."
	goxz -d $(DIST_DIR) -os $(XC_OS) -arch $(XC_ARCH) -pv $(VERSION) 

clean:
	@echo "===> Clean artifacts..."
	go clean
	rm -rf $(DIST_DIR)

.PHONY: help setup test deps updatedeps clean
