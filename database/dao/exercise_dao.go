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

func (e *ExerciseDao) GetExercises(ctx context.Context) ([]models.Exercise, error) {
	var exercises []models.Exercise

	cursor, err := e.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var exercise models.Exercise
		if err := cursor.Decode(&exercise); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return exercises, nil
}

func (e *ExerciseDao) UpdateExercise(ctx context.Context, req models.UpdateExerciseRequest) (models.Exercise, error) {
	var existing models.Exercise

	err := e.collection.FindOne(ctx, bson.M{"name": req.Name}).Decode(&existing)
	if err != nil {
		return models.Exercise{}, err
	}

	updateFields := bson.M{}

	if req.NewName != "" {
		updateFields["name"] = req.NewName
	}

	if req.NewType != "" {
		updateFields["type"] = req.NewType
	}

	if len(updateFields) == 0 {
		return existing, nil
	}

	_, err = e.collection.UpdateOne(
		ctx,
		bson.M{"name": req.Name},
		bson.M{"$set": updateFields},
	)
	if err != nil {
		return models.Exercise{}, err
	}

	err = e.collection.FindOne(ctx, bson.M{"name": req.Name}).Decode(&existing)
	if err != nil {
		return models.Exercise{}, err
	}

	return existing, nil
}

func (e *ExerciseDao) DeleteExercise(ctx context.Context, name string) (int64, error) {
	deleted_exercise, err := e.collection.DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		return -1, err
	} else {
		return deleted_exercise.DeletedCount, nil
	}
}
