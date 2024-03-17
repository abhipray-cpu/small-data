// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(":8080", router))
}
