package repository

import (
	"backend/models"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CitizenRepo struct {
	MongoCollection *mongo.Collection
}

// CreateCitizen creates a new entry
func (c *CitizenRepo) CreateCitizen(citizen *models.Citizen) (interface{}, error) {
	validate := validator.New()
	if err := validate.Struct(citizen); err != nil {
		return nil, err
	}
	result, err := c.MongoCollection.InsertOne(context.Background(), citizen)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// updates all the matched documents
func (c *CitizenRepo) UpdateCitizen(citizenId primitive.ObjectID, updatedDoc models.Citizen) (int64, error) {
	result, err := c.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "_id", Value: citizenId}},
		bson.D{{Key: "$set", Value: updatedDoc}})

	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

// delete all matched documents
func (c *CitizenRepo) DeleteCitizen(citizenId primitive.ObjectID) (int64, error) {
	result, err := c.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "_id", Value: citizenId}})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

// delete all matched entries in the database
func (c *CitizenRepo) DeleteAll() (int64, error) {
	result, err := c.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

// fetches a singe matched citizen
func (c *CitizenRepo) GetCitizen(citizenId primitive.ObjectID) (*models.Citizen, error) {
	var citizen models.Citizen

	err := c.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "_id", Value: citizenId}}).Decode(&citizen)
	if err != nil {
		return nil, err
	}
	return &citizen, nil
}

// filtered results
func (c *CitizenRepo) FilteredCitizens(key string, value string, page int) ([]models.Citizen, error) {
	pageSize := 10
	skip := (page - 1) * pageSize
	options := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(skip))
	result, err := c.MongoCollection.Find(context.Background(),
		bson.D{{Key: key, Value: value}}, options)
	if err != nil {
		return nil, err
	}
	var citizens []models.Citizen
	err = result.All(context.Background(), &citizens)
	if err != nil {
		return nil, fmt.Errorf("results decode error:%s", err.Error())
	}
	return citizens, nil
}

// fetch all citizens
func (c *CitizenRepo) GetCitizens(page int) ([]models.Citizen, error) {
	pageSize := 10
	skip := (page - 1) * pageSize
	options := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(skip))
	result, err := c.MongoCollection.Find(context.Background(),
		bson.D{}, options)
	if err != nil {
		return nil, err
	}
	var citizens []models.Citizen
	err = result.All(context.Background(), &citizens)
	if err != nil {
		return nil, fmt.Errorf("results decode error:%s", err.Error())
	}
	return citizens, nil
}

func (c *CitizenRepo) ComplexQuery(page int, filters map[string]string) ([]models.Citizen, error) {
	pageSize := 10
	skip := (page - 1) * pageSize
	options := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(skip))
	filter := bson.D{}
	for key, value := range filters {
		filter = append(filter, bson.E{Key: key, Value: value})
	}
	result, err := c.MongoCollection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	var citizens []models.Citizen
	err = result.All(context.Background(), &citizens)
	if err != nil {
		return nil, fmt.Errorf("results decode error: %s", err.Error())
	}
	return citizens, nil
}
