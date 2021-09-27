SHELL := /bin/bash

SUB_MAKEFILES_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
PROJECT_ROOT := $(abspath $(shell cd $(SUB_MAKEFILES_DIR)/../../ && pwd -P))
OUTPUT_DIR := $(PROJECT_ROOT)/_output
