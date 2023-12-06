# Todo Backend with Gin

## Overview
This repository contains a backend application developed with the Gin-Gonic framework for managing and storing todo notes. The backend provides RESTful API endpoints to perform CRUD (Create, Read, Update, Delete) operations on todo notes.

## Getting Started
To run the project locally, follow these steps:

### Prerequisites
- [Go](https://golang.org/dl/) installed on your machine
- [Gin-Gonic](https://github.com/gin-gonic/gin) package
- [Cors](https://github.com/gin-contrib/cors) middleware package

### Installation
1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/your-username/todo-backend-gin.git
   cd todo-backend-gin
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The application will start on `localhost:8080`.

## API Endpoints
- **Create Note:** `POST /notes`
- **Find All Notes:** `GET /notes`
- **Find Note by ID:** `GET /notes/:id`
- **Update Note:** `PATCH /notes/:id`
- **Delete Note:** `DELETE /notes/:id`

## Database
The application connects to a database to persistently store todo notes. Ensure that the database connection details are configured in the `models.ConnectDatabase()` function.

## CORS Configuration
CORS is enabled by default using the `github.com/gin-contrib/cors` middleware. Adjust CORS settings in the `main.go` file as needed.
