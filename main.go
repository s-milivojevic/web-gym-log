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

	mainMux.Handle("/exercise/", http.StripPrefix("/exercise", exercise_mux))

	fmt.Println("Server is listening on http://localhost:8080")
	http.ListenAndServe(":8080", mainMux)
}
