package main

import (
	"fmt"
	"net/http"

	"github.com/4shb0rne/goapi-basic/internal/handlers"
	"github.com/4shb0rne/goapi-basic/internal/models"
	"github.com/4shb0rne/goapi-basic/internal/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	mysqlHandler := tools.NewMySQLHandler("root", "", "go_learn")

	// Example usage
	db, er := mysqlHandler.GetDatabase()
	if er != nil {
		log.Fatalf("Error connecting to the database: %v", er)
	}

	result, er := mysqlHandler.ViewAll(db)
	if er != nil {
		log.Errorf("Error retrieving products: %v", er)
		// Handle the error appropriately
	} else {
		// Assert the type of result to []models.Product
		products, ok := result.([]models.Product)
		if !ok {
			// Handle the case where the type assertion failed
			log.Error("Failed to assert type of products")
			// Handle the error appropriately
		} else {
			// Access the products slice
			for _, product := range products {
				// Do something with each product
				fmt.Printf("ProductID: %d, ProductName: %s, ProductPrice: %d, ProductStock: %d\n", product.ProductID, product.ProductName, product.ProductPrice, product.ProductStock)
			}
		}
	}
	defer db.Close()

	fmt.Println("Starting GO API Services")

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
