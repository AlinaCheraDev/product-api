package database

// ProductStore simulates a database for products

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	IsActive    bool    `json:"isActive"`
}

type ProductStore struct {
	Products map[int]Product
	NextID   int
}

var Store = ProductStore{
	Products: make(map[int]Product),
	NextID:   1,
}
