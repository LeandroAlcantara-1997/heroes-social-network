DB_URL = postgres://${DB_USER}:${DB_PASSWORD}@postgres-database:5432/${DB_NAME}?sslmode=disable
MODULE_NAME=$(shell grep ^module go.mod | cut -d " " -f2)
GIT_COMMIT_HASH=$(shell git rev-parse HEAD)
LD_FLAGS=-ldflags="-X $(MODULE_NAME)/internal/config.gitCommitHash=$(GIT_COMMIT_HASH)"

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
	@echo __Installing hot reload__
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
	@air -v
	@chmod -R a+w /go/pkg
	@git config --global --add safe.directory '*'
	@go install github.com/golang/mock/mockgen@v1.6.0


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
