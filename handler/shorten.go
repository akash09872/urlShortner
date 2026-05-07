package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"urlshortner/db"
	"urlshortner/storage"
)

type URL struct {
	FullUrl string `json:"url"`
}
type Shortened struct {
	Short string `json:"short"`
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode(length int) string {
	code := ""

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(chars))

		code += string(chars[randomIndex])
	}

	return code
}
func CodeExists(code string) bool {

	var exists bool

	err := db.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM urls WHERE short_code = $1)",
		code,
	).Scan(&exists)

	if err != nil {
		panic(err)
	}

	return exists
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var url URL
	json.NewDecoder(r.Body).Decode(&url)
	if url.FullUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: "URL Required",
		})
		return
	}
	code := GenerateCode(6)
	for CodeExists(code) {
		code = GenerateCode(6)
	}
	if len(url.FullUrl) <= 8 || url.FullUrl[0:8] != "https://" {
		url.FullUrl = "https://" + url.FullUrl
	}
	storage.SaveUrl(code, url.FullUrl)
	s := Shortened{Short: code}
	json.NewEncoder(w).Encode(s)
}
