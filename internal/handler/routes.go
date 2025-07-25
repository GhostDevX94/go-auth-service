package handler

import (
	"net/http"
)

func Route(h *Handler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", h.UserRegister)
	mux.HandleFunc("POST /login", h.UserLogin)

	return mux
}
