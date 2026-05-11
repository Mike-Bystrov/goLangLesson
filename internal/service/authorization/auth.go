package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("JWT_SECRET")

type userClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func generateToken(userID string) (string, error) {
	now := time.Now()
	claims := &userClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24*time.Hour)),
			IssuedAt: jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer: "oplati-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

var loginData = map[string]string{
	"user1": "password1",
}

var accountData = map[string]string{
	"user1": "993",
}

func validateLogin(login, password string) bool {
	return password == loginData[login]
}

func Login(login, password string) (string, error) {
	if !validateLogin(login, password) {
		return "",errors.New("invalid login or password")
	}

	accountId, ok := accountData[login]
	if !ok {
		return "", errors.New("accout not found")
	}

	return generateToken(accountId)
}

func getAccountIdFromToken(token string) (string, error) {
	claims, err := jwt.ParseWithClaims(token, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	parsedClaims, ok := claims.Claims.(*userClaims)
	if !ok {
		return "", errors.New("invalid claims type")
	}
	return parsedClaims.UserID, nil
}

type AccountIdContextKey struct{}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		accountId, err := getAccountIdFromToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AccountIdContextKey{}, accountId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
