package middleware

import (
	"context"
	"log"
	"net/http"
)

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("X-User-ID")
		if userId == "" {
			log.Printf("empty X-User-ID header")
		}

		ctx := context.WithValue(r.Context(), "user", userId)
		newReq := r.WithContext(ctx)
		next.ServeHTTP(w, newReq)
	})
}
