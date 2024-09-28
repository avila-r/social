# Docker Compose file name
DOCKER_COMPOSE_FILE = docker-compose.yml

# Command to run your go app
run:
	@echo "Starting go application..."
	go run cmd/main.go

# Command to start the docker-compose services
start:
	@echo "Starting services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "Services started successfully!"

# Command to pause (stop) the containers without removing them
stop:
	@echo "Stopping Docker Compose services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) stop
	@echo "Services stopped!"

# Command to completely remove containers, volumes, and images
kill:
	@echo "Removing containers, volumes, and images..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes --rmi all
	@echo "All containers, volumes, and images removed!"

# Command to show logs of services in real time
logs:
	@echo "Displaying services' logs..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

#
# Tests
#
test_start:
	@echo "Using test environment"
	@make -f Makefile.test.mk start -s

test_stop:
	@echo "Stopping test environment"
	@make -f Makefile.test.mk stop -s

test_kill:
	@echo "Killing test environment"
	@make -f Makefile.test.mk kill -s

test_logs:
	@echo "Get logging for test environment"
	@make -f Makefile.test.mk logs -s