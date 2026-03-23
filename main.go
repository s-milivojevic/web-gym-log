package main

import (
	"fmt"
	"gym-log/handlers"
	"gym-log/routers"
	"net/http"
)

func main() {
	exercise_mux := http.NewServeMux()
	exercise_handler := handlers.NewExerciseHandler()
	routers.RegisterExerciseRoutes(exercise_mux, exercise_handler)
	fmt.Println("Server is listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
