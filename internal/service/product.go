package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"product-api/internal/database"

	"github.com/gorilla/mux"
)

// GET /products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := make([]database.Product, 0, len(database.Store.Products))
	for _, product := range database.Store.Products {
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

	product, exists := database.Store.Products[id]

	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// POST /products
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product database.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product.ID = database.Store.NextID
	database.Store.NextID++
	database.Store.Products[product.ID] = product

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

	var product database.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if _, exists := database.Store.Products[id]; !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	product.ID = id
	database.Store.Products[id] = product

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

	if _, exists := database.Store.Products[id]; !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	delete(database.Store.Products, id)
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

	product, exists := database.Store.Products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	product.IsActive = statusUpdate.IsActive
	database.Store.Products[id] = product

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}