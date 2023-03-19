DB_URL = postgres://${DB_USER}:${DB_PASSWORD}@postgres-database:5432/${DB_NAME}?sslmode=disable


# docker
.PHONY: docker-up
docker-up:
	@docker compose -f ./build/docker-compose.yaml up

.PHONY: docker-build
docker-build:
	@docker compose -f ./build/docker-compose.yaml build


# lint
.PHONY: lint
lint:
	@staticcheck ./...

.PHONY: build
build:
	@go build main.go


# setup
.PHONY: setup
setup:
	@echo __Installing migrate__
	@curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
	@sudo apt-get update
	@sudo apt-get install migrate
	@echo __Installing staticcheck__
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@staticcheck --version


# migrate
.PHONY: migration-up
migration-up:
	@echo ${DB_URL}
	@migrate -database ${DB_URL} -path migration up


.PHONY: migration-down
migration-down:
	@echo ${DB_URL}
	@migrate -database ${DB_URL} -path migration down

.PHONY: migration-drop
migration-drop:
	@echo ${DB_URL}
	@migrate -database ${DB_URL} -path migration drop -f

.PHONY: migration-create
migration-create:
	@migrate create -ext sql -dir migration/ -seq create_base_tables
