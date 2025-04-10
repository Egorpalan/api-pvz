package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Egorpalan/api-pvz/pkg/jwt"
)

type contextKey string

const RoleKey contextKey = "userRole"

func AuthMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// Проверка роли
			authorized := false
			for _, role := range allowedRoles {
				if claims.Role == role {
					authorized = true
					break
				}
			}

			if !authorized {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			// Прокинем роль в context
			ctx := context.WithValue(r.Context(), RoleKey, claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
