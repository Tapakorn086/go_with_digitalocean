.PHONY: db-up db-down create-table

db-up:
	@echo "Running database migrations..."
	goose postgres "host=localhost port=5432 user=postgres password=postgres dbname=learning sslmode=disable" up

db-down:
	@echo "Rolling back database migrations..."
	goose postgres "host=localhost port=5432 user=postgres password=postgres dbname=learning sslmode=disable" down

create-table:
	@echo "Creating todos table..."
	goose -dir db/migrate create ${name} sql

generate-sqlc:
	@echo "Generating sqlc..."
	sqlc generate