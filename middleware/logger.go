package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request: ", r.Method, r.URL.Path)
		next(w, r)
	}
}
