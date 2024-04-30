package utils_test

import (
	"auth-service/internal/types"
	"auth-service/utils"
	"errors"
	"net/smtp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Note this will be moved to types
type MockEmailSender struct {
	mock.Mock
}

// Note also be move to a centralized mock collection having issues while moving.
func (m *MockEmailSender) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	args := m.Called(addr, a, from, to, msg)
	return args.Error(0)
}

func TestSendVerificationEmail(t *testing.T) {
	tests := []struct {
		name          string
		emailOpt      types.EmailOpt
		expectedError error
	}{
		{
			name: "successful email send",
			emailOpt: types.EmailOpt{
				EmailFrom: "from@example.com",
				Password:  "password123",
				Email:     "to@example.com",
				Message:   "Subject Line",
				Body:      "Email Body",
			},
			expectedError: nil,
		},
		{
			name: "failed email send",
			emailOpt: types.EmailOpt{
				EmailFrom: "from@example.com",
				Password:  "password123",
				Email:     "to@example.com",
				Message:   "Subject Line",
				Body:      "Email Body",
			},
			expectedError: errors.New("smtp error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSender := new(MockEmailSender)
			mockSender.On("SendMail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.expectedError)

			err := utils.SendVerificationEmail(tt.emailOpt, mockSender)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			mockSender.AssertExpectations(t)
		})
	}
}
