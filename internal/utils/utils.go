package utils

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

func HashText(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost);
	if err != nil {
		return "", err
	}
	return string(hash), nil
}


func CompareHash(text string, hashedText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(text));
}

func createToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp": "",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	signedToken, err := token.SignedString("")
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
