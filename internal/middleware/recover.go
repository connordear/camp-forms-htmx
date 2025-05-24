package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoverPanic(errLog *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Connection", "close")

					trace := fmt.Sprintf("%s\n%s", err, debug.Stack())
					errLog.Output(2, trace)

					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
