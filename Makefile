.PHONY: build
build:
	@docker compose -f ./build/docker-compose.yaml up

.PHONY: migration
migration:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations up

.PHONY: lint
lint:
	@staticcheck ./...

.PHONY: build
build:
	@go build -o ./bin/api ./cmd/api