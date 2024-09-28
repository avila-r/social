## Overview

The main `Makefile` provides commands for managing the Docker services and running the Go application. It also includes commands to interface with a separate `Makefile.test.mk`, which manages the test environment. The use of Makefiles simplifies command execution and improves the development workflow.

- **Docker Compose Management**: Start, stop, and manage Docker containers for both development and test environments.
- **Log Management**: Display logs from the Docker services.

---

## Makefile Commands

The following table summarizes the available commands in the main `Makefile`:

| Command     | Description                                            |
|-------------|--------------------------------------------------------|
| `make run`  | Starts the Go application.                             |
| `make start`| Starts the Docker Compose services defined in `docker-compose.yml`. |
| `make stop` | Stops the running Docker Compose services.            |
| `make kill` | Removes containers, volumes, and images associated with the Docker Compose services. |
| `make logs` | Displays the logs of the running services in real-time. |
| `make test_start` | Starts the test environment using `Makefile.test.mk`. |
| `make test_stop` | Stops the test environment using `Makefile.test.mk`. |
| `make test_kill` | Removes test environment resources using `Makefile.test.mk`. |
| `make test_logs` | Displays logs for the test environment using `Makefile.test.mk`. |

---

## Makefile.test.mk Commands

The following table summarizes the available commands in the `Makefile.test.mk`:

| Command     | Description                                            |
|-------------|--------------------------------------------------------|
| `make start`| Starts the Docker Compose services defined in `docker-compose.test.yml` for the test environment. |
| `make stop` | Stops the running Docker Compose services in the test environment. |
| `make kill` | Removes containers, volumes, and images associated with the test environment services. |
| `make logs` | Displays logs from the test environment services in real-time. |

---

## Additional Notes

- Ensure that you have Docker and Docker Compose installed and running on your machine.
- The Makefile commands utilize the `docker-compose` commands internally, so familiarize yourself with Docker Compose for better understanding.
- You can add additional commands to the Makefiles as necessary to suit your development workflow.