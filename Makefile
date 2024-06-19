run:
	@go run ./cmd/tratodo --config=./configs/development.yaml

migration:
	@go run ./cmd/migrator --db-path=./storage/tratodo.db --migrations-path=./migrations