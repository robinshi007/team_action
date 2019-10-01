GOCMD=go
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
SERVER_BINARY_NAME=team_action
SERVER_BINARY_UNIX=$(SERVER_BINARY_NAME)_unix
APP_BINARY_NAME=team_action_app
APP_BINARY_UNIX=$(APP_BINARY_NAME)_unix
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
	$(GOBUILD) -o $(SERVER_BINARY_NAME) -v $(SERVER_ENTRY)
build_app:
	$(GOBUILD) -o $(APP_BINARY_NAME) -v $(APP_ENTRY)
build_asset:
	statik -src=./dist
clean:
	$(GOCLEAN) $(SERVER_ENTRY)
	$(GOCLEAN) $(APP_ENTRY)
	rm -f $(SERVER_BINARY_NAME)
	rm -f $(SERVER_BINARY_UNIX)
	rm -f $(APP_BINARY_NAME)
	rm -f $(APP_BINARY_UNIX)
	rm -rf statik

