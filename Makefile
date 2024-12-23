# /Makefile
.PHONY: up down build logs clean dev

# Production commands
up:
	docker-compose -f deploy/docker-compose.yml up -d

down:
	docker-compose -f deploy/docker-compose.yml down

build:
	docker-compose -f deploy/docker-compose.yml build

logs:
	docker-compose -f deploy/docker-compose.yml logs -f

clean:
	podman pod rm -f pod_deploy || true
	podman rm -f -a || true
	podman network rm formease_network || true

dev: clean
	podman network create formease_network || true
	podman-compose -f deploy/docker-compose.dev.yml up --build

dev-down:
	docker-compose -f deploy/docker-compose.dev.yml down

dev-build:
	docker-compose -f deploy/docker-compose.dev.yml build

# Database commands
migrate-up:
	docker-compose -f deploy/docker-compose.yml exec backend ./main migrate up

migrate-down:
	docker-compose -f deploy/docker-compose.yml exec backend ./main migrate down

# Testing
test-backend:
	cd backend && go test ./...

test-frontend:
	cd frontend && npm test
