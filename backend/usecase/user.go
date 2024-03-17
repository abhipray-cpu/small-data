package usecase

import (
	"backend/models"
	"backend/repository"
	authentication "backend/utils"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *UserService) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var login struct {
		Email    string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}
	// check if user exists
	userRepo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	result, err := userRepo.GetMail(login.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = "Something went wrong at our end"
		log.Println("Failed to fetch user by email", err.Error())
		return
	}
	if len(result) == 0 {
		w.WriteHeader(http.StatusNotFound)
		res.Error = "No such user exists"
		return
	}
	user := result[0]
	// password matches
	if authentication.VerifyPassword(user.Password, login.Password) {
		// generate jwt token
		token, err := authentication.GenerateToken(user.ID.Hex())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res.Error = "Something went wrong at our end"
			log.Println("Failed to generate token", err.Error())
			return
		}
		// send token
		w.WriteHeader(http.StatusOK)
		res.Data = token
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	res.Error = "Wrong Password"
	return

}

func (svc *UserService) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}
	// can also add validation here

	// like check for existing user let's add it
	if svc.existingUser(user.Email) {
		w.WriteHeader(http.StatusAlreadyReported)
		log.Println(user.Email, "already exists")
		res.Error = "email already exists"
		return
	}
	userRepo := repository.UserRepo{MongoCollection: svc.MongoCollection}

	hashedPassword, err := authentication.HashPassword(user.Password)
	if err != nil {
		log.Println("Failed to hash password")
		w.WriteHeader(http.StatusInternalServerError)
		res.Error = "Something went wrong at our end"
		return
	}
	user.Password = hashedPassword
	insertID, err := userRepo.CreateUser(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Insert error", err)
		res.Error = err.Error()
		return
	}
	res.Data = user
	w.WriteHeader(http.StatusOK)
	log.Println("Citizen created with id", insertID, user)

}

func (svc *UserService) existingUser(email string) bool {
	userRepo := repository.UserRepo{MongoCollection: svc.MongoCollection}
	result, err := userRepo.GetMail(email)

	if err != nil {
		return false
	}
	return len(result) > 0
}
