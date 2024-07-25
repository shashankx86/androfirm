package api

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// RateLimiter is a global rate limiter with a rate of 60 requests per minute.
var limiter = rate.NewLimiter(rate.Every(time.Minute), 60)

// RateLimitMiddleware is a middleware that limits the rate of requests.
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
