package types

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}
