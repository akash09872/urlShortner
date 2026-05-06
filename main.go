package main

import (
	"fmt"
	"net/http"
	"os"
	"urlshortner/handler"
	"urlshortner/middleware"
	"urlshortner/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	http.HandleFunc("/shorten", middleware.Logger(handler.Shorten))
	http.HandleFunc("/", middleware.Logger(handler.Redirect))
	fmt.Println("Server Started")
	storage.Load()
	http.ListenAndServe(":"+PORT, nil)
}
