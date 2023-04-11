SHELL := /bin/bash

.PHONY: run
run: build
	go run ./cmd/account-producer-service/main.go

.PHONY: build run compose-up compose-down compose-infra-down compose-infra-up
compose-infra-up:
	docker-compose -f local-dev/docker-compose.yaml --profile infra up -d
compose-infra-down:
	docker-compose -f local-dev/docker-compose.yaml --profile infra down
