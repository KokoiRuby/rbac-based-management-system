PROJECT_NAME = rbac-mgmt-system
DOCKER_COMPOSE = docker compose -f ./docker/docker-compose.yaml -p $(PROJECT_NAME)

# Default target: runs the 'up' command
.DEFAULT_GOAL := up

# Target to start the Consul container in detached mode (-d)
up:
	$(DOCKER_COMPOSE) up -d
	@echo "Access Consul UI at http://localhost:8500"

# Target to stop and remove the Consul container and associated resources
down:
	$(DOCKER_COMPOSE) down

tidy:
	go mod tidy

run:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && if [ -z "$$GIN_MODE" ]; then export GIN_MODE="release"; fi && go run cmd/main.go

debug:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && if [ -z "$$GIN_MODE" ]; then export GIN_MODE="debug"; fi && go run cmd/main.go

migrate:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && go run cmd/main.go -m

# Phony targets to avoid conflicts with files named 'up' or 'down'
.PHONY: up down tidy run debug migrate
