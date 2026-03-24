package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Exercise struct {
	ID   bson.ObjectID `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
	Type string        `bson:"type" json:"type"`
}

type ExerciseLog struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Exercises Exercise
	Reps      []int
	Sets      []int
}

type Training struct {
	ID        string
	Date      time.Time
	Exercises []ExerciseLog
	Comment   string
	UserId    string
}

type User struct {
	ID        string
	Name      string
	LastName  string
	Email     string
	Password  string
	Username  string
	Trainings []Training
	UserInfo  UserInfo
}

type UserInfo struct {
	StartWeight   float64
	Height        float64
	CurrentWeight float64
	Age           int
}

type UpdateExerciseRequest struct {
	Name    string `json:"name"`
	NewName string `json:"newName"`
	NewType string `json:"newType"`
}

type CreateExerciseRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateExerciseLogRequest struct {
	Exercises CreateExerciseRequest `json:"exercises"`
	Sets      []int
	Reps      []int
}
