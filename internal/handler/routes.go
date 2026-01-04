package handler

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Route(h *Handler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", h.UserRegister)
	mux.HandleFunc("POST /login", h.UserLogin)
	mux.HandleFunc("POST /update-token", h.UpdateToken)

	// Swagger UI
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
