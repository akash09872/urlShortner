package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request: ", r.Method, r.URL.Path, strings.Split(r.RemoteAddr, ":")[0])
		next(w, r)
	}
}
