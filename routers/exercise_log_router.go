package routers

import "net/http"
import "gym-log/handlers"

func RegisterExerciseLogRoutes(mux *http.ServeMux, handler *handlers.ExerciseLogHandler) {
	mux.HandleFunc("/add", handler.AddExerciseLog)
	mux.HandleFunc("/delete", handler.DeleteExerciseLog)
	mux.HandleFunc("/update", handler.UpdateExerciseLog)
}
