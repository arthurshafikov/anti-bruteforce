BIN := "./bin/app"
DOCKER_COMPOSE_FILE := "./deployments/docker-compose.yml"
DOCKER_COMPOSE_TEST_FILE := "./deployments/docker-compose.tests.yml"
APP_TEST_NAME="anti_bruteforce_test"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -a -o $(BIN) -ldflags "$(LDFLAGS)" main.go

run: build 
	${BIN} serve

test: 
	go test --short -race ./internal/... ./pkg/...

lint: 
	golangci-lint run ./...

.PHONY: build test

enterdb:
	docker exec -it ab-postgres psql -U homestead;

generate:
	protoc -I=api --go_out=internal/server/grpc/generated --go-grpc_out=internal/server/grpc/generated api/AppService.proto

up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d --build

down:
	docker-compose -f ${DOCKER_COMPOSE_FILE} down --volumes
	
integration-tests:
	docker-compose -f ${DOCKER_COMPOSE_TEST_FILE}  -p ${APP_TEST_NAME} up --build --abort-on-container-exit --exit-code-from integration --attach integration
	docker-compose -f ${DOCKER_COMPOSE_TEST_FILE}  -p ${APP_TEST_NAME} down --volumes

reset-integration-tests:
	docker-compose -f ${DOCKER_COMPOSE_TEST_FILE}  -p ${APP_TEST_NAME} down --volumes	
