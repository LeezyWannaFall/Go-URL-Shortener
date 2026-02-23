# Go URL Shortener

A fast, lightweight, and containerized URL shortening service written in Go. The application is built using Clean Architecture principles and supports interchangeable storage engines (In-Memory and PostgreSQL), making it highly flexible for both rapid local development and production-like deployments.

## Features

* **URL Shortening:** Generate unique, compact aliases for long URLs.
* **Redirection:** Fast HTTP redirects from short links to original destinations.
* **Pluggable Storage:** Seamlessly switch between PostgreSQL (persistent) and In-Memory (volatile) storage using environment variables.
* **Containerized:** Fully configured with Docker and Docker Compose for zero-setup deployments.
* **Automated Migrations:** Database schemas are automatically applied on startup.
* **Tested:** High test coverage across handlers, services, and repositories using table-driven unit tests.

## Architecture and Project Structure

The project follows standard Go project layout and Clean Architecture principles, ensuring separation of concerns.
```text
Go-URL-Shortener
├── cmd
│   └── shortener
│       └── main.go              # Application entry point
├── config.yaml                  # Application configuration (ports, db credentials)
├── .env                         # Memory change file
├── docker-compose.yml           # Infrastructure definition
├── Dockerfile                   # Application container build instructions
├── internal                     # Private application and business logic
│   ├── config                   # Configuration loader and parser
│   ├── handler                  # HTTP delivery layer (routing, request/response)
│   ├── model                    # Core business entities
│   ├── repository               # Data access layer (Postgres and Memory implementations)
│   └── service                  # Core business logic and helper functions
├── Makefile                     # Task automation commands
└── migrations                   # Raw SQL files for database schema
```

## Prerequisites
To run this project, you need to have the following installed on your machine:
- **Docker**
- **Docker Compose**
- **Make (Optional, but recommended for ease of use)**

## Configuration
The application is configured using a combination of ```config.yaml``` and environment variables.

Edit a ```.env``` file in the root directory to define the storage type and database secrets:
```
# Choose storage: "postgres" or "memory"
STORAGE_TYPE=postgres

# PostgreSQL Configuration
POSTGRES_USER=url-shortener
POSTGRES_PASSWORD=your_secure_password
POSTGRES_DB=urldb
```

## Getting Started
You can manage the application lifecycle using the provided Makefile.

**1. Start the application:**
```bash
make service_start
```
**2. You can create, delete and check postgres database with commands:**
```bash
make create_table
```
```bash
make delete_table
```
```bash
make check_table
```
**3. Stop the application and remove containers:**
```bash
make down_service
```

## API Documentation
By default, the application runs on ```http://localhost:8080```

**1. Create a Short URL**
Converts a full URL into a short alias.
- **Endpoint:** ```POST /shorten```

- **Content-Type:** ```application/json```

- **Request Body:**
    ```json
    {
        "full": "https://github.com/leezywannafall"
    }
    ```
**2.Get Original URL and Redirect**

Get original URL and redirects the client to the original URL.
- **Endpoint:** ```GET /{short}```

- **Response:** ```302 Found``` (Redirects to the original URL)

## Testing
The project utilizes standard Go table-driven tests with mocked dependencies to ensure reliability.

To run all unit tests across the project:
```bash
go test -v ./... 
```