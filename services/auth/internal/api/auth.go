package api

import (
	"auth/internal/messages"
	"auth/internal/service"
	"errors"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type AuthAPI struct {
	authService    service.AuthService
	sessionService service.SessionService
}

func NewAuthAPI(authService service.AuthService, sessionService service.SessionService) *AuthAPI {
	return &AuthAPI{authService: authService, sessionService: sessionService}
}

// RegisterPublicRoutes registers the public routes for the auth API
// These routes may require a token
func (api *AuthAPI) RegisterPublicRoutes(router *gin.RouterGroup) {
	logging.Logger.Info("Registering public routes")
	router.GET("/verify/:token", api.Verify)
}

// RegisterPublicOnlyRoutes registers the public only routes for the auth API
// These routes do not require a token
func (api *AuthAPI) RegisterPublicOnlyRoutes(router *gin.RouterGroup) {
	logging.Logger.Info("Registering public only routes")
	router.POST("/login", api.Login)
	router.POST("/register", api.Register)
	router.POST("/change-password", api.ChangePassword)
	router.POST("/change-password/:token", api.ChangePasswordWithToken)
}

// RegisterPrivateRoutes registers the private routes for the auth API
// These routes require a token
func (api *AuthAPI) RegisterPrivateRoutes(router *gin.RouterGroup) {
	logging.Logger.Info("Registering private routes")
	router.GET("/logout", api.Logout)
	router.POST("/validate", api.ValidateToken)
	//	Admin routes
	router.DELETE("/admin/sessions/hard-delete", api.HardDeleteSessions)
	router.DELETE("/admin/sessions/delete-inactive", api.DeleteInactiveSessions)
}

func (api *AuthAPI) Login(c *gin.Context) {
	logging.Logger.Info("Logging in user")

	// Parse the request
	var req messages.AuthRequest
	err := c.ShouldBindJSON(&req)

	// Check if the request is valid
	if err != nil {
		logging.Logger.WithError(err).Error("Invalid request")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	// Authenticate the user
	resp, err := api.authService.Login(&req)
	if errors.Is(err, service.ErrInvalidCredentials) {
		logging.Logger.WithError(err).Error("Wrong email or password")
		c.JSON(http.StatusUnauthorized, messages.ApiResponse{
			Code:    http.StatusUnauthorized,
			Type:    "error",
			Message: "Wrong email or password",
		})
		return
	} else if err != nil {
		logging.Logger.WithError(err).Error("Internal server error")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Internal server error. Details: " + err.Error(),
		})
		return
	}

	if resp != nil {
		logging.Logger.Info("User logged in successfully")
		c.SetCookie("token", resp.Token, 31536000, "/", "", false, true)
		c.JSON(http.StatusOK, resp)
		return
	}

	logging.Logger.Error("Wrong email or password")
	// Return an error if the user is not authenticated
	c.JSON(http.StatusUnauthorized, messages.ApiResponse{
		Code:    401,
		Type:    "error",
		Message: "Wrong email or password",
	})

}

func (api *AuthAPI) Register(c *gin.Context) {
	logging.Logger.Info("Registering user")

	var req messages.AuthRequest
	err := c.ShouldBindJSON(&req)

	// Check if the request is valid
	if err != nil {
		logging.Logger.WithError(err).Error("Invalid request")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	logging.Logger.Info("Registering user with email: " + req.Email)

	// Register the user
	resp, err := api.authService.Register(&req)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		logging.Logger.WithError(err).Error("User with this email already registered")
		c.JSON(http.StatusConflict, messages.ApiResponse{
			Code:    http.StatusConflict,
			Type:    "error",
			Message: "User with this email already registered",
		})
		return
	} else if err != nil {
		logging.Logger.WithError(err).Error("Internal server error")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Internal server error. Details: " + err.Error(),
		})
		return
	}

	if resp != nil {
		logging.Logger.Info("User registered successfully")
		c.SetCookie("token", resp.Token, 31536000, "/", "", false, true)
		c.JSON(http.StatusOK, resp)
		return
	}

	logging.Logger.Error("User with this email already registered")
	// WTF, this is a duplicate of the error above
	// TODO: Refactor this
	// Return an error if the user is already registered
	c.JSON(http.StatusConflict, messages.ApiResponse{
		Code:    http.StatusConflict,
		Type:    "error",
		Message: "User with this email already registered",
	})
}

