package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"main.go/models"
)

type ClientClaims struct {
	ID    uint   `jsom:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func TokenGenerate(user *models.ClientToken, role string) (string, error) {
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
	TokenUnpacked, err := jwt.ParseWithClaims(Token, &ClientClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("1")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKENSECRETKEY")), nil
	})
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return "", err
	}

	if claims, ok := TokenUnpacked.Claims.(*ClientClaims); ok && TokenUnpacked.Valid {
		return claims.Role, nil
	}
	fmt.Println("3")
	return "", fmt.Errorf("invalid token")
}

func ExtractUserIdFromToken(Token string) (uint, error) {

	TokenUnpacked, err := jwt.ParseWithClaims(Token, &ClientClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("1")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKENSECRETKEY")), nil
	})
	if err != nil {

		fmt.Println(err)
		return 0, err
	}

	if claims, ok := TokenUnpacked.Claims.(*ClientClaims); ok && TokenUnpacked.Valid {
		return claims.ID, nil
	}

	return 0, fmt.Errorf("invalid token")
}
