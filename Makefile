run:
	@go run ./cmd/tratodo --config=./configs/development.yaml

migration:
	@go run ./cmd/migrator --data=./storage/tratodo.db --migrations=./migrations