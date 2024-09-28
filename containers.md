## Docker Compose Files Overview

### `docker-compose.yml`
This file is for the main development environment. It defines a PostgreSQL service (with optional Redis and App services commented out). 

### `docker-compose.test.yml`
This file is for the test environment. It includes the same PostgreSQL service but is configured to work with environment variables from a `.env.test` file.

---

## Environment Variables Setup

Before running the Docker Compose commands, ensure you have the correct environment variable files in the root directory.

### 1. `.env` (for development)

Create a `.env` file with the following content for your development environment:

```bash
POSTGRES_USER=app-default-user
POSTGRES_PASSWORD=app-default-password
POSTGRES_DB_NAME=app-db
```

### 2. `.env.test` (for unit tests)

Create a .env.test file with the following content for your test environment:

```bash
POSTGRES_USER=app-test-user
POSTGRES_PASSWORD=app-test-password
POSTGRES_DB_NAME=app-test-db
```

## Running the Development Environment

To run the development environment, use the docker-compose.yml file. This will spin up the PostgreSQL service and allow you to work on the application locally.

### 1. Ensure you have the .env file in the project root with the correct environment variables.
### 2. Run the following command to start the services:

```bash
docker-compose up -d
```

### 3. Run the following command to down the services:
```bash
docker-compose down # --volumes --rmi all # (optional)
```

## Running the Unit Test Environment

For testing, you should use the docker-compose.test.yml file along with the .env.test file. This ensures the correct environment variables are used during testing.

```bash
docker-compose --env-file .env.test -f docker-compose.test.yml up -d
```
