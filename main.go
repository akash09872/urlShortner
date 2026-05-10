package main

import (
	"fmt"
	"net/http"
	"os"
	"urlshortner/db"
	"urlshortner/routes"
	"urlshortner/storage"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	db.Connect()
	go storage.Clean()
	routes.Route()
	fmt.Println("Server Started")
	http.ListenAndServe(":"+PORT, nil)
}
