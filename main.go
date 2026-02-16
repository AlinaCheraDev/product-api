package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	IsActive		bool    `json:"isActive"`
}

// Store products in memory
type ProductStore struct {
	products map[int]Product
	nextID   int
}

var store = ProductStore{
	products: make(map[int]Product),
	nextID:   1,
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	router.HandleFunc("/products/{id}", PatchProductStatus).Methods("PATCH")

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GET /products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := make([]Product, 0, len(store.products))
	for _, product := range store.products {
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GET /products/{id}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, exists := store.products[id]

	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// POST /products
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product.ID = store.nextID
	store.nextID++
	store.products[product.ID] = product

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// PUT /products/{id}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, exists := store.products[id]; !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	product.ID = id
	store.products[id] = product

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// DELETE /products/{id}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if _, exists := store.products[id]; !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	delete(store.products, id)
	w.WriteHeader(http.StatusNoContent)
}

// PATCH /products/{id}
func PatchProductStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var statusUpdate struct {
		IsActive bool `json:"isActive"`
	}
	if err := json.NewDecoder(r.Body).Decode(&statusUpdate); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, exists := store.products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	product.IsActive = statusUpdate.IsActive
	store.products[id] = product

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}