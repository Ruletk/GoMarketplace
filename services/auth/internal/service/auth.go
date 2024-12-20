package service

import (
	"auth/internal/api"
	"auth/internal/repository"
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService interface {
	Login(req *api.AuthRequest) (*api.AuthResponse, error)
	Register(req *api.AuthRequest) (*api.AuthResponse, error)
	ChangePassword(req *api.PasswordChangeRequest) error
	ResetPassword(req *api.PasswordChange, token string) error
	VerifyUser(token string) error
}

type authService struct {
	authRepo       repository.AuthRepository
	sessionService SessionService
}

func NewAuthService(authRepo repository.AuthRepository, sessionService SessionService) AuthService {
	return &authService{
		authRepo:       authRepo,
		sessionService: sessionService,
	}
}

// Login authenticates a user
func (a authService) Login(req *api.AuthRequest) (*api.AuthResponse, error) {
	user, err := a.authRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !user.ComparePassword(req.Password) {
		return nil, ErrInvalidCredentials
	}

	session, err := a.sessionService.CreateSession(user.ID)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

// Register creates a new user
func (a authService) Register(req *api.AuthRequest) (*api.AuthResponse, error) {
	user := &repository.Auth{
		Email:        req.Email,
		PasswordHash: req.Password,
	}

	err := a.authRepo.Create(user)
	if err != nil {
		return nil, err
	}

	session, err := a.sessionService.CreateSession(user.ID)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

// ChangePassword requests a password change for a user. Link is sent to the user's email
func (a authService) ChangePassword(req *api.PasswordChangeRequest) error {
	user, err := a.authRepo.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	// Send email with link to change password
	// TODO: Implement email sending
	return a.authRepo.Update(user)
}

// ResetPassword resets the password for a user
func (a authService) ResetPassword(req *api.PasswordChange, token string) error {
	// TODO: Implement
	// Verify token

	// Find user by token

	// Update user password

	// Delete token
	return nil
}

// VerifyUser verifies a user
func (a authService) VerifyUser(token string) error {
	// TODO: Implement
	// Verify token

	// Find user by token

	// Update user

	// Delete token
	return nil
}
