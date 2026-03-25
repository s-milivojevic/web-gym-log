package dao

import (
	"context"
	"gym-log/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TrainingDao struct {
	dao *mongo.Collection
}

func NewTrainingDao(client *mongo.Client, dbName string, collectionName string) *TrainingDao {
	return &TrainingDao{dao: client.Database(dbName).Collection(collectionName)}
}

func (t *TrainingDao) CreateTraining(ctx context.Context, userId string) (any, error) {
	object_id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return bson.ObjectID{}, err
	}
	object_id = primitive.ObjectID(object_id)
	doc := bson.M{
		"user_id": object_id,
		"date":    time.Now(),
	}
	result, err := t.dao.InsertOne(ctx, doc)
	if err != nil {
		return bson.ObjectID{}, err
	}
	return result.InsertedID, nil
}

func (t *TrainingDao) AddExerciseLog(ctx context.Context, exerciseLog models.ExerciseLog, trainingId string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(trainingId)
	if err != nil {
		return -1, err
	}

	filter := bson.M{
		"_id": objectID,
	}

	update := bson.M{
		"$push": bson.M{
			"exercise_log": exerciseLog,
		},
	}

	result, err := t.dao.UpdateOne(ctx, filter, update)
	if err != nil {
		return -1, err
	}

	return result.ModifiedCount, nil
}

func (t *TrainingDao) GetUsersTrainings(ctx context.Context, userId string) ([]models.Training, error) {
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"_id": objectID,
	}
	var trainings []models.Training
	cursor, err := t.dao.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var training models.Training
		if err := cursor.Decode(&trainings); err != nil {
			return nil, err
		}
		trainings = append(trainings, training)
	}
	return trainings, nil
}

func (t *TrainingDao) UpdateTrainingDuration(ctx context.Context, trainingId string, duration time.Duration) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(trainingId)
	if err != nil {
		return -1, err
	}
	filter := bson.M{
		"_id": objectID,
	}
	doc := bson.M{
		"$update": bson.M{
			"duration": duration,
		},
	}
	result, err := t.dao.UpdateOne(ctx, filter, doc)
	if err != nil {
		return -1, err
	}
	return result.ModifiedCount, nil
}

func (t *TrainingDao) UpdateTrainingDate(ctx context.Context, trainingId string, date time.Time) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(trainingId)
	if err != nil {
		return -1, err
	}
	filter := bson.M{
		"_id": objectID,
	}
	doc := bson.M{
		"$update": bson.M{
			"date": date,
		},
	}
	result, err := t.dao.UpdateOne(ctx, filter, doc)
	if err != nil {
		return -1, err
	}
	return result.ModifiedCount, nil
}

func (t *TrainingDao) UpdateTrainingComment(ctx context.Context, trainingId string, comment string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(trainingId)
	if err != nil {
		return -1, err
	}
	filter := bson.M{
		"_id": objectID,
	}
	doc := bson.M{
		"$update": bson.M{
			"comment": comment,
		},
	}
	result, err := t.dao.UpdateOne(ctx, filter, doc)
	if err != nil {
		return -1, err
	}
	return result.ModifiedCount, nil
}

func (t *TrainingDao) UpdateTrainingExerciseLog(ctx context.Context, trainingId string, exerciseLog []models.ExerciseLog) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(trainingId)
	if err != nil {
		return -1, err
	}
	filter := bson.M{
		"_id": objectID,
	}
	update := bson.M{
		"$update": bson.M{
			"exercise_log": exerciseLog,
		},
	}
	result, err := t.dao.UpdateOne(ctx, filter, update)
	if err != nil {
		return -1, err
	}
	return result.ModifiedCount, nil
}
