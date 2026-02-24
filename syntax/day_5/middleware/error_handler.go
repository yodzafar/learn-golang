package middleware

import (
	"learn-golang/syntax/day_5/utils"
	"log"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Println("Recovered from panic", rec)
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)
				utils.JSON(w, http.StatusInternalServerError, utils.ApiResponse{
					Success: false,
					Error:   http.StatusText(http.StatusInternalServerError),
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
