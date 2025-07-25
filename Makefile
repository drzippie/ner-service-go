# NER Service Go Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# CGO flags for MITIE
CGO_CFLAGS=-I/opt/homebrew/include
CGO_LDFLAGS=-L/opt/homebrew/lib

# Binary names
SERVER_BINARY=ner-server
CLI_BINARY=ner-cli

# Directories
SERVER_DIR=cmd/server
CLI_DIR=cmd/cli

.PHONY: all build clean test deps server cli

all: deps build

deps:
	$(GOMOD) download
	$(GOMOD) tidy

build: server cli

server:
	CGO_CFLAGS="$(CGO_CFLAGS)" CGO_LDFLAGS="$(CGO_LDFLAGS)" $(GOBUILD) -o $(SERVER_BINARY) $(SERVER_DIR)/main.go

cli:
	CGO_CFLAGS="$(CGO_CFLAGS)" CGO_LDFLAGS="$(CGO_LDFLAGS)" $(GOBUILD) -o $(CLI_BINARY) $(CLI_DIR)/main.go

clean:
	$(GOCLEAN)
	rm -f $(SERVER_BINARY)
	rm -f $(CLI_BINARY)

test:
	$(GOTEST) -v ./...

run-server:
	CGO_CFLAGS="$(CGO_CFLAGS)" CGO_LDFLAGS="$(CGO_LDFLAGS)" $(GOCMD) run $(SERVER_DIR)/main.go

run-cli:
	CGO_CFLAGS="$(CGO_CFLAGS)" CGO_LDFLAGS="$(CGO_LDFLAGS)" $(GOCMD) run $(CLI_DIR)/main.go

install-mitie:
	@echo "Installing MITIE..."
	@command -v brew >/dev/null 2>&1 && brew install mitie || echo "Please install MITIE manually"

download-model:
	@echo "Downloading Spanish MITIE model..."
	@mkdir -p models
	@wget https://sourceforge.net/projects/mitie.mirror/files/v0.4/MITIE-models-v0.2-Spanish.zip/download -O models/spanish_model.zip
	@unzip models/spanish_model.zip -d models/
	@mv models/MITIE-models/spanish/* models/ 2>/dev/null || true
	@rm -rf models/MITIE-models models/spanish_model.zip
	@echo "Spanish model downloaded to models/ner_model.dat"

setup: install-mitie download-model deps build

.DEFAULT_GOAL := all