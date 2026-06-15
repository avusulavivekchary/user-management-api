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
├── logger/
├── middleware/
├── models/
├── repository/
├── routes/
└── service/
```

## Setup

### Prerequisites

* Go 1.24+
* PostgreSQL
* Git

### Environment Variables

Create a `.env` file:

```env
APP_PORT=3000

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=user_management
DB_SSLMODE=disable
```

### Database Setup

Create the database:

```sql
CREATE DATABASE user_management;
```

Create the users table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);
```

## Example Response

### GET /users/1

```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

\# User Management API



A RESTful User Management API built with Go, Fiber, PostgreSQL, and SQLC.



\## Features



\* Create User

\* Get All Users

\* Get User By ID

\* Update User

\* Delete User

\* Age Calculation

\* Request Validation

\* PostgreSQL Integration

\* SQLC Query Generation



\## Tech Stack



\* Go

\* Fiber

\* PostgreSQL

\* SQLC

\* Git



\## API Endpoints



\### Create User



POST /users



```json

{

&#x20; "name": "Alice",

&#x20; "dob": "1990-05-10"

}

```



\### Get All Users



GET /users



\### Get User By ID



GET /users/:id



\### Update User



PUT /users/:id



```json

{

&#x20; "name": "Alice Updated",

&#x20; "dob": "1991-01-01"

}

```



\### Delete User



DELETE /users/:id



\## Run Locally



```bash

git clone <repository-url>

cd user-management-api

go mod tidy

go run cmd/server/main.go

```



\## Database



```sql

CREATE TABLE users (

&#x20;   id SERIAL PRIMARY KEY,

&#x20;   name TEXT NOT NULL,

&#x20;   dob DATE NOT NULL

);

```



