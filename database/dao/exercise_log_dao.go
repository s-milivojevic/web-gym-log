package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)
import "gym-log/models"

type ExerciseLogDao struct {
	collection *mongo.Collection
}

func NewExerciseLogDao(db *mongo.Client, dbName string, collectionName string) *ExerciseLogDao {
	return &ExerciseLogDao{
		collection: db.Database(dbName).Collection(collectionName),
	}
}

func (e *ExerciseLogDao) AddExerciseLog(ctx context.Context, exerciseLogRequest models.CreateExerciseLogRequest) (error, *mongo.InsertOneResult) {
	doc := bson.M{
		"exercises": exerciseLogRequest.Exercises,
		"sets":      exerciseLogRequest.Sets,
		"reps":      exerciseLogRequest.Reps,
	}
	result, err := e.collection.InsertOne(ctx, doc)
	if err != nil {
		return err, nil
	}
	if result.InsertedID == nil {
		return nil, nil
	} else {
		return nil, result
	}

}

func (e *ExerciseLogDao) GetAllExerciseLogs(ctx context.Context) ([]models.ExerciseLog, error) {

}
