include .env

POSTGRES_HOST=127.0.0.1
POSTGRES_MIGRATIONS=./database/migrations
DATA_SOURCE_NAME=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)
OPTIONS=sslmode=disable

migration:
	@migrate create -ext=sql -dir=$(POSTGRES_MIGRATIONS) -tz America/Sao_Paulo $(name)

migration-up:
	@migrate -path=$(POSTGRES_MIGRATIONS) -database "$(DATA_SOURCE_NAME)?$(OPTIONS)" -verbose up

migration-down:
	@migrate -path=$(POSTGRES_MIGRATIONS) -database "$(DATA_SOURCE_NAME)?$(OPTIONS)" -verbose down

pg-connect:
	@psql -h $(POSTGRES_HOST) -p $(POSTGRES_PORT) -U $(POSTGRES_USER) -d $(POSTGRES_DB)


start:
	docker compose build
	docker compose up -d
	make migration-up

stop:
	docker compose down
