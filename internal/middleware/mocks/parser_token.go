package mocks

import "errors"

func MockParseToken(tokenString, secretKey string) (string, error) {
	if tokenString == "ValidToken" {
		return "662d188f705e90f11bac6cb7", nil
	}
	return "", errors.New("invalid token")
}
