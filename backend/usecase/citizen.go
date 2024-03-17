package usecase

import (
	"backend/models"
	"backend/repository"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CitizenService struct {
	MongoCollection *mongo.Collection
}

func (svc *CitizenService) CreateCitizen(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var citizen models.Citizen
	err := json.NewDecoder(r.Body).Decode(&citizen)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}

	citizenRepo := repository.CitizenRepo{MongoCollection: svc.MongoCollection}
	insertID, err := citizenRepo.CreateCitizen(&citizen)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Insert error", err)
		res.Error = err.Error()
		return
	}
	res.Data = citizen
	w.WriteHeader(http.StatusOK)
	log.Println("Citizen created with id", insertID, citizen)

}

func (svc *CitizenService) UpdateCitizen(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	citizenID := mux.Vars(r)["id"]
	objectID, err := primitive.ObjectIDFromHex(citizenID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to initialize object id", err)
		res.Error = err.Error()
		return
	}

	var updatedCitizen models.Citizen
	err = json.NewDecoder(r.Body).Decode(&updatedCitizen)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to parse the body", err.Error())
		res.Error = err.Error()
		return
	}
	citizenRepo := repository.CitizenRepo{MongoCollection: svc.MongoCollection}

	count, err := citizenRepo.UpdateCitizen(objectID, updatedCitizen)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to update citizen", err.Error())
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)

}

func (svc *CitizenService) DeleteCitizen(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	citizenID := mux.Vars(r)["id"]
	citizenRepo := repository.CitizenRepo{MongoCollection: svc.MongoCollection}
	objectID, err := primitive.ObjectIDFromHex(citizenID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to initialize object id", err)
		res.Error = err.Error()
		return
	}

	citizen, err := citizenRepo.DeleteCitizen(objectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to delete citizen", err)
		res.Error = err.Error()
		return
	}
	res.Data = citizen
	w.WriteHeader(http.StatusOK)
}

func (svc *CitizenService) FindCitizens(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	pageNumber, err := strconv.Atoi(mux.Vars(r)["pageNumber"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Unable to parse the page number", err)
		res.Error = err.Error()
		return
	}
	var filters map[string]string
	err = json.NewDecoder(r.Body).Decode(&filters)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Unable to parse the request body for filters", err)
		res.Error = err.Error()
		return
	}
	citizenRepo := repository.CitizenRepo{MongoCollection: svc.MongoCollection}
	result, err := citizenRepo.ComplexQuery(pageNumber, filters)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to fetch documents", err)
		res.Error = err.Error()
		return
	}

	res.Data = result
	w.WriteHeader(http.StatusOK)

}

func (svc *CitizenService) FindCitizen(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	citizenID := mux.Vars(r)["id"]
	citizenRepo := repository.CitizenRepo{MongoCollection: svc.MongoCollection}
	objectID, err := primitive.ObjectIDFromHex(citizenID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to initialize object id", err)
		res.Error = err.Error()
		return
	}

	citizen, err := citizenRepo.GetCitizen(objectID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to get citizen", err)
		res.Error = err.Error()
		return
	}
	res.Data = citizen
	w.WriteHeader(http.StatusOK)
}

func (svc *CitizenService) GetCitizens(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	pageNumber, err := strconv.Atoi(mux.Vars(r)["pageNumber"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Unable to parse the page number", err)
		res.Error = err.Error()
		return
	}
	citizenRepo := repository.CitizenRepo{MongoCollection: svc.MongoCollection}

	citizen, err := citizenRepo.GetCitizens(pageNumber)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Failed to get citizen", err)
		res.Error = err.Error()
		return
	}
	res.Data = citizen
	w.WriteHeader(http.StatusOK)
}
