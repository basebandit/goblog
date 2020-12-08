SHELL:=/bin/bash

dev-up:
	@docker-compose -f dev.docker-compose.yml up --build

dev-down:
	@docker-compose -f dev.docker-compose.yml down

prod-up:
	@docker-compose -f prod.docker-compose.yml up --build

prod-down:
	@docker-compose -f prod.docker-compose.yml down

tests:
	@go test -v ./...