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
doc:
	$(shell $(GOCMD) env GOPATH)/bin/swag init -g $(BINARY_PATH_SERVER)/main.go

test: pretest runtest
runtest:
	$(GOTEST) -v ./...
pretest: mock
	$(shell $(GOCMD) env GOPATH)/bin/mockgen -source=repository/dao/interface.go -destination=mocks/mock_dao/mock_dao.go
	$(shell $(GOCMD) env GOPATH)/bin/mockgen -source=service/core/interface.go -destination=mocks/mock_service/mock_service.go
	$(shell $(GOCMD) env GOPATH)/bin/mockgen -source=pkg/ordertime/interface.go -destination=mocks/mock_ordertime/mock_ordertime.go
	$(shell $(GOCMD) env GOPATH)/bin/mockgen -source=infrastructure/cache/redis_interface.go -destination=mocks/mock_cache/mock_cache_redis.go
	$(shell $(GOCMD) env GOPATH)/bin/mockgen -source=messaging/interface.go -destination=mocks/mock_messaging/mock_messaging.go
	$(shell $(GOCMD) env GOPATH)/bin/mockgen -source=mybot/interface.go -destination=mocks/mock_mybot/mock_mybot.go
	$(shell $(GOCMD) env GOPATH)/bin/mockgen -source=mybot/autoreply/interface.go -destination=mocks/mock_autoreply/mock_autoreply.go

swag:
	$(GOGET) github.com/swaggo/swag/cmd/swag@v1.6.7
mock:
	GO111MODULE=on $(GOGET) github.com/golang/mock/mockgen@v1.4.4

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME_SERVER)
	rm -f $(BINARY_UNIX_SERVER)
	rm -f $(BINARY_NAME_MESSAGING)
	rm -f $(BINARY_UNIX_MESSAGING)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX_SERVER) -v
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX_MESSAGING) -v