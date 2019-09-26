GOCMD=go
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=team_action
BINARY_UNIX=$(BINARY_NAME)_unix
SERVER_ENTRY=cmd/server/main.go
APP_ENTRY=cmd/app/main.go

.PHONY: all
all: run test build clean

run:
	$(GORUN) $(SERVER_ENTRY)
run_app:
	$(GORUN) $(APP_ENTRY)
test:
	$(GOTEST) -v ./...
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(SERVER_ENTRY)
clean:
	$(GOCLEAN) cmd/server/main.go
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

