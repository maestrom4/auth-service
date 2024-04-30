package utils_test

import (
	"context"
	"testing"

	u "auth-service/utils"

	"github.com/stretchr/testify/assert"
)

func TestGetStringFromContext(t *testing.T) {
	tests := []struct {
		name           string
		ctx            context.Context
		key            string
		expectedString string
		expectedError  string
	}{
		{
			name:           "Valid string value",
			ctx:            context.WithValue(context.Background(), "userID", "123"),
			key:            "userID",
			expectedString: "123",
			expectedError:  "",
		},
		{
			name:           "Missing key",
			ctx:            context.Background(),
			key:            "nonExistentKey",
			expectedString: "",
			expectedError:  "key 'nonExistentKey' not found in context",
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			strValue, err := u.GetStringFromContext(tc.ctx, tc.key)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
				assert.Empty(t, strValue)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedString, strValue)
			}
		})
	}
}

func TestGetBoolFromContext(t *testing.T) {
	tests := []struct {
		name          string
		ctx           context.Context
		key           string
		expectedBool  bool
		expectedError string
	}{
		{
			name:          "Valid boolean value",
			ctx:           context.WithValue(context.Background(), "authorized", true),
			key:           "authorized",
			expectedBool:  true,
			expectedError: "",
		},
		{
			name:          "Missing key",
			ctx:           context.Background(),
			key:           "nonExistentKey",
			expectedBool:  false,
			expectedError: "key 'nonExistentKey' not found in context",
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			boolValue, err := u.GetBoolFromContext(tc.ctx, tc.key)

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError)
				assert.False(t, boolValue)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedBool, boolValue)
			}
		})
	}
}
