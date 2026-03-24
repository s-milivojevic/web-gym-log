package handlers

import (
	"encoding/json"
	"gym-log/database/dao"
	"gym-log/models"
	"net/http"
)

type ExerciseLogHandler struct {
	dao *dao.ExerciseLogDao
}

func NewExerciseLogHandler(dao *dao.ExerciseLogDao) *ExerciseLogHandler {
	return &ExerciseLogHandler{dao: dao}
}

func (h *ExerciseLogHandler) AddExerciseLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}
	var request models.CreateExerciseLogRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err, result := h.dao.AddExerciseLog(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result == nil {
		http.Error(w, "Couldn't insert exercise log", http.StatusNoContent)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ExerciseLogHandler) GetExerciseLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
		return
	}
	var exerciseLogs []models.ExerciseLog

	exerciseLogs, err := h.dao.GetAllExerciseLogs(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//TODO return exercise log
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(exerciseLogs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
