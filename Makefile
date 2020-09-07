SHELL := /bin/bash

#!make

#include .env.test
#export $(shell sed 's/=.*//' .env.test)

.EXPORT_ALL_VARIABLES:
SRC_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
OUT_DIR := $(SRC_DIR)/_output
BIN_DIR := $(OUT_DIR)/bin
PLUGIN_DIR := $(BIN_DIR)/plugin
FTEST_DIR := test/procs
CONFIG_DIR := test/config
GOPROXY ?= https://proxy.golang.org
GO111MODULE := on
CONFIG_LOCATION := $(SRC_DIR)

$(@info $(shell mkdir -p $(OUT_DIR) $(BIN_DIR) $(PLUGIN_DIR))

.PHONY: build
build: proto_files cli server executor

.PHONY: proto_files
proto_files:
	bash ./proto_gen.sh

.PHONY: executor
executor:
	go build -race -o $(BIN_DIR)/executor ./cmd/executor/main.go

.PHONY: server
server:
	go build -race -o $(BIN_DIR)/controller ./cmd/controller/main.go

.PHONY: start-server
start-server:
	PROCTOR_AUTH_PLUGIN_BINARY=$(PLUGIN_DIR)/auth.so \
	PROCTOR_NOTIFICATION_PLUGIN_BINARY=$(PLUGIN_DIR)/slack.so \
	$(BIN_DIR)/server s

.PHONY: cli
cli:
	go build -race -o $(BIN_DIR)/cli ./cmd/cli/main.go

