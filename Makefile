run:
	JWT_SECRET="secret" go run ./cmd/tratodo --config=./configs/development.yaml

migration:
	@go run ./cmd/migrator --data=./storage/tratodo.db --migrations=./migrations

swag:
	@swag init -g cmd/tratodo/main.go