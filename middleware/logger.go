package middleware

import (
	"log"
	"net/http"
)

func RequestLogger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s Path: %s", r.Method, r.URL)
		next.ServeHTTP(w,r)
	})
}