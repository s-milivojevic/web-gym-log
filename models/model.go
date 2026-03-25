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
	Exercises Exercise      `bson:"exercises" json:"exercises"`
	Reps      []int         `bson:"reps" json:"reps"`
	Sets      []int         `bson:"sets" json:"sets"`
}

type Training struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	Date      time.Time     `bson:"date" json:"date"`
	Exercises []ExerciseLog `bson:"exercises" json:"exercises"`
	Comment   string        `bson:"comment" json:"comment"`
	UserId    bson.ObjectID `bson:"user_id" json:"user_id"`
	Duration  time.Duration `bson:"duration" json:"duration"`
}

type User struct {
	ID        string
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

type UpdateExerciseLogRequest struct {
	ID        string                `bson:"_id" json:"id"`
	Exercises UpdateExerciseRequest `json:"exercises"`
	Sets      []int
	Reps      []int
}
