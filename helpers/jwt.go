package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey []byte

// set jwt key globally
func InitJWTKey(key []byte) {
	jwtKey = key
}

type TokenClaims struct {
	Email string `json:"email"`
	Exp   int    `json:"exp"`
	jwt.StandardClaims
}

// generate token
func CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// validate token
func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid Token")
	}

	// Check the expiration time
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return errors.New("token is expired")
	}

	return nil
}

// access token calims
func GetTokenClaims(tokenString string) (*TokenClaims, error) {
	claim := &TokenClaims{}

	// Parse the token into the custom Claims struct
	_, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	return claim, err
}
