package utils

import (
	"tbo_backend/objects"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(mobile *string, name *string) string {

	// generate bearer token from the information
	userClaims := objects.UserClaims{
		Name:    *name,
		Contact: *mobile,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1000000 * time.Hour)), // Token expiration time
			Issuer:    "tbo_hive",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(objects.ConfigObj.SecretKey))
	if err != nil {
		return ""
	}

	return tokenString
}
