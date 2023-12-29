package common

import (
	"fmt"
	"net/http"

	"github.com/EdimarRibeiro/inventory/api/models"
	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(r *http.Request) (string, uint64, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", 0, fmt.Errorf("authorization token not provided")
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return models.SecretKey, nil
	})

	if err != nil {
		return "", 0, err
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok || !token.Valid {
		return "", 0, fmt.Errorf("invalid token")
	}

	return claims.Username, claims.ExternalId, nil
}
