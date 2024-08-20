package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	SECRET_KEY = []byte("dajgdhjgrjghdjagjhsdkashdhvrak")
	JWT_ALGO   = jwt.SigningMethodHS256
)

type Claims struct {
	userName string
	userId   string
	jwt.RegisteredClaims
}

func CreateJwtToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(JWT_ALGO, claims)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != JWT_ALGO {
			return nil, errors.New("unexpected signing method")
		}
		return SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// How to use it.
func main() {

	claims := &Claims{
		userName: "Mohmad Sabban",
		userId:   "1",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "your-app-name",
			Audience:  jwt.ClaimStrings{"your-app-audience"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(2000) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := CreateJwtToken(*claims)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

}
