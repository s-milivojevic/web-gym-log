package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (e *ExerciseLogDao) DeleteExerciseLog(ctx context.Context, id bson.ObjectID) (int64, error) {
	result, err := e.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return -1, err
	}
	return result.DeletedCount, nil
}

func (e *ExerciseLogDao) UpdateExerciseLog(ctx context.Context, exerciseLogRequest models.UpdateExerciseLogRequest) (int64, error) {
	id := exerciseLogRequest.ID
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return -1, err
	}
	object_id = primitive.ObjectID(bson.ObjectID(object_id))
	doc := bson.M{
		"exercises": exerciseLogRequest.Exercises,
		"sets":      exerciseLogRequest.Sets,
		"reps":      exerciseLogRequest.Reps,
	}
	result, err := e.collection.UpdateOne(ctx, bson.M{"_id": id}, doc)
	if err != nil {
		return -1, err
	}
	return result.ModifiedCount, nil
}
