BIN := "./bin/app"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -a -o $(BIN) -ldflags "$(LDFLAGS)" main.go

test: 
	go test --short -race ./internal/... ./pkg/...

.PHONY: build test

enterdb:
	docker exec -it ab-postgres psql -U homestead;

generate:
	protoc -I=api --go_out=internal/server/grpc/generated --go-grpc_out=internal/server/grpc/generated api/AppService.proto

up:
	docker-compose up --build

down:
	docker-compose down --volumes
	
integration-tests:
	docker-compose -f docker-compose.tests.yml up --build --abort-on-container-exit --exit-code-from integration-tests
	docker-compose -f docker-compose.tests.yml down --volumes

reset-integration-tests:
	docker-compose -f docker-compose.tests.yml down --volumes	
