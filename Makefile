SHELL := /bin/bash

.PHONY: all
all: deps

.PHONY: deps
deps:
	@echo "===> Fetching dependencies"
	@GO111MODULE=on go mod vendor

.PHONY: clean
clean:
	@echo "===> Cleaning up"
	@rm -rf vendor/
