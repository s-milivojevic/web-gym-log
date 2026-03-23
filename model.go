package main

import "time"

type Exercise struct {
	ID   string
	Name string
	Type string
}

type ExerciseLog struct {
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
