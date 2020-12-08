SHELL:=/bin/bash

dev-up:
	@docker-compose -f dev.docker-compose.yml up --build

dev-down:
	@docker-compose -f dev.docker-compose.yml down


tests:
	@go test -v ./...