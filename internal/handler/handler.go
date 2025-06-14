package handler

import (
	"github.com/go-playground/validator/v10"
	"user-service/internal/service"
)

type Handler struct {
	Services  *service.Services
	Validator *validator.Validate
}

func NewHandler(Services *service.Services) *Handler {
	return &Handler{
		Services:  Services,
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}
