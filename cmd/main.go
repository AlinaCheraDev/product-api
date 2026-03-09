package main

import (
	"log"
	"net/http"
	"product-api/internal/service"

	"github.com/gorilla/mux"
)

func main() {
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
