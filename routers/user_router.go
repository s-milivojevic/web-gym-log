package routers

import (
	"gym-log/handlers"
	"net/http"
)

func RegisterUserRoutes(mux *http.ServeMux, handler *handlers.UserHandler) {
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/add_training", handler.AddTraining)
}
