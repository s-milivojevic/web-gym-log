package main

import (
	"context"
	"fmt"
	"gym-log/database"
	"gym-log/database/dao"
	"gym-log/handlers"
	"gym-log/routers"
	"net/http"
)

func main() {

	//TODO Client needs to be closed
	db, err := database.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Disconnect(context.Background())

	mainMux := http.NewServeMux()

	exercise_mux := http.NewServeMux()
	exerciseDao := dao.NewExerciseDao(db, "gym-log", "exercises")
	exercise_handler := handlers.NewExerciseHandler(exerciseDao)
	routers.RegisterExerciseRoutes(exercise_mux, exercise_handler)

	exercise_log_mux := http.NewServeMux()
	exercise_log_dao := dao.NewExerciseLogDao(db, "gym-log", "exercise_logs")
	exercise_log_handler := handlers.NewExerciseLogHandler(exercise_log_dao)
	routers.RegisterExerciseLogRoutes(exercise_log_mux, exercise_log_handler)

	training_mux := http.NewServeMux()
	training_dao := dao.NewTrainingDao(db, "gym-log", "trainings")
	training_handler := handlers.NewTrainingHandler(training_dao)
	routers.RegisterTrainingRoutes(training_mux, *training_handler)

	mainMux.Handle("/exercise", http.StripPrefix("/exercise", exercise_mux))
	mainMux.Handle("/exercise_log", http.StripPrefix("/exercise_log", exercise_log_mux))
	mainMux.Handle("/training", http.StripPrefix("/training", training_mux))

	fmt.Println("Server is listening on http://localhost:8080")
	http.ListenAndServe(":8080", mainMux)
}
