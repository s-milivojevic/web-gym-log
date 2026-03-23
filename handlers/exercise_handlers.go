package handlers

import (
	"encoding/json"
	"gym-log/database/dao"
	"gym-log/models"
	"net/http"
)

type ExerciseHandler struct {
	dao *dao.ExerciseDao
}

type createExerciseRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewExerciseHandler(exerciseDao *dao.ExerciseDao) *ExerciseHandler {
	return &ExerciseHandler{dao: exerciseDao}
}

func (h *ExerciseHandler) AddNewExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var body createExerciseRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil {
		http.Error(w, "invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if body.Name == "" || body.Type == "" {
		http.Error(w, "name and type are required", http.StatusBadRequest)
		return
	}

	exercise := models.Exercise{
		Name: body.Name,
		Type: body.Type,
	}

	if err := h.dao.AddExercise(r.Context(), exercise); err != nil {
		http.Error(w, "failed to save exercise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("exercise created"))
}

func (h *ExerciseHandler) GetExercises(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	//TODO Call dao func for getting all exercises from database
	w.Write([]byte("get exercises"))
}
