MIGRATE := sql-migrate

.PHONY: dbmigrate
dbmigrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: newmigration
newmigration: ## create an empty migration file with provided NAME
	@echo "Creating empty migration with name $(NAME)"
	@$(MIGRATE) new $(NAME)

.PHONY: dbstatus
dbstatus: ## check all migrations status if they're applied to db
	@echo "Getting db migration status..."
	@$(MIGRATE) status

.PHONY: test
test: ## run unit tests
	@echo "Running all tests"
	@go test ./... -v -count=-11