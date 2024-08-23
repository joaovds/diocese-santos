package auth

import "net/http"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("auth middleware")
		next.ServeHTTP(w, r)
	}
}
