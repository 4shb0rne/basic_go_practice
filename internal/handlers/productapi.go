package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/4shb0rne/goapi-basic/internal/models"
	"github.com/4shb0rne/goapi-basic/internal/tools"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	// Create an instance of MySQLHandler
	mysqlHandler := tools.NewMySQLHandler("root", "", "go_learn")
	// Get a database connection
	db, err := mysqlHandler.GetDatabase()
	if err != nil {
		log.Errorf("Error connecting to the database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Retrieve all products from the database
	result, err := mysqlHandler.ViewAll(db)
	if err != nil {
		log.Errorf("Error retrieving products: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Convert the result to JSON
	productsJSON, err := json.Marshal(result)
	if err != nil {
		log.Errorf("Error marshalling products to JSON: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(productsJSON)
	if err != nil {
		log.Errorf("Error writing JSON response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request body into Product structs
	var products []models.Product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		log.Errorf("Error decoding JSON: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create an instance of MySQLHandler
	mysqlHandler := tools.NewMySQLHandler("root", "", "go_learn")

	// Insert products into the database
	err = mysqlHandler.Insert(products)
	if err != nil {
		log.Errorf("Error inserting products: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Products inserted successfully.")
}
