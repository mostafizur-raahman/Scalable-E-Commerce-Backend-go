# 🔑 Add this line FIRST so make can find 'migrate'
export PATH := $(PATH):$(shell go env GOPATH)/bin

DB_URL ?= postgres://postgres:root@localhost:5432/dokan_backend?sslmode=disable
MIGRATIONS_PATH ?= ./migrations

.PHONY: run up down

run:
	@go run main.go

up:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

down:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down