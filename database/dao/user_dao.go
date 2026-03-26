package dao

import (
	"context"
	"crypto/sha256"
	"gym-log/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserDao struct {
	dao *mongo.Collection
}

func NewUserDao(client *mongo.Client, dbName string, collection string) *UserDao {
	return &UserDao{
		dao: client.Database(dbName).Collection(collection),
	}
}

func (u *UserDao) Register(ctx context.Context, username string, password string, name string, lastName string, email string) (any, error) {
	hashPassword := sha256.Sum256([]byte(password))
	doc := bson.M{
		"username": username,
		"password": string(hashPassword[:]),
		"name":     name,
		"lastName": lastName,
		"email":    email,
	}
	result, err := u.dao.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (u *UserDao) Login(ctx context.Context, username string, password string) (*models.User, error) {
	hashPassword := sha256.Sum256([]byte(password))
	filter := bson.M{
		"username": username,
		"password": string(hashPassword[:]),
	}
	result := u.dao.FindOne(ctx, filter)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	var user models.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) AddTraining(ctx context.Context, training models.Training, userId string) (int64, error) {
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return -1, err
	}
	filter := bson.M{
		"_id": objectID,
	}
	doc := bson.M{
		"$push": bson.M{
			"trainings": training,
		},
	}
	result, err := u.dao.UpdateOne(ctx, filter, doc)
	if err != nil {
		return -1, err
	}

	return result.ModifiedCount, nil
}
