run:
	@go run ./cmd/tratodo

migration:
	@go run ./cmd/migrator --db-path=./storage/tratodo.db --migrations-path=./migrations