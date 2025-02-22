# Cimri Internship Case - Products Service
This repository contains the Favorites Service, which is a part of a microservices-based architecture developed in Go. The system consists of the following core services:

User Service: Handles user-related operations, including managing user profiles, preferences, and interactions.

Product Service: Manages product data and retrieves product details based on product ID.

Favorites Service: Enables users to create, update, delete, and view their favorite lists. Users can also add or remove products from these lists. This service communicates with both the User Service and the Product Service.

The service consists of the following key components:

- **Handlers**: HTTP endpoints that handle incoming requests for product data.
- **Models**: Defines the structure of the product data.
- **Repositories**: Interacts with the database to fetch product data.
- **Services**: Contains the business logic for handling product-related operations.

## 📌 Technologies Used

- **Go (Golang)**
- **Fiber** (Web Framework)
- **GORM** (Object Relational Mapping)
- **PostgreSQL**
- **Swagger / OpenAPI**

## 📂 Project Structure

```
/cmd
  /main.go         # Entry point for the HTTP server
/internal
  /handler         # Handles HTTP requests related to products
  /service         # Contains business logic for managing products
  /repository      # Handles database interactions (with GORM)
  /models          # Defines the structure of product data
/utils
  /envloader       # Loads environment variables
/pkg
  /postgres        # Handles PostgreSQL connection (with GORM)
```

### **Product Handlers**:

The **Product Handler** exposes the following key routes for interacting with product data:

- **`/products`**: This route retrieves all products from the database and returns them in JSON format.
- **`/products/{id}`**: This route retrieves a product by its unique ID.

In addition to these routes, there is also a random error generation mechanism that can simulate **503 Service Unavailable** errors for testing purposes. This is implemented in the function `generateRandom503Error`.

### **Product Service**:

- The **Product Service** acts as the business layer for handling product-related operations.
- It retrieves product data from the **Product Repository** and provides the necessary logic to process and return the data to the handler.

The **Product Service** offers two key functions:
1. **GetProducts()**: Retrieves all products from the repository.
2. **GetProductById(productId int)**: Retrieves a specific product by its ID from the repository.

### **Product Repository**:

- The **Product Repository** is responsible for interacting with the database to fetch product data.
- It defines two primary methods:
  1. **GetProducts()**: Fetches all products from the database.
  2. **GetProductById(productId int)**: Fetches a single product by its ID from the database.

### Key Points:

- **Fiber**: The web framework used to handle routing and HTTP requests.
- **GORM**: Used for ORM-based database interactions with **PostgreSQL**.
- **Models**: The **Product** model is defined with fields such as `id`, `product_name`, `product_code`, and `product_price`.
- **`/utils`**: Contains helper functions like the **env loader** for loading environment variables.
- **Main Function**: The **main** function is located in the `/cmd` directory, where the HTTP server is initialized, and routes are set up.

## 🚀 Getting Started

### 1️⃣ Install Dependencies

Ensure Go modules are up to date:

```sh
go mod tidy
```

### 2️⃣ Run with Docker

To start the service along with PostgreSQL:

```sh
docker-compose up --build
```

### 3️⃣ Run Manually (Without Docker)

If you prefer to run the service manually:

```sh
go run cmd/main.go
```

Make sure to configure your `.env` file correctly before running the service.

## ✅ Running Unit Tests

To execute all tests:

```sh
go test ./...
```

To run tests for a specific package:

```sh
go test ./internal/handler
```

Unit tests utilize mock data for testing, so no real database connection is required.

## 📖 API Documentation

API endpoints are documented using Swagger/OpenAPI. Once the service is running, access the API documentation at:

```
http://localhost:8082/swagger
```
