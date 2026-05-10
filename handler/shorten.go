package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"urlshortner/db"
	"urlshortner/model"
	"urlshortner/storage"
)

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

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.Context().Value("username").(string)
	userID, err := storage.GetUserId(username)

	if err != nil {
		http.Error(w, "User not found", 400)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var url model.URL
	json.NewDecoder(r.Body).Decode(&url)
	if url.FullUrl == "" {
		http.Error(w, "URL required", http.StatusBadRequest)
		return
	}
	if len(url.FullUrl) <= 8 || url.FullUrl[0:8] != "https://" {
		url.FullUrl = "https://" + url.FullUrl
	}
	var code string

	err = db.DB.QueryRow(
		"SELECT short_code FROM urls WHERE user_id = $2 AND original_url = $1",
		url.FullUrl, userID,
	).Scan(&code)

	if err == nil {
		// already exists
		json.NewEncoder(w).Encode(model.Shortened{
			Short: code,
		})
		db.DB.QueryRow(
			"UPDATE urls SET expires_at = $1 WHERE original_url = $2",
			time.Now().Add(6*time.Hour),
			url.FullUrl,
		)
		return
	}

	code = GenerateCode(6)
	for CodeExists(code) {
		code = GenerateCode(6)
	}
	storage.SaveUrl(code, url.FullUrl, userID)
	s := model.Shortened{Short: code}
	json.NewEncoder(w).Encode(s)
}
