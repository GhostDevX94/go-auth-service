# User Service

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/GhostDevX94/go-auth-service)](https://goreportcard.com/report/github.com/GhostDevX94/go-auth-service)

**Repository:** [https://github.com/GhostDevX94/go-auth-service](https://github.com/GhostDevX94/go-auth-service)

A RESTful user management service built with Go, providing user registration and authentication functionality with JWT tokens.

## 🚀 Features

- **User Registration** - Secure user registration with password hashing
- **User Authentication** - JWT-based login system
- **Password Security** - Bcrypt password hashing
- **Input Validation** - Request validation using go-playground/validator
- **CORS Support** - Cross-origin resource sharing enabled
- **Request Logging** - Structured logging with logrus
- **Database Migrations** - SQL migration system
- **Clean Architecture** - Well-structured codebase following Go best practices

## 📋 Prerequisites

- **Go 1.24+** - [Download Go](https://golang.org/dl/)
- **PostgreSQL** - [Download PostgreSQL](https://www.postgresql.org/download/)
- **Golang-migrate** - [Install migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

## 🛠️ Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/GhostDevX94/go-auth-service
   cd go-auth-service
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up environment variables:**
   ```bash
   cp env.example .env
   ```
   
   Edit `.env` file with your configuration:
   ```env
   APP_PORT=:8080
   DATABASE_URL=postgres://username:password@localhost:5432/user_service?sslmode=disable
   JWT_SECRET=your-super-secret-jwt-key-here
   ```

4. **Run database migrations:**
   ```bash
   make migrate-up
   ```

## 🚀 Running the Service

### Development Mode

**Build and run:**
```bash
make run
```

**Or run directly:**
```bash
go run cmd/api/main.go
```

### Production Mode

**Build the binary:**
```bash
make build
```

**Run the binary:**
```bash
./bin/myapp
```

### Using Makefile Commands

```bash
# Build the application
make build

# Run the application
make run

# Stop the application
make stop

# Restart the application
make restart

# Clean build artifacts
make clean
```

## 🗄️ Database Management

### Migration Commands

```bash
# Apply all pending migrations
make migrate-up

# Rollback last migration
make migrate-down

# Apply specific number of migrations
make migrate-steps

# Rollback specific number of migrations
make migrate-rollback

# Create new migration
make create-migration
```

### Database Schema

The service uses the following database schema:

```sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    phone VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
```

## 📡 API Endpoints

### Base URL
```
http://localhost:8080
```

### Endpoints

#### 1. User Registration
```http
POST /register
Content-Type: application/json

{
    "name": "John",
    "surname": "Doe",
    "email": "john.doe@example.com",
    "phone": "+1234567890",
    "password": "securepassword123"
}
```

**Response:**
```json
{
    "id": 1,
    "name": "John",
    "surname": "Doe",
    "email": "john.doe@example.com",
    "phone": "+1234567890"
}
```

#### 2. User Login
```http
POST /login
Content-Type: application/json

{
    "email": "john.doe@example.com",
    "password": "securepassword123"
}
```

**Response:**
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `APP_PORT` | Server port | `:8080` | No |
| `DATABASE_URL` | PostgreSQL connection string | - | Yes |
| `JWT_SECRET` | Secret key for JWT signing | - | Yes |

### Example .env file
```env
APP_PORT=:8080
DATABASE_URL=postgres://username:password@localhost:5432/user_service?sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-here
```

## 🏗️ Project Structure

```
user-service/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── configs/
│   │   └── Application.go       # Application configuration
│   ├── db/
│   │   └── connect.go           # Database connection
│   ├── dto/
│   │   └── user.go              # Data Transfer Objects
│   ├── handler/
│   │   ├── handler.go           # HTTP handler interface
│   │   ├── routes.go            # Route definitions
│   │   └── user-handler.go      # User-specific handlers
│   ├── middleware/
│   │   └── middleware.go        # HTTP middleware (CORS, logging)
│   ├── model/
│   │   └── User.go              # Domain models
│   ├── repository/
│   │   ├── repositories.go      # Repository interfaces
│   │   └── UserRepository.go    # User data access layer
│   └── service/
│       ├── service.go           # Service interfaces
│       └── UserService.go       # Business logic
├── migrations/
│   ├── 000001_create_users.up.sql
│   └── 000001_create_users.down.sql
├── pkg/
│   ├── hash.go                  # Password hashing utilities
│   ├── jwt.go                   # JWT token utilities
│   └── response.go              # HTTP response utilities
├── go.mod                       # Go module file
├── go.sum                       # Go module checksums
├── Makefile                     # Build and deployment scripts
└── README.md                    # This file
```

## 🔒 Security Features

- **Password Hashing**: Passwords are hashed using bcrypt with cost factor 14
- **JWT Tokens**: Secure authentication using JSON Web Tokens
- **Input Validation**: All user inputs are validated using go-playground/validator
- **CORS Protection**: Cross-origin requests are properly handled
- **Request Logging**: All requests are logged for monitoring and debugging

## 🧪 Testing

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/service
```

## 📊 Monitoring

The service includes structured logging for monitoring:

- **Request Logging**: All HTTP requests are logged with method, path, and body
- **Error Logging**: All errors are logged with context
- **Performance**: Request timing and performance metrics

## 🚀 Deployment

### Docker (Recommended)

1. **Build Docker image:**
   ```bash
   docker build -t user-service .
   ```

2. **Run container:**
   ```bash
   docker run -p 8080:8080 --env-file .env user-service
   ```

### Manual Deployment

1. **Build for target platform:**
   ```bash
   GOOS=linux GOARCH=amd64 make build
   ```

2. **Deploy binary and configuration files**

3. **Run with process manager (systemd, supervisor, etc.)**

## 🔧 Development

### Adding New Features

1. **Create migration for database changes:**
   ```bash
   make create-migration
   ```

2. **Update models in `internal/model/`**

3. **Add repository methods in `internal/repository/`**

4. **Implement business logic in `internal/service/`**

5. **Create handlers in `internal/handler/`**

6. **Add routes in `internal/handler/routes.go`**

### Code Style

- Follow Go conventions and best practices
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions small and focused
- Use interfaces for dependency injection

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions:
- Create an issue in the repository
- Contact the development team
- Check the documentation

## 🔄 Version History

- **v1.0.0** - Initial release with user registration and authentication
- **v1.1.0** - Added JWT authentication and improved security
- **v1.2.0** - Added middleware and structured logging 