# Employee Management API

A production-ready backend API built using **Go (Golang)** and **PostgreSQL**.  
This project demonstrates clean backend architecture, secure authentication, database design, migrations, and containerization best practices.

---

## Features

- JWT-based Authentication (Register & Login)
- Employee CRUD operations
- Salary calculation logic
- Salary metrics & aggregations
- PostgreSQL relational database
- Versioned SQL migrations
- Docker & Docker Compose support

---

## Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin
- **Database:** PostgreSQL
- **Authentication:** JWT
- **Migrations:** SQL-based
- **Containerization:** Docker, Docker Compose

---

## Project Structure

employee-management-api/
├── cmd/
│ └── main.go
├── config
│  └── config.go
├── controllers/
│  ├── auth.go
│  ├── controllers.go
│  └── salary.go
├── helpers/
│  ├── auth.go
│  ├── employee.go
│  ├── jwt.go
│  ├── pg_manager.go
│  └── salary.go
├── middlewares/
│  └── auth.go
├── models/
├── routes/
│  └── routes.go
├── db_migrations/
│ ├── 001_create_users.up.sql
│ ├── 001_create_users.down.sql
│ ├── 002_create_employees.up.sql
│ └── 002_create_employees.down.sql
├── ERD.md
├── README.md
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum


---

## Database Design (ERD)

The database schema and relationships are documented in **ERD.md**.

Main entities:
- Users
- Employees

Relationships are normalized and designed to support scalability.

---


## Database Migrations

Migrations are written using raw SQL and stored inside the `migrations/` folder.


### Apply migrations
```bash
migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
Rollback migrations
migrate -path migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down

## Authentication

Authentication is handled using **JWT (JSON Web Token)**.

### Auth Endpoints
| Method | Endpoint   |
|------  |------------|
| POST   | `/api/v1/auth/register`
| POST   | `/api/v1/auth/login`


After login, the JWT token must be passed in the `Authorization` header:
Authorization: Bearer <token>


---

## 👨‍💼 Employee APIs (Protected)

| Method | Endpoint | Description |
|--------|------------------------|-------------------|
|POST    | `/api/v1/employees`    | Create employee   |
|GET     | `/api/v1/employees/:id`| Get employee by ID|
|PUT     | `/api/v1/employees/:id`| Update employee   |
|DELETE  | `/api/v1/employees/:id`| Delete employee   |

---

## Salary APIs (Protected)

### 1: Salary Calculation (Per Employee)

GET /api/v1/employees/:id/salary
Returns:
- country
- Employee_id
- Gross salary
- Deduction
- Net salary (based on country rules)

---


### 2: Salary Metrics

#### For total employee
GET /api/v1/employees/salary/metrics
Returns:
- Minimum salary
- Maximum salary
- Average salary
- Total employees

#### By Country

GET /api/v1/employees/salary/country/:country
Returns:
- Minimum salary
- Maximum salary
- Average salary
- Total employees


#### By Job Title
GET /api/v1/employees/salary/job/:jobtitle
Returns:
- Average salary for the given job title

Edge_cases like **no matching data** are handled gracefully.

---

Go API Server

PostgreSQL Database

API will be available at:

http://localhost:8080

API Testing -
You can test APIs using:
Postman
Curl

Example:

GET /employees
Authorization: Bearer <JWT_TOKEN>

----

Docker Setup
Build & Run using Docker Compose
docker-compose up --build
Services started:

---

## Notes

Environment variables are managed via docker-compose.yml

cp env.example .env

Migrations must be run before accessing APIs

Designed to be production-ready and easily extendable

---

Author 
Hemant Mukati
Backend Developer (Go)

---