func (api *AuthAPI) Logout(c *gin.Context) {
	logging.Logger.Info("Logging out user")

	token, _ := c.Get("token")

	// Logout the user
	_ = api.authService.Logout(token.(string))

	logging.Logger.Info("User logged out successfully, token: " + token.(string))

	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Successfully logged out",
	})
}

func (api *AuthAPI) ChangePassword(c *gin.Context) {
	logging.Logger.Info("Changing password")

	var req messages.PasswordChangeRequest
	err := c.ShouldBindJSON(&req)

	// Check if the request is valid
	if err != nil {
		logging.Logger.WithError(err).Error("Invalid request")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request or email",
		})
		return
	}

	// Send an email with a token to the user
	logging.Logger.Info("Changing password for email: " + req.Email)
	err = api.authService.ChangePassword(&req)
	if err == nil {
		logging.Logger.Info("Password change sent for email: " + req.Email)
		domain := string([]rune(req.Email)[strings.Index(req.Email, "@")+1:])
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Password change request sent successfully. Check your email for further instructions. Domain: " + domain,
		})
		return
	}

	// Return an error if the email is invalid
	logging.Logger.WithError(err).Error("Invalid request or email")
	c.JSON(http.StatusBadRequest, messages.ApiResponse{
		Code:    http.StatusBadRequest,
		Type:    "error",
		Message: "Invalid request or email",
	})
}

func (api *AuthAPI) ChangePasswordWithToken(c *gin.Context) {
	logging.Logger.Info("Changing password with token")

	token := c.Param("token")
	var req messages.PasswordChange
	err := c.ShouldBindJSON(&req)

	// Check if the request is valid
	if err != nil || token == "" {
		logging.Logger.WithError(err).Error("Invalid request")
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
		logging.Logger.Info("Password changed successfully")
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Password changed successfully",
		})
		return
	}

	// Return an error if the token is invalid
	logging.Logger.WithError(err).Error("Invalid token")
	c.JSON(http.StatusUnauthorized, messages.ApiResponse{
		Code:    http.StatusUnauthorized,
		Type:    "error",
		Message: "Invalid token",
	})
}

func (api *AuthAPI) Verify(c *gin.Context) {
	//TODO: Refactor this to send X-Access-Token header
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
		logging.Logger.Info("Sessions deleted successfully")
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Sessions deleted successfully",
		})
		return
	}

	logging.Logger.WithError(err).Error("Internal server error")
	c.JSON(http.StatusInternalServerError, messages.ApiResponse{
		Code:    http.StatusInternalServerError,
		Type:    "error",
		Message: "Internal server error. Details: " + err.Error(),
	})
}

func (api *AuthAPI) DeleteInactiveSessions(c *gin.Context) {
	// TODO: Add admin check
	logging.Logger.Info("Starting delete all inactive sessions...")

	err := api.sessionService.DeleteInactiveSessions()

	if err == nil {
		logging.Logger.Info("Sessions deleted successfully")
		c.JSON(http.StatusOK, messages.ApiResponse{
			Code:    http.StatusOK,
			Type:    "success",
			Message: "Sessions deleted successfully",
		})
		return
	}

	logging.Logger.WithError(err).Error("Internal server error")

	c.JSON(http.StatusInternalServerError, messages.ApiResponse{
		Code:    http.StatusInternalServerError,
		Type:    "error",
		Message: "Internal server error. Details: " + err.Error(),
	})
}

func (api *AuthAPI) ValidateToken(c *gin.Context) {
	//TODO: Delete this method
	var req messages.TokenRequest
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

	// Validate the token
	userID, err := api.sessionService.GetUserID(req.Token)
	if err != nil {
		logging.Logger.Debug(err)
		c.JSON(http.StatusUnauthorized, messages.ApiResponse{
			Code:    401,
			Type:    "error",
			Message: "Invalid token",
		})
		return
	}

	logging.Logger.Debug(err)

	resp, err := api.authService.GetUserData(userID)
	if err == nil {
		c.JSON(http.StatusOK, resp)
		return
	}

	logging.Logger.Error(err)

	// Return an error if the token is invalid
	c.JSON(http.StatusUnauthorized, messages.ApiResponse{
		Code:    http.StatusUnauthorized,
		Type:    "error",
		Message: "Invalid token",
	})
}
