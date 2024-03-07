package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var secrete_key = "SECRETE_KEY"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secrete_key))

	if err != nil {
		return "", err
	}
	return tokenString, err

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	tokenJWT, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)

		if !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secrete_key), nil

	})

	if err != nil {
		return nil, err
	}

	return tokenJWT, nil

}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	claims, isOK := token.Claims.(jwt.MapClaims)

	if isOK && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
