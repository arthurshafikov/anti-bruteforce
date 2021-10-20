BIN := "./bin/app"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -a -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd

run: build
	 $(BIN)

.PHONY: build run test

enterdb:
	docker exec -it postgres psql -U homestead;

up:
	docker-compose up --build

down:
	docker-compose down --volumes
