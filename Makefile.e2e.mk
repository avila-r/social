# Docker Compose file name
DOCKER_COMPOSE_FILE = docker-compose.e2e.yml
WAIT_TIME = 5 # seconds
APP_PORT = 8888

setup:
	@echo "Starting Docker Compose in detached mode..."
	@docker-compose --env-file .env.e2e -f $(DOCKER_COMPOSE_FILE) up -d

	@echo "Waiting for $(WAIT_TIME) seconds for services to be ready..."
	@sleep $(WAIT_TIME)

	@echo "Starting Go application..."
	@go run cmd/main.go &

test:
	@echo "Running tests..."
	@go test ./e2e/...

	@echo "Stopping Docker Compose..."
	@docker-compose --env-file .env.e2e -f $(DOCKER_COMPOSE_FILE) down

	@echo "E2E tests completed."

clean:
	@echo "Cleaning up environment."
	@docker-compose --env-file .env.e2e -f $(DOCKER_COMPOSE_FILE) down --volumes --rmi all
	@echo "Checking if any service is running on port $(APP_PORT)..."
	@PID=$$(lsof -t -i:$(APP_PORT)); \
	if [ -n "$$PID" ]; then \
		echo "Stopping go application running on port $(APP_PORT) (PID: $$PID)..."; \
		kill $$PID; \
		echo "Go application on port $(APP_PORT) stopped."; \
	else \
		echo "No go application running on port $(APP_PORT)."; \
	fi