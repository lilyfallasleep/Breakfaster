# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_PATH_SERVER=./cmd/server
BINARY_NAME_SERVER=$(BINARY_PATH_SERVER)/breakfaster
BINARY_UNIX_SERVER=$(BINARY_NAME_SERVER)_unix

BINARY_PATH_MESSAGING=./cmd/messaging
BINARY_NAME_MESSAGING=$(BINARY_PATH_MESSAGING)/messaging
BINARY_UNIX_MESSAGING=$(BINARY_NAME_MESSAGING)_unix

all: deps doc messaging build
build:
	$(GOBUILD) -o $(BINARY_NAME_SERVER) -v $(BINARY_PATH_SERVER)/main.go
messaging:
	$(GOBUILD) -o $(BINARY_NAME_MESSAGING) -v $(BINARY_PATH_MESSAGING)/main.go
test:
	$(GOTEST) -v ./...
doc:
	$(shell $(GOCMD) env GOPATH)/bin/swag init -g $(BINARY_PATH_SERVER)/main.go
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME_SERVER)
	rm -f $(BINARY_UNIX_SERVER)
	rm -f $(BINARY_NAME_MESSAGING)
	rm -f $(BINARY_UNIX_MESSAGING)
deps:
	$(GOGET) -u github.com/swaggo/swag/cmd/swag
	

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX_SERVER) -v
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX_MESSAGING) -v