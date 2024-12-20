package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthAPI struct {
}

func NewAuthAPI() *AuthAPI {
	return &AuthAPI{}
}

func (api *AuthAPI) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", api.Login)
	router.POST("/register", api.Register)
	router.POST("/logout", api.Logout)
	router.POST("/change-password", api.ChangePassword)
	router.POST("/change-password/:token", api.ChangePasswordWithToken)
	router.GET("/verify/:token", api.Verify)
}

func (api *AuthAPI) Login(c *gin.Context) {
	var req AuthRequest
	err := c.BindJSON(&req)
	// Check if the request is valid
	if err != nil {

		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Authenticate the user
	// Let's assume that the logic is here
	// TODO Implements the logic to authenticate the user
	if req.Email == "user@gmail.com" && req.Password == "password" {
		c.JSON(http.StatusOK, AuthResponse{
			Token: "eyJpdiI6Inhwd3VZTG1PeVR6cG5KVUpUcFBBb",
		})
		return
	}

	// Return an error if the user is not authenticated
	c.JSON(http.StatusUnauthorized, ApiResponse{
		Code:    401,
		Type:    "error",
		Message: "Wrong email or password",
	})

}

func (api *AuthAPI) Register(c *gin.Context) {
	var req AuthRequest
	err := c.BindJSON(&req)
	// Check if the request is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Register the user
	// Let's assume that the logic is here
	// TODO Implements the logic to register the user
	if req.Email == "user@gmail.com" && req.Password == "password" {
		c.JSON(http.StatusOK, AuthResponse{
			Token: "eyJpdiI6Inhwd3VZTG1PeVR6cG5KVUpUcFBBb",
		})
		return
	}

	// Return an error if the user is already registered
	c.JSON(http.StatusConflict, ApiResponse{
		Code:    http.StatusConflict,
		Type:    "error",
		Message: "User with this email already registered",
	})
}

func (api *AuthAPI) Logout(c *gin.Context) {
	var req TokenRequest
	err := c.BindJSON(&req)
	// Check if the request is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request or token",
		})
		return
	}

	// Logout the user
	// Let's assume that the logic is here
	// TODO Implements the logic to logout the user
	if req.Token == "token" {
		c.JSON(http.StatusOK, ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "User logged out",
		})
		return
	}

	// Return an error if the token is invalid
	c.JSON(http.StatusUnauthorized, ApiResponse{
		Code:    http.StatusUnauthorized,
		Type:    "error",
		Message: "Invalid token",
	})
}

func (api *AuthAPI) ChangePassword(c *gin.Context) {
	var req PasswordChangeRequest
	err := c.BindJSON(&req)
	// Check if the request is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request or email",
		})
		return
	}

	// Send an email with a token to the user
	// Let's assume that the logic is here
	// TODO Implements the logic to send an email with a token
	// Split the email to get the domain
	domain := string([]rune(req.Email)[strings.Index(req.Email, "@")+1:])
	c.JSON(http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Password change request sent successfully. Check your email for further instructions. Domain: " + domain,
	})
}

func (api *AuthAPI) ChangePasswordWithToken(c *gin.Context) {
	token := c.Param("token")
	var req PasswordChange
	err := c.BindJSON(&req)
	// Check if the request is valid
	if err != nil || token == "" {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Change the password
	// Let's assume that the logic is here
	// TODO Implements the logic to change the password
	if req.NewPassword == "newPassword" {
		c.JSON(http.StatusOK, ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Password changed successfully",
		})
		return
	}

	// Return an error if the token is invalid
	c.JSON(http.StatusUnauthorized, ApiResponse{
		Code:    http.StatusUnauthorized,
		Type:    "error",
		Message: "Invalid token",
	})
}

func (api *AuthAPI) Verify(c *gin.Context) {
	token := c.Param("token")
	// Check if the token is valid
	if token == "" {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Verify the token
	// Let's assume that the logic is here
	// TODO Implements the logic to verify the token
	if token == "token" {
		c.JSON(http.StatusOK, ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Successfully verified",
		})
		return
	}

	// Return an error if the token is invalid
	c.JSON(http.StatusUnauthorized, ApiResponse{
		Code:    http.StatusUnauthorized,
		Type:    "error",
		Message: "Invalid token",
	})
}
