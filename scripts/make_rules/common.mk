SHELL := /bin/bash

SUB_MAKEFILES_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
PROJECT_ROOT := $(abspath $(shell cd $(SUB_MAKEFILES_DIR)/../../ && pwd -P))
OUTPUT_DIR := $(PROJECT_ROOT)/_output

ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, , $(PLATFORM)))
	GOARCH := $(word 2, $(subst _, , $(PLATFORM)))
endif

SPACE :=
SPACE +=
