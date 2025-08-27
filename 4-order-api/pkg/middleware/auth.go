package middleware

import (
	"api/order/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

type key string

const (
	ContextSessionKey key = "ContextSessionKey"
)

type AuthMiddleware struct {
	Secret string
}

func writeUnauth(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func Auth(next http.Handler, secret AuthMiddleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer") {
			writeUnauth(w)
			return
		}
		// Here you would typically validate the auth token
		// For simplicity, we assume the token is valid if it exists
		token := strings.TrimPrefix(auth, "Bearer ")
		isValid, data := jwt.NewJWT(secret.Secret).Parse(token)
		if !isValid {
			writeUnauth(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextSessionKey, data.Phone)

		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
