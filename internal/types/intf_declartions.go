// Interface declartions go
package types

import "github.com/golang-jwt/jwt"

type TokenParser interface {
	ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error)
}
