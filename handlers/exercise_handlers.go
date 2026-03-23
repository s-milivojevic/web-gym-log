package handlers

import "net/http"

type ExerciseHandler struct {
}

func NewExerciseHandler() *ExerciseHandler {
	return &ExerciseHandler{}
}

func (h *ExerciseHandler) AddNewExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	//TODO Call dao func for adding new exercise in database
	w.Write([]byte("add new exercise"))
}

func (h *ExerciseHandler) GetExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	//TODO Call dao func for getting all exercises from database
	w.Write([]byte("get exercises"))
}
