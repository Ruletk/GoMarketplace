package messages

// AuthRequest represents a login request
type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// TokenRequest represents a token request
type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

// AuthResponse represents a successful authentication response
type AuthResponse struct {
	Token string `json:"token"`
}

// ApiResponse represents a generic API response
type ApiResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// PasswordChangeRequest represents a request to change a password
type PasswordChangeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// PasswordChange represents the new password
type PasswordChange struct {
	NewPassword string `json:"newPassword" binding:"required"`
}
