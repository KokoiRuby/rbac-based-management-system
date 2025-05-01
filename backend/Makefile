PROJECT_NAME = rbac-mgmt-system
DOCKER_COMPOSE = docker compose -f ./docker/docker-compose.yaml -p $(PROJECT_NAME)

# Default target: runs the 'up' command
.DEFAULT_GOAL := up

# Target to start the Consul container in detached mode (-d)
up:
	@echo "Starting Consul container..mak."
	$(DOCKER_COMPOSE) up -d
	@echo "Consul container started. Access UI at http://localhost:8500"

# Target to stop and remove the Consul container and associated resources
down:
	@echo "Stopping and removing Consul container..."
	$(DOCKER_COMPOSE) down
	@echo "Consul container stopped and removed."

# Phony targets to avoid conflicts with files named 'up' or 'down'
.PHONY: up down
