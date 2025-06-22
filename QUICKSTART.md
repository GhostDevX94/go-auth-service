# Quick Start Guide

This guide will help you get the User Service up and running in minutes.

## 🚀 Prerequisites

Make sure you have the following installed:
- Go 1.24+
- PostgreSQL
- golang-migrate

## ⚡ Quick Setup

### 1. Clone and Setup
```bash
git clone https://github.com/GhostDevX94/go-auth-service
cd go-auth-service
go mod download
```

### 2. Database Setup
```bash
# Create PostgreSQL database
createdb user_service

# Copy environment file
cp env.example .env

# Edit .env with your database credentials
# DATABASE_URL=postgres://username:password@localhost:5432/user_service?sslmode=disable
```

### 3. Run Migrations
```bash
make migrate-up
```

### 4. Start the Service
```bash
make run
```

The service will be available at `http://localhost:8080`

## 🧪 Test the API

### Register a User
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John",
    "surname": "Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "password": "password123"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

## 🔧 Common Issues

### Database Connection Error
- Check if PostgreSQL is running
- Verify DATABASE_URL in .env file
- Ensure database exists

### JWT Secret Error
- Make sure JWT_SECRET is set in .env
- Use a strong, random secret key

### Port Already in Use
- Change APP_PORT in .env file
- Or stop the existing service: `make stop`

## 📚 Next Steps

- Read the full [README.md](README.md) for detailed documentation
- Check the API endpoints in the main README
- Explore the project structure
- Add your own features following the development guide 