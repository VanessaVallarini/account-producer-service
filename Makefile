SHELL := /bin/bash

.PHONY: run
run: build
	@ENV_PROFILE=local ./bin/app/authorization-service

.PHONY: build run compose-up compose-down compose-infra-down compose-infra-up
compose-infra-up:
	docker-compose -f local/docker-compose.yml --profile infra up -d
compose-infra-down:
	docker-compose -f local/docker-compose.yml --profile infra down
