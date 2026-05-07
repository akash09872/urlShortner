# URL Shortener Backend (Go + PostgreSQL)

A backend URL shortener built using Go, PostgreSQL, Docker, and the standard `net/http` package.

This project was built while learning backend development fundamentals step-by-step.

---

# Features

- URL shortening
- Redirect using short code
- PostgreSQL persistence
- Dockerized PostgreSQL
- Middleware support
- JSON request/response handling
- Layered backend structure
- Environment variable support using `.env`

---

# Tech Stack

- Go
- PostgreSQL
- Docker
- net/http

---

# Project Structure

```text
urlshortner/
│
├── main.go
│
├── handler/
│   └── shorten.go
│
├── middleware/
│   └── logger.go
│
├── storage/
│   ├── helper.go
│   └── url.go
│
├── database/
│   └── db.go
│
├── model/
│   ├── response.go
│   └── url.go
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
DB_PASSWORD=YOUR_PASS
DB_NAME=urlshortner
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
      POSTGRES_PASSWORD: 1234
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

# Create Table

```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_code TEXT UNIQUE,
    original_url TEXT NOT NULL
);
```

---

# Install Dependencies

```bash
go get github.com/lib/pq
go get github.com/joho/godotenv
```

---

# Run Backend

```bash
go run .
```

---

# API Endpoints

## Shorten URL

### Request

```http
POST /shorten
Content-Type: application/json
```

### Body

```json
{
  "url": "https://google.com"
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

Redirects user to original URL.

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

# Concepts Learned

- HTTP fundamentals
- Routing using `net/http`
- JSON encoding/decoding
- Middleware
- Request lifecycle
- PostgreSQL integration
- Docker basics
- Environment variables
- Layered backend architecture
- SQL queries from Go

---

# Future Improvements

- Proper error handling
- Duplicate URL prevention
- Random short code optimization
- Visit analytics
- Authentication
- Rate limiting
- Custom response format
- Database migrations

---

# Learning Purpose

This project was built to deeply understand backend fundamentals instead of relying on frameworks too early.

