DB_HOST ?= localhost
DB_PORT ?= $(shell grep "DB_PORT" .env | cut -d '=' -f2)
DB_NAME ?= $(shell grep "DB_NAME" .env | cut -d '=' -f2)
DB_USER ?= $(shell grep "DB_USERNAME" .env | cut -d '=' -f2)
DB_PASS ?= $(shell grep "DB_PASSWORD" .env | cut -d '=' -f2)

MYSQL_DSN ?= ${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}

# Version of migrations - this is optionally used on goto command
V?=

# Number of migrations - this is optionally used on up and down commands
N?=

migrate-setup:
	@if [ -z "$$(which migrate)" ]; then echo "Installing golang-migrate..."; go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; fi

.PHONY: migrate-up
migrate-up: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path $$(pwd)/database/migrations up $(N)

.PHONY: migrate-down
migrate-down: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path $$(pwd)/database/migrations down $(N)

.PHONY: migrate-to-version
migrate-to-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path $$(pwd)/database/migrations goto $(V)

.PHONY: drop-db
drop-db: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path $$(pwd)/database/migrations drop

.PHONY: force-version
force-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path $$(pwd)/database/migrations force $(V)

.PHONY: migration-version
migration-version: migrate-setup
	@ migrate -database 'mysql://$(MYSQL_DSN)?multiStatements=true' -path $$(pwd)/database/migrations version

.PHONY: start-dev
start-dev:
	docker compose up
