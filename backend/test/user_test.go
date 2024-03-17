package test

import (
	"backend/models"
	"backend/repository"
	authentication "backend/utils"
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var user1_id primitive.ObjectID
var user2_id primitive.ObjectID

func newMongoClient() *mongo.Client {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the value of the MONGO_URI environment variable
	mongoURI := os.Getenv("MONGO_URI")
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}
	log.Println("Connected to database successfully")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("Ping failed: %s", err.Error())
	}
	log.Println("Ping successful")

	// creating unique index no email
	collection := mongoTestClient.Database("Test").Collection("user_test")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "user_mail", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Error creating unique index: %v", err)
	}
	log.Println("Unique index on 'email' field created successfully")
	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	collection := mongoTestClient.Database("Test").Collection("user_test")

	userRepo := repository.UserRepo{
		MongoCollection: collection,
	}
	// Create employee
	t.Run("Create users", func(t *testing.T) {
		hashedPassword, err := authentication.HashPassword("password")
		user1 := models.User{
			Name:     "Abhipray Dumka",
			Email:    "dumkaabhipray@gmail.com",
			Type:     "Admin",
			Password: hashedPassword,
		}
		user2 := models.User{
			Name:     "Puttanpal",
			Email:    "abhipraydumka@gmail.com",
			Type:     "Employee",
			Password: hashedPassword,
		}

		result, err := userRepo.CreateUser(&user1)
		if err != nil {
			t.Fatal("Insert 1 operation failed", err.Error())
		}
		t.Log("Insert 1 succeeded", result)

		result, err = userRepo.CreateUser(&user2)
		if err != nil {
			t.Fatal("Insert 2 operation failed", err.Error())
		}
		t.Log("Insert 2 succeeded", result)
	})
	// find employees
	t.Run("Fetch users", func(t *testing.T) {
		result, err := userRepo.GetUsers()
		if err != nil {
			t.Fatal("Failed to fetch users", err.Error())
		}
		t.Log("Fetched all users", result)
		// You need to perform type assertion to access the fields of the struct
		user1_id = result[0].ID
		user2_id = result[1].ID
	})

	// find employee
	t.Run("Fetch User", func(t *testing.T) {
		result, err := userRepo.GetUser(user1_id)
		if err != nil {
			t.Fatal("Failed to get user 1", err.Error())
		}
		t.Log("Fetched user", result)
	})
	// update employee
	t.Run("Update user 1", func(t *testing.T) {
		updatedDoc := models.User{
			Name:     "Abhipray Dumka",
			Email:    "dumkaabhipray@gmail.com",
			Type:     "Admin",
			Password: "password_updated",
		}
		result, err := userRepo.UpdateUser(user1_id, updatedDoc)
		if err != nil {
			t.Fatal("Failed to update user 1", err.Error())
		}
		t.Log("Updated user1", result)
	})
	// delete employee
	t.Run("Delete user 1", func(t *testing.T) {
		result, err := userRepo.DeleteUser(user1_id)

		if err != nil {
			t.Fatal("Failed to delete user 1", err.Error())
		}
		t.Log("Deleted user 1", result)
	})
	t.Run("Find by mail", func(t *testing.T) {
		result, err := userRepo.GetMail("dumkaabhipray@gmail.com")
		if err != nil {
			t.Fatal("Failed to get doc by mail", err.Error())
		}
		t.Log("Fetched doc by mail", result)
	})
	// delete all
	t.Run("Delete all employees", func(t *testing.T) {
		result, err := userRepo.DeleteAll()
		if err != nil {
			t.Fatal("Failed to delete all users", err.Error())
		}
		t.Log("Delete all users:", result)
	})
}
