run:
	@echo "Running server program..."
	go run cmd/app/main.go

create-migration:
	@echo "Generating empty migrations with name $(NAME)"
	migrate create -ext sql -dir db/migrations -seq $(NAME)

run-migration:
	@echo "Running migrations"
	go run tools/run-migration.go