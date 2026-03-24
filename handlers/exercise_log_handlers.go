package handlers

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"

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

func (h *ExerciseLogHandler) DeleteExerciseLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method is supported.", http.StatusMethodNotAllowed)
		return
	}
	var string_id string
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&string_id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := primitive.ObjectIDFromHex(string_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.dao.DeleteExerciseLog(r.Context(), bson.ObjectID(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *ExerciseLogHandler) UpdateExerciseLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is supported.", http.StatusMethodNotAllowed)
		return
	}
	var request models.UpdateExerciseLogRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.dao.UpdateExerciseLog(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
