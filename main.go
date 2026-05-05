package main

import (
	"fmt"
	"net/http"
	"urlshortner/handler"
)

func main() {
	http.HandleFunc("/shorten", handler.Shorten)
	http.HandleFunc("/", handler.Redirect)
	fmt.Println("Server Started")
	http.ListenAndServe(":4000", nil)
}
