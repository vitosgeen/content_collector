build:
	go build ./cmd/collector/main.go

run:
	go run ./cmd/collector/main.go

test:
	go test -v -cover ./...

lint:
	gofumpt -w -s ./
	golangci-lint run --fix

generate:
	mockgen -source=./internal/services/collector.go -destination=./internal/services/mocks.go -package=services
	mockgen -source=./internal/repository/irepository.go -destination=./internal/repository/mocks.go -package=repositories

# ==============================================================================
# Docker compose commands

dev:
	echo "Starting docker dev environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.dev.yml up content_collector --build -d

prod:
	echo "Starting docker prod environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.prod.yml up content_collector --build

local:
	echo "Starting docker local environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.local.yml up content_collector --build

stop:
	echo "Stopping docker environment"
	docker-compose stop

dev-down:
	echo "Stopping docker dev environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.dev.yml down

prod-down:
	echo "Stopping docker prod environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.prod.yml down

local-down:
	echo "Stopping docker local environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.local.yml down