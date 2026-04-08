BINARY    := bin/server
GOBIN     := $(shell go env GOPATH)/bin
SWAG      := $(GOBIN)/swag
LINT      := $(GOBIN)/golangci-lint
GOIMPORTS := $(GOBIN)/goimports

.PHONY: dev backend-run frontend-dev frontend-build build \
        lint lint-backend lint-frontend \
        fmt fmt-backend fmt-frontend \
        test test-backend test-frontend \
        swagger

dev:
	@$(MAKE) -j2 backend-run frontend-dev

backend-run:
	go run ./cmd/server

frontend-dev:
	cd frontend && npm run dev

frontend-build:
	cd frontend && npm run build

swagger:
	$(SWAG) init -g cmd/server/main.go -o docs

build: frontend-build swagger
	mkdir -p bin
	go build -o $(BINARY) ./cmd/server

lint: lint-backend lint-frontend

lint-backend:
	$(LINT) run ./...

lint-frontend:
	cd frontend && npm run lint

fmt: fmt-backend fmt-frontend

fmt-backend:
	gofmt -w ./cmd ./internal
	$(GOIMPORTS) -w ./cmd ./internal

fmt-frontend:
	cd frontend && npm run format

test: test-backend test-frontend

test-backend:
	go test $(shell go list ./... | grep -v frontend) -race -count=1

test-frontend:
	cd frontend && npm test
