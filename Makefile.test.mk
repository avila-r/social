#
# TESTS
#
# Test Docker Compose file name
TEST_DOCKER_COMPOSE_FILE = docker-compose.test.yml

# Command to start the docker-compose.test services
start:
	@echo "Starting services for test environment..."
	docker-compose --env-file .env.test -f $(TEST_DOCKER_COMPOSE_FILE) up -d
	@echo "Services started successfully!"

# Command to pause (stop) the containers without removing them
stop:
	@echo "Stopping Docker Compose services..."
	docker-compose -f $(TEST_DOCKER_COMPOSE_FILE) stop
	@echo "Services stopped!"

# Command to completely remove containers, volumes, and images
kill:
	@echo "Removing containers, volumes, and images..."
	docker-compose -f $(TEST_DOCKER_COMPOSE_FILE) down --volumes --rmi all
	@echo "All containers, volumes, and images removed!"

# Command to show logs of services in real time
logs:
	@echo "Displaying services' logs..."
	docker-compose -f $(TEST_DOCKER_COMPOSE_FILE) logs -f