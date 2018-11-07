#
# Makefile for building all things related to this repo
#
NAME := go-schema
ORG := pinpt
PKG := $(ORG)/$(NAME)
SHELL := /bin/bash

.PHONY: all test dependencies build generate bindata schema clean

all: build test

dependencies:
	@dep ensure
	@go get golang.org/x/tools/cmd/goimports
	@go get -u github.com/jhaynie/protoc-gen-gator

build: dependencies
	@go run main.go all

generate:
	@go run main.go generate

bindata:
	@go run main.go bindata

schema:
	@go run main.go schema

clean:
	@go run main.go clean

test:
	@go test -v ./... | grep -v "?"
