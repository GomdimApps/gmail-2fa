package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT with the provided claims
func GenerateToken(claims jwt.Claims) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "mysecret"
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenJWT.SignedString([]byte(secretKey))
}

// GenerateUserToken creates a JWT token for a user
func GenerateUserToken(userID, email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := UserClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "linxygestor-api",
			Subject:   userID,
		},
	}
	return GenerateToken(claims)
}
