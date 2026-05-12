package auth

import (
	"errors"
	"os"
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
		return "", errors.New("account not found")
	}

	return generateToken(accountId)
}

func SignUp(login, password, userId string) error {
	_, exists := loginData[login]

	if !exists {
		return errors.New("login already exists")
	}

	loginData[login] = password
	accountData[login] = userId

	return nil
}

func GetAccountIdFromToken(token string) (string, error) {
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