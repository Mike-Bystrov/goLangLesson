package hmiddlewares

import (
	"context"
	"net/http"
	"strings"
	"github.com/thnxvlad/oplati/internal/service/authorization"
)

type AccountIdContextKey struct{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		accountId, err := auth.GetAccountIdFromToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AccountIdContextKey{}, accountId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
