# User Management API

A RESTful API built with Go, Fiber, PostgreSQL, and SQLC to manage users and calculate age dynamically from their date of birth.

## Features

* Create User
* Get User By ID
* List All Users
* Update User
* Delete User
* Dynamic Age Calculation
* Input Validation using go-playground/validator
* Structured Logging using Uber Zap
* Pagination Support
* Request ID Middleware
* Request Duration Logging Middleware
* Docker Support

## Tech Stack

* Go
* Fiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator
* Docker

## Project Structure

```text
cmd/
└── server/
    └── main.go

config/

db/
├── migrations/
├── query/
└── sqlc/

internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── logger/
└── models/
```

## Prerequisites

Before running the application, make sure the following are installed:

* Go 1.24 or later
* PostgreSQL
* Git

## Environment Variables

Create a `.env` file in the project root:

```env
APP_PORT=3000

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=user_management
DB_SSLMODE=disable
```

## Database Setup

### Create Database

Connect to PostgreSQL and run:

```sql
CREATE DATABASE user_management;
```

### Create Users Table

Connect to the `user_management` database and run:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

## Clone the Repository

```bash
git clone https://github.com/avusulavivekchary/user-management-api.git
cd user-management-api
```

## Install Dependencies

```bash
go mod tidy
```

## Run the Application

Start the API server:

```bash
go run cmd/server/main.go
```

If everything is configured correctly, the server will start on:

```text
http://localhost:3000
```

## API Endpoints

### Create User

**POST** `/users`

Request:

```json
{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

Response:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### Get User By ID

**GET** `/users/1`

Response:

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

---

### List Users

**GET** `/users`

Pagination:

```http
GET /users?page=1&limit=10
```

Example Response:

```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 35
  }
]
```

---

### Update User

**PUT** `/users/1`

Request:

```json
{
  "name": "Alice Updated",
  "dob": "1991-01-01"
}
```

Response:

```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-01-01"
}
```

---

### Delete User

**DELETE** `/users/1`

Response:

```http
204 No Content
```

## Running Tests

Run all tests:

```bash
go test ./...
```

## Docker

Build the Docker image:

```bash
docker compose build
```

Run the application using Docker:

```bash
docker compose up
```

## Author

Avusula Vivek Chary
