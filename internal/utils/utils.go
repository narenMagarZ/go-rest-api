package utils

import (
	"fmt"
	"rest-api/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JwtTokenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

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

func GenerateToken(email string) (string, error) {
	claims := JwtTokenClaims{
		Email: email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(config.AppConfig().JwtSecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}


func VerifyToken(token string) (*JwtTokenClaims, error) {
	// parses, validates, verifies the signature and returns the parsed token
	// keyFunc received parsed token and must return cryptographic key for verifying the signature
	jwtToken, err := jwt.ParseWithClaims(token, &JwtTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig().JwtSecretKey), nil
	})

	if err != nil {
		fmt.Println("Failed to verify jwt token")
	}

	if claims, ok := jwtToken.Claims.(*JwtTokenClaims); ok && jwtToken.Valid {
		return claims, nil
	}

	var error error = fmt.Errorf("Error occured verifying token")
	return nil, error
}