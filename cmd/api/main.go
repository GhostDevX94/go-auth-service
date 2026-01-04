package main

// @title User Authentication Service API
// @version 1.0
// @description REST API for user authentication and token management
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https

import (
	"net/http"
	"os"
	"user-service/internal/configs"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/middleware"
	"user-service/internal/service"
	"user-service/pkg"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "user-service/docs"
)

func init() {
	pkg.SetupLogger()
}

func main() {

	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load environment variables")
	}

	dns := os.Getenv("DATABASE_URL")

	port := os.Getenv("APP_PORT")

	app := configs.Application{
		AppPort: port,
		DNS:     dns,
	}

	connect, err := db.Connect(app.DNS)

	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to database")
	}

	services := service.NewServices(connect)
	newHandler := handler.NewHandler(services)

	server := &http.Server{
		Addr:    app.AppPort,
		Handler: middleware.ApiMiddleware(handler.Route(newHandler)),
	}

	logrus.WithField("port", port).Info("Starting user service server")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatal("Server failed to start")
	}
}
