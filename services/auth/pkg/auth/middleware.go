package auth

import (
	"fmt"
	"github.com/Ruletk/GoMarketplace/pkg/communication"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	TokenKey           string = "token"
	TokenValidationKey string = "token_validation"
)

// ApiResponse represents a generic API response
// In future, messages will be moved to a separate package
// This is a temporary solution
// TODO: Move to a separate package
type ApiResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func BearerTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Internal-Call") == "true" {
			c.Next()
			return
		}
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		tokenValidation := validateToken(token)

		c.Set(TokenKey, token)
		c.Set(TokenValidationKey, tokenValidation)

		if !tokenValidation {
			c.JSON(http.StatusUnauthorized, ApiResponse{
				Code:    http.StatusUnauthorized,
				Type:    "error",
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func validateToken(token string) bool {
	if token == "" {
		return false
	}
	// TODO: Make a discovery service to get the URL of the auth service
	resp, err := communication.PostJSON("http://web:80/api/v1/auth/validate", []byte(`{"token": "`+token+`"}`))
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
