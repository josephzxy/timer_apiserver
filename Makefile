# Build all by default
.DEFAULT_GOAL := all

.PHONY: all
all: tidy format

# Includes
include scripts/make_rules/common.mk # must always be the first
include scripts/make_rules/golang.mk

.PHONY: tidy
tidy: go.tidy

.PHONY: format
format: go.format

.PHONY: clean
clean: go.clean
