package utils

import (
	cfg "auth-service/internal/config"
	"auth-service/internal/models"
	t "auth-service/internal/types"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ValidateToken(tokenString string) (*models.User, error) {
	secretKey := cfg.JwtSecretKey
	if secretKey == "" {
		return nil, errors.New("JWT secret key is not set")
	}

	claims := &t.TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*t.TokenClaims); ok && token.Valid {
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, errors.New("token has expired")
		}

		user := &models.User{
			ID: claims.UserID,
		}
		return user, nil
	}

	return nil, errors.New("invalid token")
}
