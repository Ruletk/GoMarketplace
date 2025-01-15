package service

import (
	"auth/internal/messages"
	"auth/internal/repository"
	"errors"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService interface {
	Login(req *messages.AuthRequest) (*messages.AuthResponse, error)
	Register(req *messages.AuthRequest) (*messages.AuthResponse, error)
	Logout(token string) error
	SendVerificationEmail(userID int64) error
	ChangePassword(req *messages.PasswordChangeRequest) error
	ResetPassword(req *messages.PasswordChange, token string) error
	VerifyUser(token string) error
	GetUserData(userID int64) (*messages.AuthDataResponse, error)
}

type authService struct {
	authRepo       repository.AuthRepository
	sessionService SessionService
	jwtService     JwtService
	emailService   EmailService
}

func NewAuthService(authRepo repository.AuthRepository, sessionService SessionService, jwtService JwtService, emailService EmailService) AuthService {
	return &authService{
		authRepo:       authRepo,
		sessionService: sessionService,
		jwtService:     jwtService,
		emailService:   emailService,
	}
}

// Login authenticates a user
func (a authService) Login(req *messages.AuthRequest) (*messages.AuthResponse, error) {
	logging.Logger.Info("Authenticating user with email: ", req.Email, "...")

	user, err := a.authRepo.GetByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logging.Logger.Debug("User with email: ", req.Email, " not found")
		return nil, err
	} else if err != nil {
		logging.Logger.WithError(err).Error("Failed to get user by email: ", req.Email)
		return nil, err
	}

	if !user.ComparePassword(req.Password) {
		// For security. If someone tries to brute force the password
		logging.Logger.WithFields(logrus.Fields{
			"type":  "auth_attempt",
			"email": req.Email,
		}).Debug("Invalid credentials for user with email: ", req.Email)
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
	logging.Logger.Info("Registering user with email: ", req.Email, "...")

	_, err := a.authRepo.GetByEmail(req.Email)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		logging.Logger.Debug("User with email: ", req.Email, " already exists")
		return nil, gorm.ErrDuplicatedKey
	} else if err == nil {
		// User exists
		logging.Logger.WithError(err).Error("Unexpected error.")
		return nil, err
	}

	user := &repository.Auth{
		Email:        req.Email,
		PasswordHash: "",
	}
	user.PasswordHash = user.GeneratePasswordHash(req.Password)

	logging.Logger.Debug("User model created: ", user)

	err = a.authRepo.Create(user)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to create user: ")
		return nil, err
	}
	logging.Logger.Debug("User with email: ", req.Email, " created successfully, id: ", user.ID)

	session, err := a.sessionService.CreateSession(user.ID)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to create session: ")
		return nil, err
	}

	logging.Logger.Debug("Session created with token: ", session.Token[:5], "...")

	// Send verification email after registration
	err = a.SendVerificationEmail(user.ID)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to send verification email.")
		return nil, err
	}

	return &session, nil
}

// Logout logs out a user
func (a authService) Logout(token string) error {
	logging.Logger.Info("Logging out user with token: ", token[:10], "...")
	return a.sessionService.DeleteSession(token)
}

func (a authService) SendVerificationEmail(userID int64) error {
	logging.Logger.Info("Sending verification email for user with ID: ", userID, "...")
	user, err := a.authRepo.GetByID(userID)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get user by ID.")
		return err
	}

	if user.Active {
		logging.Logger.Warn("User with ID: ", userID, " is already verified")
		return errors.New("user is already verified")
	}

	logging.Logger.Debug("User found: ", user, ". Generating verification token...")
	token, err := a.jwtService.GenerateVerificationToken(userID)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to generate verification token.")
		return err
	}
	logging.Logger.Debug("Token generated: ", token[:10], ". Sending verification email...")

	err = a.emailService.SendVerificationEmail(user.Email, token)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to send verification email.")
		return err
	}

	logging.Logger.Debug("Verification email sent successfully for user with ID: ", userID)
	return nil
}

// ChangePassword requests a password change for a user. Link is sent to the user's email
func (a authService) ChangePassword(req *messages.PasswordChangeRequest) error {
	logging.Logger.Debug("Sending changing password request for user with email: ", req.Email, "...")
	user, err := a.authRepo.GetByEmail(req.Email)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get user by email.")
		return err
	}
	logging.Logger.Debug("User found: ", user, ". Generating password reset token...")
	token, err := a.jwtService.GeneratePasswordResetToken(user.ID)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to generate password reset token.")
		return err
	}

	logging.Logger.Debug("Sending password reset email...")
	err = a.emailService.SendPasswordResetEmail(user.Email, token)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to send password reset email.")
		return err
	}

	logging.Logger.Debug("Password reset email sent successfully")
	return nil
}

// ResetPassword resets the password for a user
func (a authService) ResetPassword(req *messages.PasswordChange, token string) error {
	// Verify token
	logging.Logger.Info("Resetting password for token: ", token[:10], "...")
	valid, userID := a.jwtService.IsPasswordResetToken(token)
	if valid == false {
		logging.Logger.Debug("Provided token is not valid")
		return jwt.ErrTokenInvalidClaims
	}

	// Find user by token
	logging.Logger.Debug("Getting user by ID: ", userID, "...")
	user, err := a.authRepo.GetByID(userID)
	if err != nil {
		logging.Logger.WithError(err).Debug("Failed to get user by ID: ", userID)
		return err
	}

	// Update user password
	logging.Logger.Debug("Updating user password...")
	user.GeneratePasswordHash(req.NewPassword)
	err = a.authRepo.Update(user)
	if err != nil {
		logging.Logger.WithError(err).Debug("Failed to update user.")
		return err
	}

	// Delete token
	err = a.jwtService.DeleteToken(token)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to delete token.")
		return err
	}

	return nil
}

// VerifyUser verifies a user
func (a authService) VerifyUser(token string) error {
	// Verify token
	logging.Logger.Info("Verifying user with token: ", token[:10], "...")
	valid, userID := a.jwtService.IsVerificationToken(token)
	if valid == false {
		logging.Logger.Debug("Provided token is not valid")
		return jwt.ErrTokenInvalidClaims
	}

	logging.Logger.Debug("Verifying user with ID: ", userID, "...")
	err := a.authRepo.VerifyUser(userID)
	if err != nil {
		logging.Logger.WithError(err).Debug("Failed to verify user.")
		return err
	}

	logging.Logger.Debug("User verified successfully, deleting token: ", token[:10], "...")
	err = a.jwtService.DeleteToken(token)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to delete token.")
		return err
	}

	return nil
}

// GetUserData returns user data
func (a authService) GetUserData(userID int64) (*messages.AuthDataResponse, error) {
	logging.Logger.Info("Getting user data for ID: ", userID)
	user, err := a.authRepo.GetByID(userID)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to get user by ID.")
		return nil, err
	}
	logging.Logger.Debug("User found: ", user.ID)

	return &messages.AuthDataResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}
