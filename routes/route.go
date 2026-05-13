package routes

import (
	"net/http"
	"urlshortner/handler"
	"urlshortner/middleware"
)

func Route() {

	http.HandleFunc("/my-urls", middleware.Logger(middleware.RateLimit(middleware.AuthMiddleware(handler.MyUrls))))
	http.HandleFunc("/login", middleware.Logger(middleware.RateLimit(handler.Login)))
	http.HandleFunc("/signup", middleware.Logger(middleware.RateLimit(handler.Signup)))
	http.HandleFunc("/shorten", middleware.Logger(middleware.RateLimit(middleware.AuthMiddleware(handler.Shorten))))
	http.HandleFunc("/", middleware.Logger(middleware.RateLimit(handler.Redirect)))
}
