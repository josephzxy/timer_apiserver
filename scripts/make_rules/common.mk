SHELL := /bin/bash

SUB_MAKEFILES_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
PROJECT_ROOT := $(abspath $(shell cd $(SUB_MAKEFILES_DIR)/../../ && pwd -P))
OUTPUT_DIR := $(PROJECT_ROOT)/_output

APP_NAME := timer_apiserver

GIT_COMMIT := $(shell cd $(PROJECT_ROOT) && git log -1 --pretty=format:'%h')

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

DKR_ARCH ?= $(GOARCH)

SPACE :=
SPACE +=
