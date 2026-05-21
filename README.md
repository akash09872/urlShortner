# URL Shortener Backend (Go + PostgreSQL)

A backend URL shortener built using Go, PostgreSQL, Docker, and the standard `net/http` package.

This project was built while learning backend development fundamentals step-by-step without relying on frameworks too early.

---

# Features

* URL shortening with unique short-code generation
* Redirect using short code
* PostgreSQL persistence
* JWT-based authentication
* User signup and login
* bcrypt password hashing
* Protected routes using middleware
* User-specific URL ownership
* Optional URL expiry support
* Background cleanup jobs for expired URLs
* Request logging middleware
* User-based API rate limiting
* JSON request/response handling
* Layered backend structure
* Environment variable support using `.env`

---

# Tech Stack

* Go
* PostgreSQL
* Docker
* net/http
* JWT
* bcrypt

---

# Project Structure

```text
urlshortner/
│
├── main.go
│
├── handler/
│   ├── auth.go
│   └── shorten.go
│
├── middleware/
│   ├── auth.go
│   ├── logger.go
│   └── ratelimit.go
│
├── storage/
│   ├── helper.go
│   ├── url.go
│   └── user.go
│
├── database/
│   └── db.go
│
├── auth/
│   ├── jwt.go
│   └── password.go
│
├── jobs/
│   └── cleanup.go
│
├── model/
│   ├── response.go
│   ├── url.go
│   └── user.go
│
├── routes/
│   └── routes.go
│
├── postgres-data/
│
├── docker-compose.yml
├── .env
├── go.mod
└── go.sum
```

---

# Environment Variables

Create a `.env` file:

```env
PORT=4000

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=YOUR_PASSWORD
DB_NAME=urlshortner

JWT_SECRET=YOUR_SECRET_KEY
```

---

# Running PostgreSQL with Docker

## docker-compose.yml

```yaml
services:
  postgres:
    image: postgres
    container_name: postgres-db

    environment:
      POSTGRES_PASSWORD: YOUR_PASSWORD
      POSTGRES_DB: urlshortner

    ports:
      - "5432:5432"

    volumes:
      - ./postgres-data:/var/lib/postgresql
```

---

# Start PostgreSQL

```bash
docker compose up -d
```

---

# Enter PostgreSQL Shell

```bash
docker exec -it postgres-db psql -U postgres
```

---

# Connect Database

```sql
\c urlshortner
```

---

# Database Schema

## Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
```

## URLs Table

```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,

    short_code TEXT UNIQUE NOT NULL,

    original_url TEXT NOT NULL,

    user_id INT REFERENCES users(id),

    expires_at TIMESTAMP
);
```

---

# Install Dependencies

```bash
go get github.com/lib/pq
go get github.com/joho/godotenv
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get golang.org/x/time/rate
```

---

# Run Backend

```bash
go run .
```

---

# API Endpoints

## Authentication

### Signup

```http
POST /signup
Content-Type: application/json
```

### Request Body

```json
{
  "username": "akash",
  "password": "hello123"
}
```

---

### Login

```http
POST /login
Content-Type: application/json
```

### Request Body

```json
{
  "username": "akash",
  "password": "hello123"
}
```

### Response

```json
{
  "token": "JWT_TOKEN"
}
```

---

## URL Shortening

### Create Short URL

```http
POST /shorten
Authorization: Bearer JWT_TOKEN
Content-Type: application/json
```

### Request Body

```json
{
  "url": "https://google.com"
}
```

### Optional Expiry Support

```json
{
  "url": "https://google.com",
  "expires_in_hours": 24
}
```

### Response

```json
{
  "short": "Ab12Xq"
}
```

---

## Redirect

```http
GET /Ab12Xq
```

Redirects user to the original URL if valid and not expired.

---

## User URLs

### Get All URLs Created By Logged-in User

```http
GET /my-urls
Authorization: Bearer JWT_TOKEN
```

### Response

```json
[
  {
    "short_code": "Ab12Xq",
    "original_url": "https://google.com"
  }
]
```

---

# Backend Request Flow

```text
Client
 ↓
Route
 ↓
Middleware
 ↓
Handler
 ↓
Storage Layer
 ↓
PostgreSQL
 ↓
Response
```

---

# Middleware Features

* JWT Authentication Middleware
* Request Logging Middleware
* User-based Rate Limiting Middleware
* Protected Routes

---

# Concepts Learned

* HTTP fundamentals
* Routing using `net/http`
* JSON encoding/decoding
* Middleware chaining
* JWT authentication
* Password hashing with bcrypt
* Request lifecycle
* PostgreSQL integration
* Docker basics
* Environment variables
* Background jobs using goroutines
* API rate limiting
* Layered backend architecture
* Relational database design
* SQL queries from Go

---

# Future Improvements

* Refresh tokens
* Visit analytics
* Pagination
* Swagger documentation
* Redis-based distributed rate limiting
* Database migrations
* Role-based authorization
* Deployment and CI/CD

---

# Learning Purpose

This project was built to deeply understand backend fundamentals, authentication systems, middleware architecture, relational databases, and scalable backend design without depending on frameworks too early.
