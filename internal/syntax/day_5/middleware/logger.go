package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w}
		start := time.Now()
		next.ServeHTTP(rw, r)
		duration := float64(time.Since(start).Microseconds()) / 1000.0
		log.Printf("%s %s %d %.2f ms", r.Method, r.URL.Path, rw.status, duration)
	})
}
