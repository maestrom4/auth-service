package utils_test

// import (
// 	u "auth-service/utils"
// 	"errors"
// 	"testing"

// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // MockTokenParser is a mock type for the jwt parsing function
// // type MockTokenParser struct {
// // 	mock.Mock
// // }

// // func (m *MockTokenParser) ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
// // 	args := m.Called(tokenString, claims, keyFunc)
// // 	return args.Get(0).(*jwt.Token), args.Error(1)
// // }

// // // TestParseToken tests the ParseToken function
// // func TestParseToken(t *testing.T) {
// // 	tests := []struct {
// // 		name           string
// // 		tokenString    string
// // 		secretKey      string
// // 		mockReturn     *jwt.Token
// // 		mockError      error
// // 		expectedError  error
// // 		expectedUserID string
// // 	}{
// // 		{
// // 			name:           "valid token",
// // 			tokenString:    "valid.token.string",
// // 			secretKey:      "secret",
// // 			mockReturn:     &jwt.Token{Valid: true},
// // 			mockError:      nil,
// // 			expectedError:  nil,
// // 			expectedUserID: "123456",
// // 		},
// // 		{
// // 			name:           "invalid token",
// // 			tokenString:    "invalid.token.string",
// // 			secretKey:      "secret",
// // 			mockReturn:     nil,
// // 			mockError:      errors.New("invalid token"),
// // 			expectedError:  errors.New("invalid token"),
// // 			expectedUserID: "",
// // 		},
// // 	}

// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			mockTokenParser := new(MockTokenParser)
// // 			mockClaims := &struct {
// // 				jwt.RegisteredClaims
// // 				UserID string `json:"userId"`
// // 			}{} // This should be your actual type
// // 			mockTokenParser.On(
// // 				"ParseWithClaims",
// // 				tt.tokenString,
// // 				mock.IsType(mockClaims),
// // 				mock.AnythingOfType("jwt.Keyfunc"),
// // 			).Return(tt.mockReturn, tt.mockError)
// // 			// Use the TokenParser interface to parse the token
// // 			userID, err := u.ParseToken(tt.tokenString, tt.secretKey, mockTokenParser)

// // 			if tt.expectedError != nil {
// // 				assert.EqualError(t, err, tt.expectedError.Error())
// // 			} else {
// // 				assert.NoError(t, err)
// // 				assert.Equal(t, tt.expectedUserID, userID)
// // 			}

// // 			mockTokenParser.AssertExpectations(t)
// // 		})
// // 	}
// // }
