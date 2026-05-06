package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"urlshortner/storage"
)

type URL struct {
	FullUrl string `json:"url"`
}
type Shortened struct {
	Short string `json:"short"`
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
		http.Error(w, "URL Required", http.StatusBadRequest)
		return
	}

	code := strconv.Itoa(len(storage.Store) + 1)
	if url.FullUrl[0:8] != "https://" {
		url.FullUrl = "https://" + url.FullUrl
	}
	storage.Store[code] = url.FullUrl
	storage.Save()
	s := Shortened{Short: code}
	fmt.Println(code, ": ", storage.Store[code])
	json.NewEncoder(w).Encode(s)
}
