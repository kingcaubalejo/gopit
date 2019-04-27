.PHONY: all
all: build
FORCE: ;

SHELL  := env BOOKMARK_ENV=$(BOOKMARK_ENV) $(SHELL)
BOOKMARK_ENV ?= dev

BIN_DIR = $(PWD)/bin

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=./bin/go-api-jwt-v2
BINARY_UNIX=$(BINARY_NAME)_unix

PROJECTNAME := $(shell basename "$(pwd)")

GO_ENV:=tests
export GO_ENV

PID := /tmp/.$(PROJECTNAME).PID

.PHONY: build

clean:
	rm -rf bin/*

all:
	test build

build:
	# $(GOBUILD) -o $(BINARY_NAME) -v
	dependencies build-api

build-api:
	$(GOBUILD) -tags $(BOOKMARK_ENV) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./..

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f ./bin/$(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME)

deps:
	$(GOGET) -u github.com/dgrijalva/jwt-go
	$(GOGET) -u github.com/auth0/go-jwt-middleware
	$(GOGET) -u github.com/go-sql-driver/mysql
	$(GOGET) -u github.com/gorilla/handlers
	$(GOGET) -u github.com/urfave/negroni
	$(GOGET) -u github.com/onsi/ginkgo/ginkgo
	$(GOGET) -u github.com/onsi/gomega