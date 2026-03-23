package routers

import "net/http"
import "gym-log/handlers"

func RegisterExerciseRoutes(mux *http.ServeMux, handler *handlers.ExerciseHandler) {
	mux.HandleFunc("/exercise", handler.GetExercises)
	mux.HandleFunc("/exercise/Add", handler.AddNewExercise)
}
