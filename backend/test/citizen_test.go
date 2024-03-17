package test

import (
	"backend/models"
	"backend/repository"
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var citizen_id primitive.ObjectID

func Client1() *mongo.Client {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	URI := os.Getenv("MONGO_URI")
	TestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}
	log.Println("Connected to database successfully")
	err = TestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("Ping failed: %s", err.Error())
	}
	log.Println("Ping successful")
	return TestClient
}

func TestClientRepository(t *testing.T) {
	client := Client1()
	defer client.Disconnect(context.Background())
	collection := client.Database("Test").Collection("citizen_test")

	citizenRepo := repository.CitizenRepo{
		MongoCollection: collection,
	}

	// create citizens
	t.Run("Create Citizen", func(t *testing.T) {
		citizenDoc := models.Citizen{
			FirstName: "Mundal",
			LastName:  "Dhaiya",
			DOB:       "21-09-2000",
			Gender:    "Male",
			Address:   "Some random address",
			City:      "Faridabad",
			PinCode:   "263139",
			State:     "Uttarakhand",
		}
		result, err := citizenRepo.CreateCitizen(&citizenDoc)
		if err != nil {
			t.Fatal("Failed to create citizen", err.Error())
		}
		t.Log("Created citizen", result)
	})

	// find all citizens
	t.Run("Find All Citizens", func(t *testing.T) {
		result, err := citizenRepo.GetCitizens(1)

		if err != nil {
			t.Fatal("Failed to get all citizens", err.Error())
		}
		t.Log("Fetched all citizens", result)
		citizen_id = result[0].ID
	})

	// find citizen
	t.Run("Find a single Citizen", func(t *testing.T) {
		result, err := citizenRepo.GetCitizen(citizen_id)

		if err != nil {
			t.Fatal("Failed to fetch citizen", err.Error())
		}
		t.Log("Fetched user", result)
	})

	// update citizen
	t.Run("Update Citizen", func(t *testing.T) {
		updatedDoc := models.Citizen{
			FirstName: "Mundal",
			LastName:  "Chaurasiya",
			DOB:       time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC),
			Gender:    "Male",
			Address:   "Some random address",
			City:      "Faridabad",
			PinCode:   "263139",
			State:     "Uttarakhand",
		}

		result, err := citizenRepo.UpdateCitizen(citizen_id, updatedDoc)

		if err != nil {
			t.Fatal("Failed to update citizen", result)
		}
		t.Log("Updated citizen", result)
	})

	// filter by first name
	t.Run("Filter by first name", func(t *testing.T) {
		result, err := citizenRepo.FilteredCitizens("first_name", "Mundal", 1)
		if err != nil {
			log.Fatal("Failed to filter by first name", err.Error())
		}
		t.Log("Successfully fetched these documents", result)

	})
	// filter by last name
	t.Run("Filter by last name", func(t *testing.T) {
		result, err := citizenRepo.FilteredCitizens("last_name", "Chaurasiya", 1)
		if err != nil {
			log.Fatal("Failed to filter by last name", err.Error())
		}
		t.Log("Successfully fetched these documents", result)
	})

	// filter by gender
	t.Run("Filter by gender", func(t *testing.T) {
		result, err := citizenRepo.FilteredCitizens("gender", "Male", 1)
		if err != nil {
			log.Fatal("Failed to filter by gender", err.Error())
		}
		t.Log("Successfully fetched these documents", result)
	})

	// filter by Address
	t.Run("Filter by address", func(t *testing.T) {
		result, err := citizenRepo.FilteredCitizens("address", "Some random address", 1)
		if err != nil {
			log.Fatal("Failed to filter by address", err.Error())
		}
		t.Log("Successfully fetched these documents", result)
	})

	// filter by city
	t.Run("Filter by city", func(t *testing.T) {
		result, err := citizenRepo.FilteredCitizens("City", "Faridabad", 1)
		if err != nil {
			log.Fatal("Failed to filter by city", err.Error())
		}
		t.Log("Successfully fetched these documents", result)
	})

	// filter by pincode
	t.Run("Filter by pincode", func(t *testing.T) {
		result, err := citizenRepo.FilteredCitizens("pin_code", "263139", 1)
		if err != nil {
			log.Fatal("Failed to filter by pincode", err.Error())
		}
		t.Log("Successfully fetched these documents", result)
	})

	// filter by state
	t.Run("Filter by state", func(t *testing.T) {
		result, err := citizenRepo.FilteredCitizens("state", "Uttarakhand", 1)
		if err != nil {
			log.Fatal("Failed to filter by state", err.Error())
		}
		t.Log("Successfully fetched these documents", result)
	})

	// complex query
	t.Run("Complex query", func(t *testing.T) {
		query := map[string]string{}
		query["first_name"] = "Mundal"
		query["gender"] = "Male"
		result, err := citizenRepo.ComplexQuery(1, query)

		if err != nil {
			t.Fatal("Failed to fetch document", err.Error())
		}
		t.Log("Fetched document", result)

	})
	// delete citizen
	t.Run("Delete a citizen", func(t *testing.T) {
		result, err := citizenRepo.DeleteCitizen(citizen_id)

		if err != nil {
			t.Fatal("Failed to delete citizen", err.Error())
		}

		t.Log("Deleted citizen", result)
	})

	// delete citizens
	t.Run("Delete all citizens", func(t *testing.T) {
		result, err := citizenRepo.DeleteAll()
		if err != nil {
			t.Fatal("Failed to delete all citizens", err.Error())
		}
		t.Log("Deleted all citizens", result)
	})
}
