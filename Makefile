DBHost=$(shell grep "DB_HOST" .env | cut -d '=' -f2)
DBPort=$(shell grep "DB_PORT" .env | cut -d '=' -f2)
DBName=$(shell grep "DB_NAME" .env | cut -d '=' -f2)
DBUser=$(shell grep "DB_USERNAME" .env | cut -d '=' -f2)
DBPass=$(shell grep "DB_PASSWORD" .env | cut -d '=' -f2)

.PHONY: start-dev
start-dev:
	docker compose up

.PHONY: migrations
migrations:
	docker run -it --network olist/backend olist/migrator -path=/migrations/ -database "mysql://${DBUser}:${DBPass}@tcp(${DBHost}:${DBPort})/${DBName}" $(command)