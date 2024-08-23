package auth

import "net/http"

func AuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("auth middleware")
		next.ServeHTTP(w, r)
	})
}
