package dao

import (
	"context"
	"gym-log/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ExerciseDao struct {
	collection *mongo.Collection
}

func NewExerciseDao(client *mongo.Client, dbName string, collectionName string) *ExerciseDao {
	return &ExerciseDao{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (e *ExerciseDao) AddExercise(ctx context.Context, exercise models.Exercise) error {
	doc := bson.M{
		"name": exercise.Name,
		"type": exercise.Type,
	}

	_, err := e.collection.InsertOne(ctx, doc)
	return err
}
