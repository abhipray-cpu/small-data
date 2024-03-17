// db.go

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database connection details
type DB struct {
	Client *mongo.Client
	Ctx    context.Context
}

// Connect initializes the MongoDB connection
func Connect() (*DB, error) {
	// Set up client options
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the value of the MONGO_URI environment variable
	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect client
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return &DB{
		Client: client,
		Ctx:    ctx,
	}, nil
}

// Close closes the MongoDB connection
func (db *DB) Close() {
	db.Client.Disconnect(db.Ctx)
	fmt.Println("Disconnected from MongoDB.")
}
