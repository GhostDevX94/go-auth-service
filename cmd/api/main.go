package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"user-service/internal/configs"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/middleware"
	"user-service/internal/service"
	"user-service/pkg"
)

func init() {
	pkg.SetupLogger()
}

func main() {

	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("❌ Failed to load environment variables")
	}

	connect, err := db.Connect()
	if err != nil {
		logrus.WithError(err).Fatal("❌ Failed to connect to database")
	}

	app := configs.Application{
		Handler: handler.NewHandler(
			service.NewServices(connect),
		),
	}

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = ":8080"
	}

	server := &http.Server{
		Addr:    port,
		Handler: middleware.ApiMiddleware(handler.Route(app.Handler)),
	}

	logrus.WithField("port", port).Info("🚀 Starting user service server")
	logrus.WithField("port", port).Info("📡 Server is listening for requests")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatal("❌ Server failed to start")
	}
}
