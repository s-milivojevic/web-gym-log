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

func NewExerciseHandler(exerciseDao *dao.ExerciseDao) *ExerciseHandler {
	return &ExerciseHandler{dao: exerciseDao}
}

func (h *ExerciseHandler) AddNewExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var body models.CreateExerciseRequest

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

	exercises, err := h.dao.GetExercises(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch exercises", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(exercises)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// TODO filter with id not with name or type
func (h *ExerciseHandler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var body models.UpdateExerciseRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&body); err != nil {
		http.Error(w, "invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if body.Name == "" {
		http.Error(w, "name type are required", http.StatusBadRequest)
		return
	}
	if body.NewType == "" && body.NewName == "" {
		http.Error(w, "new type or new name is required", http.StatusBadRequest)
		return
	}

	var updatedExercise models.Exercise
	updatedExercise, err := h.dao.UpdateExercise(r.Context(), body)
	if err != nil {
		http.Error(w, "failed to save exercise", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updatedExercise)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// TODO filter with id
func (h *ExerciseHandler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var exercise_name string
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&exercise_name); err != nil {
		http.Error(w, "invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.dao.DeleteExercise(r.Context(), exercise_name)
	if err != nil {
		http.Error(w, "failed to delete exercise", http.StatusInternalServerError)
		return
	}
	if result != 0 {
		http.Error(w, "there is no such exercise", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
