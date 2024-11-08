package token

import (
	// "fmt"
	// "strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// FOR INTERPRICE MUST STORAGE THIS CONST IN ANOTHER PLACE
var secretKey = []byte("asdfaslj1234")

func MakeToken(email string) (string, error) {
	payload := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string) (string, bool, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", false, err
	}

	email, err := claims.GetSubject()
	if err != nil {
		return "", false, err
	}
	exp, err := claims.GetExpirationTime()
	if err != nil {
		return "", false, err
	}

	return email, exp.After(time.Now()), nil
}
