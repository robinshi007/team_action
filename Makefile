GOCMD=go
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=team_action
BINARY_UNIX=$(BINARY_NAME)_unix
PACKAGE_ENTRY=cmd/server/main.go

.PHONY: all
all: run test build clean

run:
	$(GORUN) $(PACKAGE_ENTRY)
test:
	$(GOTEST) -v ./...
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(PACKAGE_ENTRY)
clean:
	$(GOCLEAN) cmd/server/main.go
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

