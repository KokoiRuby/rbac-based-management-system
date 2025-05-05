FROM bitnami/consul:1.20.6

# https://hub.docker.com/r/bitnami/consul/

# Optional: install curl
USER root

RUN apt-get update && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

USER 1001