package repository

import (
	"backend/models"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	MongoCollection *mongo.Collection
}

// createUser creates a new user in the database
func (u *UserRepo) CreateUser(user *models.User) (interface{}, error) {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return nil, err
	}
	result, err := u.MongoCollection.InsertOne(context.Background(), user)

	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// updateUser updates an all matched documents
func (u *UserRepo) UpdateUser(userID primitive.ObjectID, updatedDoc models.User) (int64, error) {
	result, err := u.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "_id", Value: userID}},
		bson.D{{Key: "$set", Value: updatedDoc}})
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil

}

// GetUser finds a single user
func (u *UserRepo) GetUser(user_id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := u.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "_id", Value: user_id}}).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers fetches all the users
func (u *UserRepo) GetUsers() ([]models.User, error) {
	result, err := u.MongoCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = result.All(context.Background(), &users)
	if err != nil {
		return nil, fmt.Errorf("results decode error:%s", err.Error())
	}
	return users, nil
}

// DeleteUser deletes all matched users
func (u *UserRepo) DeleteUser(userId primitive.ObjectID) (int64, error) {
	result, err := u.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "_id", Value: userId}})
	if err != nil {
		return 0, nil
	}
	return result.DeletedCount, nil
}

// DeleteAll deletes all the users
func (u *UserRepo) DeleteAll() (int64, error) {
	result, err := u.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, nil
	}
	return result.DeletedCount, nil
}

func (u *UserRepo) GetMail(mail string) ([]models.User, error) {
	result, err := u.MongoCollection.Find(context.Background(), bson.D{{
		Key: "user_mail", Value: mail,
	}})
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = result.All(context.Background(), &users)
	if err != nil {
		return nil, fmt.Errorf("results decode error:%s", err.Error())
	}
	return users, nil
}
