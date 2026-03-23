package routers

import "net/http"
import "gym-log/handlers"

func RegisterExerciseRoutes(mux *http.ServeMux, handler *handlers.ExerciseHandler) {
	mux.HandleFunc("/", handler.GetExercises)
	mux.HandleFunc("/add", handler.AddNewExercise)
}
