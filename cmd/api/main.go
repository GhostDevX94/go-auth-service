package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"user-service/configs"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/service"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error : %s", err)
	}

	connect, err := db.Connect()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return
	}

	services := service.NewServices(connect)

	app := configs.Application{
		Handler: handler.NewHandler(services),
	}

	server := &http.Server{
		Addr:    os.Getenv("APP_PORT"),
		Handler: handler.Route(app.Handler),
	}

	fmt.Println("Listening on port " + os.Getenv("APP_PORT"))

	log.Printf("Server is running on %s\n", os.Getenv("APP_PORT"))

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
