package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler, logger *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
