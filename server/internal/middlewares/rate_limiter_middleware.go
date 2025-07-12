package middlewares

import (
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

const (
	rateLimitRPS   = 1   // requests per second
	rateLimitBurst = 5   // max burst size
)

var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

// getIP extracts the IP from request (can be enhanced to support X-Forwarded-For)
func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func getVisitorLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(rateLimitRPS), rateLimitBurst)
		visitors[ip] = limiter
	}
	return limiter
}

// RateLimiter returns a chi-compatible middleware
func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		limiter := getVisitorLimiter(ip)

		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
