package middleware

import "net/http"

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if user is aunthenticated
		token := r.Header.Get("Authorization")
		if token != "Bearer Secret" {
			http.Error(w,"Unathorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w,r)
	})
}