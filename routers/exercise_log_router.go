package routers

import "net/http"
import "gym-log/handlers"

func RegisterExerciseLogRoutes(mux *http.ServeMux, handler *handlers.ExerciseLogHandler) {
	mux.HandleFunc("/", handler.AddExerciseLog)
}
