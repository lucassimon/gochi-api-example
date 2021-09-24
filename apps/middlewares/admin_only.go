package middlewares

import (
	"math/rand"
	"net/http"
)

func IsLoggedInAdmin(r *http.Request) bool {
	return rand.Float32() < 0.5
}

// AdminOnly middleware restricts access to just administrators.
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok := IsLoggedInAdmin(r)

		if !ok {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
