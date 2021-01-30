package swagger

import (
	"log"
	"net/http"
	"time"
)

// Logger is a logging middleware to output requests and their response times
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// CorsHeaderSetter middleware sets the Access-Control-Allow-Origin header and
// returns the handler
func CorsHeaderSetter(inner http.Handler, origin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		inner.ServeHTTP(w, r)
	})
}
