# User API

A RESTful User Management API built using Go, Fiber, PostgreSQL, and SQLC. The application provides CRUD operations for managing users while dynamically calculating age from the stored date of birth.

## Features

* Create, Read, Update, and Delete (CRUD) users
* PostgreSQL database integration
* SQLC-generated database queries
* Repository pattern architecture
* Dynamic age calculation from date of birth
* Input validation using go-playground/validator
* Structured logging using Uber Zap
* Request logging middleware
* Request ID middleware
* Docker support
* Pagination support for user listing
* Unit testing for age calculation service

## Project Structure

```text
/cmd/server/main.go
/config/
/db/migrations/
/db/sqlc/
/internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── models/
└── logger/
```

## Tech Stack

* Go
* Fiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator
* Docker

## Running Locally

### Prerequisites

* Go 1.24+
* PostgreSQL 18+
* Git

### Clone Repository

```bash
git clone https://github.com/vigneshxprabhu/user-api.git
cd user-api
```

### Install Dependencies

```bash
go mod download
```

### Create Database

```sql
CREATE DATABASE user_api;
```

Connect to the database:

```sql
\c user_api
```

Create the users table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

### Configure Database Connection

Set the DATABASE_URL environment variable:

```powershell
$env:DATABASE_URL="postgres://postgres:<password>@localhost:5432/user_api?sslmode=disable"
```

### Run Application

```bash
go run ./cmd/server
```

The API will be available at:

```text
http://localhost:8080
```

---

## Running with Docker

Build and start the application:

```bash
docker compose up --build
```

The API will be available at:

```text
http://localhost:8080
```

To stop containers:

```bash
docker compose down
```

---

## API Endpoints

| Method | Endpoint   | Description       |
| ------ | ---------- | ----------------- |
| GET    | /users     | Get all users     |
| GET    | /users/:id | Get user by ID    |
| POST   | /users     | Create a new user |
| PUT    | /users/:id | Update a user     |
| DELETE | /users/:id | Delete a user     |

### Pagination

The user listing endpoint supports pagination:

```http
GET /users?page=1&limit=5
```

Example:

```http
GET /users?page=2&limit=5
```

Returns the next set of users.

---

## Example Request

### Create User

```http
POST /users
```

Request Body:

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

Example Response:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

---

## Validation Rules

* Name is required
* DOB must be in YYYY-MM-DD format
* DOB cannot be a future date

---

## Logging

The application uses Uber Zap for structured logging.

Logged events include:

* User creation
* User updates
* User deletion

Request logging middleware records:

* HTTP method
* Request path
* Request duration

Each request is assigned a unique Request ID using the `X-Request-ID` header.

---

## Running Tests

Run all tests:

```bash
go test ./...
```

---

## Bonus Features Implemented

* Docker support
* Request ID middleware
* Request logging middleware
* Pagination support
* Unit testing for age calculation service

---

## Future Improvements

* Database migrations automation
* Authentication and authorization
* Advanced filtering and sorting
* API documentation using Swagger/OpenAPI
* Integration tests

```
```
