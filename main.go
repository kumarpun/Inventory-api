package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("sqlite3", "../inventory.db")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Device{})

	// Handle Subsequent requests

	fmt.Println("Api running on port 4000...")

	r := mux.NewRouter().StrictSlash(true)

	headers := handlers.AllowedHeaders([]string{"X-Request", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// For Device table
	r.HandleFunc("/devices", getDevices).Methods("GET")

	r.HandleFunc("/devices/{id}", getDevice).Methods("GET")

	r.HandleFunc("/devices", postDevice).Methods("POST")

	r.HandleFunc("/devices/{id}", putDevice).Methods("PUT")

	r.HandleFunc("/devices/{id}", deleteDevice).Methods("DELETE")

	// For List table

	r.HandleFunc("/devices/{id}", approveDevice).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":4000", handlers.CORS(headers, methods, origins)(r)))

}
