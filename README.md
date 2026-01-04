# User Service

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/GhostDevX94/go-auth-service)](https://goreportcard.com/report/github.com/GhostDevX94/go-auth-service)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://www.docker.com/)
[![Tests](https://img.shields.io/badge/Tests-Passing-green.svg)](https://github.com/GhostDevX94/go-auth-service/actions)

**Repository:** [https://github.com/GhostDevX94/go-auth-service](https://github.com/GhostDevX94/go-auth-service)

##  About the Project

User Service is a modern RESTful microservice for authentication and user management built with Go. The service provides complete functionality for user registration, authentication using JWT tokens, and access token refresh.

The project demonstrates Go development best practices, including clean architecture, comprehensive testing, structured logging, and modern DevOps practices.

## Features

- **User Registration** - Secure registration with Bcrypt password hashing
- **Authentication** - JWT-based login system with access and refresh tokens
- **Token Refresh** - Mechanism for refreshing access tokens via refresh tokens
- **Password Security** - Bcrypt hashing with configurable cost
- **Data Validation** - Automatic request validation using go-playground/validator
- **CORS Support** - Configured cross-origin request support
- **Structured Logging** - Logging using logrus
- **Database Migrations** - Migration system for PostgreSQL
- **Clean Architecture** - Layered separation: handlers, services, repositories
- **Testing** - Unit tests with mocks and integration tests
- **Docker Support** - Containerization with Docker Compose
- **API Documentation** - Interactive Swagger/OpenAPI documentation

## Architecture

The project follows clean architecture principles with clear separation of concerns:

```
user-service/
├── cmd/api/              # Application entry point
├── internal/
│   ├── handler/          # HTTP handlers (controllers)
│   ├── service/          # Business logic
│   ├── repository/       # Database operations
│   ├── model/            # Data models
│   ├── dto/              # Data Transfer Objects
│   ├── middleware/       # HTTP middleware
│   ├── db/               # Database connection
│   └── configs/          # Configuration
├── pkg/                  # Reusable utilities
├── migrations/           # SQL migrations
├── docs/                 # Swagger documentation
└── test/                 # Tests
```

### Swagger UI

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/swagger/index.html` | Interactive API documentation |

## 🚦 Quick Start

### Prerequisites

- Go 1.24 or higher
- PostgreSQL 14+
- Make
- migrate CLI (for migrations)
- Docker and Docker Compose (optional)

### Installation

1. **Clone the repository:**
```bash
git clone https://github.com/GhostDevX94/go-auth-service.git
cd go-auth-service
```

2. **Configure environment variables:**
```bash
cp env.example .env
```

Edit the `.env` file:
```env
DATABASE_URL=postgres://user:password@localhost:5432/userdb?sslmode=disable
APP_PORT=:8080
JWT_SECRET=your-secret-key-here
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=7d
```

3. **Install dependencies:**
```bash
go mod download
```

4. **Start PostgreSQL** (if using Docker):
```bash
docker-compose up -d postgres
```

5. **Apply migrations:**
```bash
make migrate-up
```

6. **Install Swagger CLI** (for documentation generation):
```bash
make swagger-install
```

7. **Generate Swagger documentation:**
```bash
make swagger-gen
```

8. **Start the service:**
```bash
make run
```

The service will be available at `http://localhost:8080`

## Running with Docker

```bash
# Start all services (application + PostgreSQL)
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Swagger UI

Open in browser: `http://localhost:8080/swagger/index.html`

The interactive documentation allows you to:
- View all available endpoints
- Explore request and response schemas
- Test the API directly from the browser

##  Makefile Commands

```bash
# Build
make build          # Build binary file
make clean          # Remove binary file

# Run
make run            # Build and run in background
make stop           # Stop service
make restart        # Restart service

# Migrations
make create-migration    # Create new migration
make migrate-up         # Apply all migrations
make migrate-down       # Rollback last migration
make migrate-steps      # Apply N migrations
make migrate-rollback   # Rollback N migrations

# Swagger
make swagger-install    # Install swag CLI
make swagger-gen       # Generate documentation
make swagger-clean     # Remove generated files

# Testing
make test              # Run tests
```

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -v -cover ./...

# Run tests for specific package
go test -v ./internal/service/...
```