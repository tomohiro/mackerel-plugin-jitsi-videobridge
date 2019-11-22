# Project information
PACKAGE=mackerel-plugin-jitsi-videobridge

# Tasks
help:
	@echo "Please type: make [target]"
	@echo "  test         Run tests"
	@echo "  deps         Install runtime dependencies"
	@echo "  updatedeps   Update runtime dependencies"
	@echo "  help         Show this help messages"

test: deps
	@echo "===> Running tests..."
	go test -v -cover ./lib

deps:
	@echo "===> Installing runtime dependencies..."
	go mod download

updatedeps:
	@echo "===> Updating runtime dependencies..."
	go get -u

.PHONY: help test deps updatedeps
