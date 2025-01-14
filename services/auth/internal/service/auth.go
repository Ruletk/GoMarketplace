package service

import (
	"auth/internal/messages"
	"auth/internal/repository"
	"errors"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService interface {
	Login(req *messages.AuthRequest) (*messages.AuthResponse, error)
	Register(req *messages.AuthRequest) (*messages.AuthResponse, error)
	Logout(token string) error
	ChangePassword(req *messages.PasswordChangeRequest) error
	ResetPassword(req *messages.PasswordChange, token string) error
	VerifyUser(token string) error
	GetUserData(userID int64) (*messages.AuthDataResponse, error)
}

type authService struct {
	authRepo       repository.AuthRepository
	sessionService SessionService
	jwtService     JwtService
}

func NewAuthService(authRepo repository.AuthRepository, sessionService SessionService, jwtService JwtService) AuthService {
	return &authService{
		authRepo:       authRepo,
		sessionService: sessionService,
		jwtService:     jwtService,
	}
}

// Login authenticates a user
func (a authService) Login(req *messages.AuthRequest) (*messages.AuthResponse, error) {
	logging.Logger.Debug("Authenticating user with email: ", req.Email, "...")

	user, err := a.authRepo.GetByEmail(req.Email)
	if err != nil {
		logging.Logger.Debug("User with email: ", req.Email, " not found")
		return nil, err
	}

	if !user.ComparePassword(req.Password) {
		// Unsafe logging, delete in production
		// TODO: Implement proper logging
		logging.Logger.Debug("Invalid credentials for user with email: ", req.Email, "Password: ", req.Password, "User password: ", user.PasswordHash)
		return nil, ErrInvalidCredentials
	}

	logging.Logger.Debug("User with email: ", req.Email, " authenticated successfully, creating session...")

	session, err := a.sessionService.CreateSession(user.ID)
	if err != nil {
		return nil, err
	}

	logging.Logger.Debug("Session created with token: ", session.Token[:5])

	return &session, nil
}

// Register creates a new user
func (a authService) Register(req *messages.AuthRequest) (*messages.AuthResponse, error) {
	logging.Logger.Debug("Registering user with email: ", req.Email, "...")

	_, err := a.authRepo.GetByEmail(req.Email)
	if err == nil {
		logging.Logger.Debug("User with email: ", req.Email, " already exists")
		return nil, gorm.ErrDuplicatedKey
	}

	user := &repository.Auth{
		Email:        req.Email,
		PasswordHash: "",
	}
	user.PasswordHash = user.GeneratePasswordHash(req.Password)

	logging.Logger.Debug("User model created: ", user)

	err = a.authRepo.Create(user)
	if err != nil {
		logging.Logger.Error("Failed to create user: ", err)
		return nil, err
	}
	logging.Logger.Debug("User with email: ", req.Email, " created successfully, id: ", user.ID)

	session, err := a.sessionService.CreateSession(user.ID)
	if err != nil {
		return nil, err
	}

	logging.Logger.Debug("Session created with token: ", session.Token[:5], "...")
	return &session, nil
}

// Logout logs out a user
func (a authService) Logout(token string) error {
	logging.Logger.Debug("Logging out user with token: ", token[:10], "...")
	return a.sessionService.DeleteSession(token)
}

// ChangePassword requests a password change for a user. Link is sent to the user's email
func (a authService) ChangePassword(req *messages.PasswordChangeRequest) error {
	user, err := a.authRepo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	token, err := a.jwtService.GeneratePasswordResetToken(user.ID)
	if err != nil {
		logging.Logger.Error("Failed to generate password reset token: ", err)
		return err
	}
	// to not make error
	_ = token

	// Send email with link to change password
	// TODO: Implement email sending
	return nil
}

// ResetPassword resets the password for a user
func (a authService) ResetPassword(req *messages.PasswordChange, token string) error {
	// Verify token
	logging.Logger.Debug("Resetting password for token: ", token[:10], "...")
	valid, userID := a.jwtService.IsPasswordResetToken(token)
	if valid == false {
		logging.Logger.Debug("Provided token is not valid")
		return jwt.ErrTokenInvalidClaims
	}

	// Find user by token
	logging.Logger.Debug("Getting user by ID: ", userID, "...")
	user, err := a.authRepo.GetByID(userID)
	if err != nil {
		logging.Logger.Debug("Failed to get user by ID: ", userID, " - ", err)
		return err
	}

	// Update user password
	logging.Logger.Debug("Updating user password...")
	user.GeneratePasswordHash(req.NewPassword)
	err = a.authRepo.Update(user)
	if err != nil {
		logging.Logger.Debug("Failed to update user: ", err)
		return err
	}

	// Delete token
	// TODO: Implement token marking as used

	return nil
}

// VerifyUser verifies a user
func (a authService) VerifyUser(token string) error {
	// Verify token
	logging.Logger.Debug("Verifying user with token: ", token[:10], "...")
	valid, userID := a.jwtService.IsVerificationToken(token)
	if valid == false {
		logging.Logger.Debug("Provided token is not valid")
		return jwt.ErrTokenInvalidClaims
	}

	// TODO: Add new repository method to change user status in one query
	// Find user by token
	user, err := a.authRepo.GetByID(userID)
	if err != nil {
		return err
	}

	// Update user
	user.Active = true
	err = a.authRepo.Update(user)

	// Delete token
	// TODO: Implement token marking as used

	return nil
}

// GetUserData returns user data
func (a authService) GetUserData(userID int64) (*messages.AuthDataResponse, error) {
	user, err := a.authRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return &messages.AuthDataResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
