# Go URL Shortener Backend

## What'S Built

A basic backend URL shortener using Go and `net/http`.

Features:
- HTTP server
- Routing
- Query parameters
- JSON handling
- Middleware
- URL shortening
- Redirect system
- File persistence using JSON
- Environment variables using `.env`

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
│   ├── store.go
│   ├── file.go
│   └── data.json
│
├── model/
│   └── url.go
│
├── .env
├── go.mod
└── go.sum
```

---

# Concepts Covered

## 1. HTTP Server

```go
http.ListenAndServe(":4000", nil)
```

Starts backend server.

---

## 2. Routing

```go
http.HandleFunc("/shorten", handler.Shorten)
```

Maps route to handler.

---

## 3. Handlers

```go
func Shorten(w http.ResponseWriter, r *http.Request)
```

Handles request and response.

---

## 4. Query Parameters

```go
r.URL.Query().Get("url")
```

Reads URL query values.

---

## 5. JSON Encoding/Decoding

### Decode

```go
json.NewDecoder(r.Body).Decode(&data)
```

### Encode

```go
json.NewEncoder(w).Encode(data)
```

---

## 6. Middleware

```go
logger(handler.Shorten)
```

Wraps handler with extra functionality.

Example:
- logging
- auth
- CORS

---

## 7. Redirects

```go
http.Redirect(w, r, url, http.StatusFound)
```

Redirects short URL to original URL.

---

## 8. Environment Variables

### .env

```env
PORT=4000
```

### Load

```go
godotenv.Load()
```

### Read

```go
os.Getenv("PORT")
```

---

## 9. File Persistence

### Save map to JSON file

```go
json.NewEncoder(file).Encode(Store)
```

### Load map from file

```go
json.NewDecoder(file).Decode(&Store)
```

---

# Commands Used

## Initialize Go Module

```bash
go mod init urlshortner
```

---

## Run Server

```bash
go run .
```

---

## Install dotenv package

```bash
go get github.com/joho/godotenv
```

---

# API Flow

```text
Client
 ↓
Route
 ↓
Middleware
 ↓
Handler
 ↓
Storage
 ↓
Response
```

---

# URL Shortener Flow

```text
POST /shorten
↓
Read URL
↓
Generate short code
↓
Store in map
↓
Save to file
↓
Return short code
```

---
