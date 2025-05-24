package middleware

import "net/http"

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: for production have different csp
		// w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src-elem 'self' https://unpkg.com http://localhost:* ws://localhost:* 'unsafe-inline'; style-src-elem 'self' http://localhost:* 'unsafe-inline'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		// w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		// w.Header().Set("X-Content-Type-Options", "nosniff")
		// w.Header().Set("X-Frame-Options", "deny")
		// w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}
