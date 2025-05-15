PROJECT_NAME = rbac-mgmt-system
DOCKER_COMPOSE = docker compose -f ./docker/docker-compose.yaml -p $(PROJECT_NAME)
USER = tutorj

.PHONY: build
build:
	$(DOCKER_COMPOSE) build

.PHONY: up
up: build
	$(DOCKER_COMPOSE) up -d
	@echo "Access Consul UI at http://localhost:8500"

.PHONY: down
down:
	$(DOCKER_COMPOSE) down

.PHONY: tidy
tidy:
	sudo chown -R $(USER):$(USER) ./docker/mysql/data
	go mod tidy
	sudo chown -R 1001:1001 ./docker/mysql/data

.PHONY: run
run:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && if [ -z "$$GIN_MODE" ]; then export GIN_MODE="release"; fi && go run cmd/main.go

.PHONY: debug
debug:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && if [ -z "$$GIN_MODE" ]; then export GIN_MODE="debug"; fi && go run cmd/main.go

.PHONY: migrate
migrate:
	if [ -z "$$ENV" ]; then export ENV="dev"; fi && if [ -z "$$CONSUL_HTTP_ADDR" ]; then export CONSUL_HTTP_ADDR="http://localhost:8500"; fi && go run cmd/main.go -m

.PHONY: ssl
ssl:
	mkdir -p ssl/certs
	cfssl gencert -initca ssl/ca-csr.json | cfssljson -bare ssl/certs/ca
	cfssl gencert \
		  -ca ssl/certs/ca.pem \
		  -ca-key ssl/certs/ca-key.pem \
		  -config ssl/ca-config.json \
		  ssl/gin-csr.json | cfssljson -bare ssl/certs/gin/gin

.PHONY: clean
clean:
	rm -rf ssl/certs
