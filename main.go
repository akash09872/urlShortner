package main

import (
	"fmt"
	"net/http"
	"os"
	"urlshortner/db"
	"urlshortner/handler"
	"urlshortner/middleware"
	"urlshortner/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	host := os.Getenv("DB_HOST")
	fmt.Println(host)
	http.HandleFunc("/shorten", middleware.Logger(handler.Shorten))
	http.HandleFunc("/", middleware.Logger(handler.Redirect))
	db.Connect()
	go storage.Clean()
	fmt.Println("Server Started")
	http.ListenAndServe(":"+PORT, nil)
}
