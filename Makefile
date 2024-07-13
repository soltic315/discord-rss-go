.PHONY: setup migrate-up migrate-down gen-models

setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/volatiletech/sqlboiler@latest
	go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql@latest

migrate-up:
	migrate -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE} -path migrations up

migrate-down:
	migrate -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE} -path migrations down

migrate-force:
	docker-compose exec app migrate -database "postgres://postgres:password@db:5432/sample?sslmode=disable" -path ./src/database/migrations force ${VERSION}

gen-models:
	sqlboiler psql
	go mod tidy