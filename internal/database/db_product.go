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

func (ps *ProductStore) SeedProducts(n int) {
	for i := 0; i < n; i++ {
		ps.Products[ps.NextID] = Product{
			ID:          ps.NextID,
			Name:        "Product " + string(rune(65+i)),
			Description: "Sample product description",
			IsActive:    true,
		}
		ps.NextID++
	}
}
