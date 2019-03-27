GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=go-api-jwt-v2
BINARY_UNIX=$(BINARY_NAME)_unix

PROJECTNAME := $(shell basename "$(pwd)")

GO_ENV:=tests
export GO_ENV

PID := /tmp/.$(PROJECTNAME).PID

all:
	test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v 
test:
	$(GOTEST) -v ./..
clean:
	$(GOCLEAN)
	rm -f ./bin/$(BINARY_NAME)
	rm -f ./bin/$(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"
	@echo "-------------------------------------"
	@echo "GO_ENV=""$$GO_ENV"
	./$(BINARY_NAME)

deps:
	$(GOGET) -u github.com/dgrijalva/jwt-go
	$(GOGET) -u github.com/auth0/go-jwt-middleware
	$(GOGET) -u github.com/go-sql-driver/mysql
	$(GOGET) -u github.com/gorilla/handlers
	$(GOGET) -u github.com/urfave/negroni
