package utils

import (
	cfg "auth-service/internal/config"
	"auth-service/internal/models"
	t "auth-service/internal/types"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userID string, secretKey string) (string, error) {
	if userID == "" {
		return "", errors.New("empty userID")
	}

	claims := t.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    cfg.AppName,
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey string) (*models.User, error) {
	claims := &t.TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token has expired")
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				return nil, errors.New("signature validation failed")
			}
			// Add more checks based on the errors you wish to catch
		}
		return nil, errors.New("token is invalid")
	}

	// Ensure the userID is available in the token
	if claims.UserID == "" {
		return nil, errors.New("user ID not found in token")
	}

	// At this point, the token is valid, and UserID is present, construct the user model
	user := &models.User{
		ID: claims.UserID,
	}
	return user, nil
}

func ParseToken(tokenString, secretKey string) (string, error) {
	claims := &t.TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return "", errors.New("invalid or expired token")
	}

	return claims.UserID, nil
}
