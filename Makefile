# Project Variables

PROJECTNAME := $(shell basename "$(PWD)")
BUILD := $(shell git rev-parse --short HEAD)

# Go variables

GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

# Provide details to build settings
LDFLAGS=-ldflags "-X=main.Build=$(BUILD)"

all: help

## build: Compile binary
build: clean go-build

## exec: Compile binary and run
exec: build go-run

clean:
	@echo " > Cleaning historical builds..."
	@-rm -rf $(GOBIN)/$(PROJECTNAME) 2> /dev/null
	@-$(MAKE) go-clean

help: Makefile
	@echo 
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

go-run:
	@echo " > Running application..."
	@-$(GOBIN)/$(PROJECTNAME)

go-build:
	@echo " > Building project..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-get:
	@echo " > Building dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(get)

go-clean:
	@echo " > Cleaning build cache..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean


