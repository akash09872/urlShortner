package handler

import (
	"encoding/json"
	"net/http"
	"urlshortner/auth"
	"urlshortner/db"
	"urlshortner/model"
	"urlshortner/storage"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, _ := auth.HashPassword(user.Password)

	err := storage.SaveUser(user.Username, hashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Message: "Signup Error",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{
		Message: "DONE",
	})
}
func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, err := GetPass(user.Username)
	if err != nil {
		http.Error(w, "Invalid Username", 400)
		return
	}
	valid := auth.CheckPassword(
		hashedPassword, user.Password,
	)
	if !valid {
		http.Error(w, "Invalid Password", 400)
		return
	}
	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		http.Error(w, "TOKEN GENERATION ERROR", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
func GetPass(username string) (string, error) {
	var hashedPassword string
	err := db.DB.QueryRow(
		"SELECT password FROM users WHERE username=$1",
		username,
	).Scan(&hashedPassword)
	return hashedPassword, err
}
func MyUrls(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	userId, err := storage.GetUserId(username)
	if err != nil {
		http.Error(w, "user Not found", 400)
		return
	}
	urls := storage.GetUserURLs(userId)
	if urls == nil {
		http.Error(w, "No urls found", 500)
	}
	json.NewEncoder(w).Encode(urls)
}
