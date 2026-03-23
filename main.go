package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Gym log server")
	})

	fmt.Println("Server sluša na http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/workouts", workoutsHandler)
}
