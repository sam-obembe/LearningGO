#!/bin/sh
podman run --rm -p 80:8080 -e SWAGGER_JSON=/app/openapi.json -v /:/app docker.io/swaggerapi/swagger-ui