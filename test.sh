#!/bin/bash

echo 'GET /products'
curl http://localhost:8080/products
echo

echo 'POST /products'
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "Low-performance laptop",
    "isActive": true
  }'
echo

echo 'POST /products'
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop"
  }'
echo

echo 'GET /products'
curl http://localhost:8080/products
echo

echo 'PUT /products/2'
curl -X PUT http://localhost:8080/products/2 \
-H "Content-Type: application/json" \
-d '{
  "name": "Gaming Laptop",
  "description": "Ultra high-performance gaming laptop",
  "isActive": false
}'
echo

echo 'GET /products'
curl http://localhost:8080/products
echo

echo 'PATCH /products/2'
curl -X PATCH http://localhost:8080/products/2 \
-H "Content-Type: application/json" \
-d '{
  "isActive": true
}'
echo

echo 'GET /products'
curl http://localhost:8080/products
echo

echo 'DELETE /products/1'
curl -X DELETE http://localhost:8080/products/1
echo

echo 'GET /products/1'
curl http://localhost:8080/products/1
echo

echo 'GET /products'
curl http://localhost:8080/products
echo