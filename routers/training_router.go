package routers

import (
	"gym-log/handlers"
	"net/http"
)

func RegisterTrainingRoutes(mux *http.ServeMux, handler handlers.TrainingHandler) {
	mux.HandleFunc("/", handler.GetUsersTrainings)
	mux.HandleFunc("/create", handler.CreateTraining)
	mux.HandleFunc("/add_exercise_log", handler.AddExerciseLog)
	mux.HandleFunc("/update_duration", handler.UpdateTrainingDuration)
	mux.HandleFunc("/update_date", handler.UpdateTrainingDate)
	mux.HandleFunc("/update_comment", handler.UpdateTrainingComment)
	mux.HandleFunc("/update_exercise_log", handler.UpdateTrainingExerciseLog)
}
