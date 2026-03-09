# Product REST API in Go

A simple REST API for managing products.

## Setup

1. Install dependencies:

```bash
go mod tidy
```

2. Run the server:

```bash
go run cmd/main.go
```

or, if I want to database seeded:

```bash
go run cmd/main.go --seed
```

The server will start on `http://localhost:8080`

3. Test the API:

- simple tests:

```bash
./test.sh
```

- performance tests (requires k6 to be installed):

```bash
k6 run internal/test/k6_products.js
```

## API Endpoints

### Get all products

```bash
GET /products
```

Example:

```bash
curl http://localhost:8080/products
```

### Get a single product

```bash
GET /products/{id}
```

Example:

```bash
curl http://localhost:8080/products/1
```

### Create a product

```bash
POST /products
Content-Type: application/json
```

Example:

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "isActive": true
  }'
```

### Update a product

```bash
PUT /products/{id}
Content-Type: application/json
```

Example:

```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Gaming Laptop",
    "description": "Ultra high-performance gaming laptop",
    "isActive": true
  }'
```

### Delete a product

```bash
DELETE /products/{id}
```

Example:

```bash
curl -X DELETE http://localhost:8080/products/1
```

### Change product active status

```bash
PATCH /products/{id}
Content-Type: application/json
```

Example:

```bash
curl -X PATCH http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "isActive": true
  }'
```

## Key Features

- **In-memory storage**: Uses a map to store products (data resets on restart)
- **Thread-safe**: Uses mutexes to handle concurrent requests
- **RESTful**: Follows REST conventions for HTTP methods and status codes
- **JSON responses**: All data exchanged in JSON format
