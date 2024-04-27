package middleware_test

import (
	mdl "auth-service/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		authHeader     string
		expectedUserID string
		expectError    bool
		statusCode     int
	}{
		{
			name:           "Valid Authorization Token",
			authHeader:     "Bearer ValidToken",
			expectedUserID: "662d188f705e90f11bac6cb7",
			expectError:    false,
			statusCode:     http.StatusOK,
		},
		{
			name:           "Invalid Authorization Token",
			authHeader:     "Bearer InvalidToken",
			expectedUserID: "",
			expectError:    true,
			statusCode:     http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.Use(mdl.AuthMiddleware(mdl.MockParseToken))
			router.GET("/graphql", func(c *gin.Context) {
				userID, exists := c.Get("userID")
				if exists {
					c.JSON(http.StatusOK, gin.H{"userID": userID})
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				}
			})

			req := httptest.NewRequest("GET", "/graphql", nil)
			req.Header.Set("Authorization", tt.authHeader)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.statusCode, w.Code)
			if !tt.expectError {
				assert.Contains(t, w.Body.String(), tt.expectedUserID)
			} else {
				assert.Contains(t, w.Body.String(), "Invalid token")
			}
		})
	}
}
