// main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Connect to MongoDB
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// creating a new server
	router := Router(db.Client)
	fmt.Println("Server is listening on port 8080...")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
