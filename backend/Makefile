PROJECT_NAME = rbac-mgmt-system
DOCKER_COMPOSE = docker compose -f ./docker/docker-compose.yaml -p $(PROJECT_NAME)

.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d
	@echo "Access Consul UI at http://localhost:8500"

.PHONY: down
down:
	$(DOCKER_COMPOSE) down

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run
run:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && if [ -z "$$GIN_MODE" ]; then export GIN_MODE="release"; fi && go run cmd/main.go

.PHONY: debug
debug:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && if [ -z "$$GIN_MODE" ]; then export GIN_MODE="debug"; fi && go run cmd/main.go

.PHONY: migrate
migrate:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && go run cmd/main.go -m

.PHONY: ssl
ssl:
	mkdir -p ssl/certs
	cfssl gencert -initca ssl/ca-csr.json | cfssljson -bare ssl/certs/ca
	cfssl gencert \
		  -ca ssl/certs/ca.pem \
		  -ca-key ssl/certs/ca-key.pem \
		  -config ssl/ca-config.json \
		  ssl/gin-csr.json | cfssljson -bare ssl/certs/gin

.PHONY: clean
clean:
	rm -rf ssl/certs
