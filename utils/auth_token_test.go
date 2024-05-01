package utils_test

import (
	"encoding/base64"
	"testing"

	u "auth-service/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateVerificationToken(t *testing.T) {
	testCases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Token should not be empty",
			wantErr: false,
		},
		{
			name:    "Token should be valid base64",
			wantErr: false,
		},
		{
			name:    "Token should have length greater than 0",
			wantErr: false,
		},
		{
			name:    "Expected token length does not match",
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := u.GenerateVerificationToken()

			if !tc.wantErr {
				require.NoError(t, err)
				assert.NotEmpty(t, token, "Token should not be empty")
				_, decodeErr := base64.StdEncoding.DecodeString(token)
				assert.NoError(t, decodeErr, "Token should be valid base64")
				assert.True(t, len(token) > 0, "Token should have length greater than 0")
				assert.Equal(t, 44, len(token), "Expected token length does not match")

			}
		})
	}
}
