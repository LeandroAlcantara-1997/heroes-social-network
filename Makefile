DB_URL = postgres://${DB_USER}:${DB_PASSWORD}@postgres-database:5432/${DB_NAME}?sslmode=disable
MODULE_NAME=$(shell grep ^module go.mod | cut -d " " -f2)
GIT_COMMIT_HASH=$(shell git rev-parse HEAD)
.INCLUDE: ./build/.env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' ./build/.env))

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

.PHONY: run
run:
	@go run main.go

.PHONY: hot
hot:
	@air

.PHONY: build
build:
	@go build main.go

.PHONY: test
test:
	@go test -coverpkg ./... -race -coverprofile coverage.out ./...

.PHONY: mock
mock: 
	@go generate ./...

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
	@go install github.com/swaggo/swag/cmd/swag@latest
	# @echo __Installing hot reload__
	# @curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
	# @air -v
	@git config --global --add safe.directory '*'
	@echo __Installing MockGen__
	@go install github.com/golang/mock/mockgen@v1.6.0
	@echo __Installing__
	@go install github.com/swaggo/swag/cmd/swag@latest


# files swagger
.PHONY: swag
swag:
	@swag init

# migrate
.PHONY: migration-up
migration-up:
	@echo ${DB_URL}
	@migrate -database ${DB_URL} -path config/migration up


.PHONY: migration-down
migration-down:
	@echo ${DB_URL}
	@migrate -database ${DB_URL} -path config/migration down

.PHONY: migration-drop
migration-drop:
	@echo ${DB_URL}
	@migrate -database ${DB_URL} -path config/migration drop -f

.PHONY: migration-create
migration-create:
	@migrate create -ext sql -dir config/migration/ -seq create_base_tables
