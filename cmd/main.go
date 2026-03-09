package main

import (
	"log"
	"net/http"
	"product-api/internal/service"

	"os"
	"product-api/internal/database"

	"github.com/gorilla/mux"
)

func main() {
	shouldSeed := false
	for _, arg := range os.Args[1:] {
			if arg == "--seed" {
					shouldSeed = true
			}
	}
	if shouldSeed {
			database.Store.SeedProducts(50)
	}

	router := mux.NewRouter()

	router.HandleFunc("/products", service.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", service.GetProduct).Methods("GET")
	router.HandleFunc("/products", service.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", service.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", service.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/{id}", service.PatchProductStatus).Methods("PATCH")

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
