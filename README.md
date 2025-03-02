# golang_sample_REST_API

## Overview

golang_sample_REST_API is a sample REST API built with Gin, GORM, and Swaggo. It serves as a template for creating RESTful services in Go, featuring endpoints for managing demo entities and user authentication using JWT.

## Features

- RESTful API endpoints for demo resources.
- User registration and login endpoints with JWT authentication.
- Database integration using PostgreSQL with auto-migration via GORM.
- Swagger documentation generated with Swaggo.
- Middleware for JWT authentication.
- Clean project structure to serve as a template for new projects.

## Project Structure

- **cmd/app**: Contains the main entry point (`main.go`) of the application.
- **docs**: Contains generated Swagger documentation files (`docs.go`, `swagger.json`, `swagger.yaml`).
- **internal/config**: Loads application configuration from environment variables.
- **internal/domain/demo**: Contains demo domain logic including model, routes, handler, service, and repository.
- **internal/domain/authentication**: Contains user authentication logic including model, routes, handler, and repository.
- **internal/abstract**: Provides generic repository and service implementations.
- **internal/middleware**: Contains JWT authentication middleware.

## Database
The API uses a PostgreSQL Database. The easiest way is using Docker:

```bash
  docker compose up -d
 ```
   
## Setup and Installation

1. **Clone the repository**
    ```bash
    git clone https://github.com/Lon60/golang_sample_REST_API.git  
    cd golang_sample_REST_API
   ```

2. **Set up Environment Variables**

   Create a `.env` file in the root directory (there is a `.env.example`, modify as needed):

3. **Generate Swagger Documentation**

   Run the following command from the project root:
    ```bash
    swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/app/main.go
    ```
4. **Run the Application**

   Start the application with:
    ```bash
    go run ./cmd/app
    ```
## API Endpoints

### Swagger Documentation

- Access the Swagger UI at:  
  `http://localhost:8080/swagger/index.html` (available when running in debug mode)

### Demo Endpoints

All demo endpoints require a valid JWT token in the `Authorization` header (`Bearer <token>`).

- **Create Demo**:  
  `POST /api/demos/`  
  Payload: JSON representing a demo entity.

- **Get All Demos**:  
  `GET /api/demos/`

- **Get Demo by ID**:  
  `GET /api/demos/{id}`

- **Update Demo**:  
  `PUT /api/demos/{id}`  
  Payload: JSON representing the updated demo entity.

- **Delete Demo**:  
  `DELETE /api/demos/{id}`

### Authentication Endpoints

- **Register**:  
  `POST /api/register`  
  Payload: JSON containing `email` and `password`.

- **Login**:  
  `POST /api/login`  
  Payload: JSON containing `email` and `password`.  
  On success, a JWT token is returned.

## Usage

1. Register a new user using the `/api/register` endpoint.
2. Log in with the registered credentials using the `/api/login` endpoint to obtain a JWT token.
3. Include the token in the `Authorization` header (as `Bearer <token>`) for all protected endpoints (demo endpoints).