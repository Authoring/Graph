# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=graph
BINARY_UNIX=$(BINARY_NAME)_unix

all: build run

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	./$(BINARY_NAME)
