package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/time/rate"
)

var limiters = make(map[string]*rate.Limiter)

func getLimiter(username string) *rate.Limiter {
	limiter, exists := limiters[username]
	if !exists {
		limiter = rate.NewLimiter(1, 5)
		limiters[username] = limiter
	}
	return limiter
}
func RateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value("username").(string)
		var limiter *rate.Limiter
		ip := strings.Split(r.RemoteAddr, ":")[0]
		if username == "" || !ok {
			limiter = getLimiter(ip)
		} else {
			limiter = getLimiter(username)
		}
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", 429)
			fmt.Println("Blocked Request", username, ip)
			return
		}
		next(w, r)
	}
}
