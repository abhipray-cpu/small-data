package main

import (
	"backend/usecase"
	authentication "backend/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(db *mongo.Client) http.Handler {

	// accessing both collections
	userCollection := db.Database("smalldata").Collection("user")
	citizenCollection := db.Database("smalldata").Collection("citizen")

	userService := usecase.UserService{MongoCollection: userCollection}
	citizenService := usecase.CitizenService{MongoCollection: citizenCollection}

	router := mux.NewRouter()
	router.HandleFunc("/health", HealthHandler).Methods(http.MethodGet)

	// user routes
	router.HandleFunc("/login", userService.Login).Methods(http.MethodPost)
	router.HandleFunc("/signup", userService.Signup).Methods(http.MethodPost)

	// citizen routes
	router.Handle("/citizens/{pageNumber}", authentication.Authenticate(http.HandlerFunc(citizenService.GetCitizens))).Methods(http.MethodGet)
	router.Handle("/createCitizen", authentication.Authenticate(http.HandlerFunc(citizenService.CreateCitizen))).Methods(http.MethodPost)
	router.Handle("/updateCitizen/{id}", authentication.Authenticate(http.HandlerFunc(citizenService.UpdateCitizen))).Methods(http.MethodPut)
	router.Handle("/deleteCitizen/{id}", authentication.Authenticate(http.HandlerFunc(citizenService.DeleteCitizen))).Methods(http.MethodDelete)
	router.Handle("/citizen/{id}", authentication.Authenticate(http.HandlerFunc(citizenService.FindCitizen))).Methods(http.MethodGet)
	router.Handle("/findCitizens/{pageNumber}", authentication.Authenticate(http.HandlerFunc(citizenService.FindCitizens))).Methods(http.MethodPost)
	return router
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running...."))
}
