package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"main.go/models"
)

type ClientClaims struct {
	ID    uint
	Phone string
	Email string
	Role  string
	jwt.RegisteredClaims
}

func TokenGenerate(user *models.ClientToken, role string) (string, error) {
	claims := ClientClaims{
		ID:    user.ID,
		Phone: user.Phone,
		Email: user.Email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cityvibe",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString([]byte(os.Getenv("TOKENSECRETKEY")))
	if err != nil {
		return "", err
	}
	return TokenString, nil
}

func AdminTokenGenerate(user models.Admin, role string) (string, error) {
	claims := ClientClaims{
		ID:    user.ID,
		Email: user.Email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cityvibe",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString([]byte(os.Getenv("TOKENSECRETKEY")))
	if err != nil {
		return "", err
	}
	return TokenString, nil
}

func GetRoleFromToken(Token string) (string, error) {
	TokenUnpacked, err := jwt.ParseWithClaims(Token, ClientClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRETKEY")), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := TokenUnpacked.Claims.(*ClientClaims); ok && TokenUnpacked.Valid {
		return claims.Role, nil
	}
	return "", fmt.Errorf("invalid token")
}
