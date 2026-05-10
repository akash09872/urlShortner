package routes

import (
	"net/http"
	"urlshortner/handler"
	"urlshortner/middleware"
)

func Route() {
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		if r != nil || w != nil {
			return
		}
		return
	})
	http.HandleFunc("/my-urls", middleware.Logger(middleware.AuthMiddleware(handler.MyUrls)))
	http.HandleFunc("/login", middleware.Logger(handler.Login))
	http.HandleFunc("/signup", middleware.Logger(handler.Signup))
	http.HandleFunc("/shorten", middleware.Logger(middleware.AuthMiddleware(handler.Shorten)))
	http.HandleFunc("/", middleware.Logger(handler.Redirect))
}
