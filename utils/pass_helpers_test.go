package utils_test

import (
	"errors"
	"testing"

	cfg "auth-service/internal/config"
	"auth-service/utils"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	// Define test cases as a table
	testCases := []struct {
		Password string

		ExpectedError error
	}{
		{

			Password:      "maestrom4",
			ExpectedError: nil,
		},
		{
			Password:      "",
			ExpectedError: errors.New("empty password"),
		},
	}

	// Initialize logrus logger
	logger := logrus.New()

	for _, tc := range testCases {
		t.Run(tc.Password, func(t *testing.T) {
			tokenString, err := utils.HashPassword(tc.Password)

			// Log the token
			logger.WithFields(logrus.Fields{
				"Password": tc.Password,
			}).Info("Generated Password")

			// Assertions
			if tc.ExpectedError != nil {
				assert.Error(t, err, "Expected error should not be nil")
				assert.EqualError(t, err, tc.ExpectedError.Error(), "Error message should match")
				assert.Empty(t, tokenString, "Token should be empty")
			} else {
				assert.NoError(t, err, "Error should be nil")
				assert.NotEmpty(t, tokenString, "Token should not be empty")
			}
		})
	}
}

func TestCreateToken(t *testing.T) {
	// Define test cases as a table
	testCases := []struct {
		Name          string
		UserID        string
		ExpectedError error
	}{
		{
			Name:          "danmats",
			UserID:        "maestrom4",
			ExpectedError: nil,
		},
		{
			Name:          "danmats2",
			UserID:        "",
			ExpectedError: errors.New("empty userID"),
		},
	}

	// Initialize logrus logger
	logger := logrus.New()

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Call the function with test case parameters
			tokenString, err := utils.CreateToken(tc.UserID, cfg.SaltPassKey)

			// Log the token
			logger.WithFields(logrus.Fields{
				"test_name": tc.Name,
				"token":     tokenString,
			}).Info("Generated token")

			// Assertions
			if tc.ExpectedError != nil {
				assert.Error(t, err, "Expected error should not be nil")
				assert.EqualError(t, err, tc.ExpectedError.Error(), "Error message should match")
				assert.Empty(t, tokenString, "Token should be empty")
			} else {
				assert.NoError(t, err, "Error should be nil")
				assert.NotEmpty(t, tokenString, "Token should not be empty")
			}
		})
	}
}
