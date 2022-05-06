LOCAL_BIN=$(CURDIR)/bin

.PHONY: build start test

.EXPORT_ALL_VARIABLES:
GO111MODULE = on
APP_STAGE = local

include bin-deps.mk

lint: $(GOLANGCI_BIN)
	$(GOLANGCI_BIN) run ./...

build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/server cmd/server/main.go

# db should be already up before starting server
run: build
	./bin/server --port 8080

test: gotest gofmt govet

gotest:
	go test ./... -v -cover

gofmt:
	go fmt ./...

govet:
	go vet ./...

gogenerate:
	go generate ./... -v

gen-server:
	swagger generate server -A flights -f ./api/spec/server.yaml --server-package=./internal/server/restapi --model-package=./pkg/models/rest --exclude-main --principal rest.Principal --with-context
