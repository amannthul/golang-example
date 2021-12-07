SOURCE_FILES?=./...

export PATH := ./bin:$(PATH)
export GO111MODULE := on
export GOPATH := $(shell go env GOPATH)
export GOPROXY := https://goproxy.io,direct
export GOPRIVATE := github.com/amannthul/*

# Install all the build and lint dependencies
setup:
	@./buf-tools-setup.sh
.PHONY: setup

update-tools:
	@./buf-tools-update.sh
.PHONY: update-tools

# Format go files
format:
	@goimports -w ./
.PHONY: format

# Run all the linters
lint:
	@./buf-check-lint.sh
.PHONY: lint

# Check proto break changes
break-check:
	@./buf-check-breaking.sh
.PHONY: break-check

# Check with lint and break-check
check: lint break-check
.PHONY: check

# Clean all generated files.
clean:
	@./buf-clean.sh
.PHONY: clean

# Generate proto documents
gen-doc:
	@./buf-gen-doc.sh
.PHONY: gen-doc

# Generate codes
gen-codes:
	@./buf-gen-codes.sh
.PHONY: gen-codes

# Generate proto descriptor set
build-images:
	@./buf-build-image.sh
.PHONY: build-images

# Build a beta version of goreleaser
build:
	@./build.sh
.PHONY: build

.DEFAULT_GOAL := build
