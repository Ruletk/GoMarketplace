package auth

import (
	"fmt"
	"github.com/Ruletk/GoMarketplace/pkg/communication"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
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

// NoAuthMiddleware is a middleware that checks if the user is authenticated.
// If the user is authenticated, it returns an error message and aborts the request.
func NoAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Internal-Call") == "true" {
			c.Next()
			return
		}
		_, err := c.Cookie("token")
		if err == nil {
			logging.Logger.Info("User is already authenticated, aborting.")
			c.JSON(http.StatusForbidden, ApiResponse{
				Code:    http.StatusForbidden,
				Type:    "error",
				Message: "You are already authenticated",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// CookieTokenMiddleware is a middleware that checks if the user is authenticated.
// If the user is authenticated, it sets the token in the context.
func CookieTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Internal-Call") == "true" {
			c.Next()
			return
		}
		logging.Logger.Info(c.Request.Cookies())
		logging.Logger.Info(c.Request.Header)
		for _, cookie := range c.Request.Cookies() {
			logging.Logger.Info("Cookie: ", cookie.Name, "=", cookie.Value)
		}
		token, err := c.Cookie("token")
		logging.Logger.Info("Token: ", token)
		if err != nil {
			logging.Logger.Info("No token provided, aborting.")
			c.JSON(http.StatusUnauthorized, ApiResponse{
				Code:    http.StatusUnauthorized,
				Type:    "error",
				Message: "No token provided",
			})
			c.Abort()
		}
		tokenValidation := validateToken(token)

		c.Set(TokenKey, token)
		c.Set(TokenValidationKey, tokenValidation)

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
