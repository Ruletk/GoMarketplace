package api

import (
	"auth/internal/messages"
	"auth/internal/service"
	"errors"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type AuthAPI struct {
	authService    service.AuthService
	sessionService service.SessionService
	tokenService   service.TokenService
}

func NewAuthAPI(authService service.AuthService, sessionService service.SessionService, tokenService service.TokenService) *AuthAPI {
	return &AuthAPI{authService: authService, sessionService: sessionService, tokenService: tokenService}
}

func (api *AuthAPI) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", api.Login)
	router.POST("/register", api.Register)
	router.POST("/logout", api.Logout)
	router.POST("/change-password", api.ChangePassword)
	router.POST("/change-password/:token", api.ChangePasswordWithToken)
	router.GET("/verify/:token", api.Verify)

	//	Admin routes
	router.GET("/admin/sessions/hard-delete", api.HardDeleteSessions)
	router.GET("/admin/sessions/delete-inactive/:time", api.DeleteInactiveSessions)
}

func (api *AuthAPI) Login(c *gin.Context) {
	var req messages.AuthRequest
	err := c.ShouldBindJSON(&req)
	// Check if the request is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Authenticate the user
	resp, err := api.authService.Login(&req)
	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, service.ErrInvalidCredentials) {
		c.JSON(http.StatusUnauthorized, messages.ApiResponse{
			Code:    http.StatusUnauthorized,
			Type:    "error",
			Message: "Wrong email or password",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Internal server error. Details: " + err.Error(),
		})
		return
	}

	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	// Return an error if the user is not authenticated
	c.JSON(http.StatusUnauthorized, messages.ApiResponse{
		Code:    401,
		Type:    "error",
		Message: "Wrong email or password",
	})

}

func (api *AuthAPI) Register(c *gin.Context) {
	var req messages.AuthRequest
	err := c.ShouldBindJSON(&req)
	// Check if the request is valid
	if err != nil {
		logging.Logger.Debug(err)
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Register the user
	resp, err := api.authService.Register(&req)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		logging.Logger.Debug(err)
		c.JSON(http.StatusConflict, messages.ApiResponse{
			Code:    http.StatusConflict,
			Type:    "error",
			Message: "User with this email already registered",
		})
		return
	} else if err != nil {
		logging.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Internal server error. Details: " + err.Error(),
		})
		return
	}

	if resp != nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	// Return an error if the user is already registered
	c.JSON(http.StatusConflict, messages.ApiResponse{
		Code:    http.StatusConflict,
		Type:    "error",
		Message: "User with this email already registered",
	})
}

func (api *AuthAPI) Logout(c *gin.Context) {
	var req messages.TokenRequest
	err := c.ShouldBindJSON(&req)
	// Check if the request is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request or token",
		})
		return
	}

	// Logout the user
	err = api.authService.Logout(req)
	if err == nil {
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Successfully logged out",
		})
		return
	}

	logging.Logger.Debug(err)

	// Return an error if the token is invalid
	c.JSON(http.StatusUnauthorized, messages.ApiResponse{
		Code:    http.StatusUnauthorized,
		Type:    "error",
		Message: "Invalid token",
	})
}

func (api *AuthAPI) ChangePassword(c *gin.Context) {
	var req messages.PasswordChangeRequest
	err := c.ShouldBindJSON(&req)
	// Check if the request is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request or email",
		})
		return
	}

	// Send an email with a token to the user
	err = api.authService.ChangePassword(&req)
	if err == nil {
		domain := string([]rune(req.Email)[strings.Index(req.Email, "@")+1:])
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Password change request sent successfully. Check your email for further instructions. Domain: " + domain,
		})
		return
	}

	// Return an error if the email is invalid
	c.JSON(http.StatusBadRequest, messages.ApiResponse{
		Code:    http.StatusBadRequest,
		Type:    "error",
		Message: "Invalid request or email",
	})
}

func (api *AuthAPI) ChangePasswordWithToken(c *gin.Context) {
	token := c.Param("token")
	var req messages.PasswordChange
	err := c.ShouldBindJSON(&req)
	// Check if the request is valid
	if err != nil || token == "" {
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Change the password
	err = api.authService.ResetPassword(&req, token)
	if err == nil {
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Password changed successfully",
		})
		return
	}

	// Return an error if the token is invalid
	c.JSON(http.StatusUnauthorized, messages.ApiResponse{
		Code:    http.StatusUnauthorized,
		Type:    "error",
		Message: "Invalid token",
	})
}

func (api *AuthAPI) Verify(c *gin.Context) {
	token := c.Param("token")
	// Check if the token is valid
	if token == "" {
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Verify the token
	err := api.authService.VerifyUser(token)
	if err == nil {
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "User verified successfully",
		})
		return
	}

	// Return an error if the token is invalid
	c.JSON(http.StatusUnauthorized, messages.ApiResponse{
		Code:    http.StatusUnauthorized,
		Type:    "error",
		Message: "Invalid token",
	})
}

func (api *AuthAPI) HardDeleteSessions(c *gin.Context) {
	// TODO: Add admin check
	logging.Logger.Info("Starting delete all expired sessions...")
	err := api.sessionService.HardDeleteSessions()
	if err == nil {
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Sessions deleted successfully",
		})
		return
	}

	logging.Logger.Error(err)
	c.JSON(http.StatusInternalServerError, messages.ApiResponse{
		Code:    http.StatusInternalServerError,
		Type:    "error",
		Message: "Internal server error. Details: " + err.Error(),
	})
}

func (api *AuthAPI) DeleteInactiveSessions(c *gin.Context) {
	// TODO: Add admin check
	logging.Logger.Info("Starting delete all inactive sessions...")
	time, err := strconv.Atoi(c.Param("time"))
	if err != nil {
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}
	// Prevent accidental deletion of active sessions
	if time < 7 {
		time = 7
	}

	logging.Logger.Info("Deleting sessions older than ", time, " days...")
	err = api.sessionService.DeleteInactiveSessions()

	if err == nil {
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Sessions deleted successfully",
		})
		return
	}

	logging.Logger.Error(err)

	c.JSON(http.StatusInternalServerError, messages.ApiResponse{
		Code:    http.StatusInternalServerError,
		Type:    "error",
		Message: "Internal server error. Details: " + err.Error(),
	})
}
