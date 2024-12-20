package service

import (
	"auth/internal/api"
	"auth/internal/repository"
	"time"
)

type SessionService interface {
	// CreateSession creates a new session. Returns prepared response with token.
	CreateSession(userId int64) (api.AuthResponse, error)

	// ValidateSession validates a session. If the session is invalid, an error is returned
	ValidateSession(token string) error

	// GetUserID returns the user ID associated with a session
	GetUserID(token string) (int64, error)

	// DeleteSession deletes a session
	DeleteSession(token string) error

	// HardDeleteSession deletes all expired sessions
	HardDeleteSession() error
}

type sessionService struct {
	sessionRepo repository.SessionRepository
}

func NewSessionService(sessionRepo repository.SessionRepository) SessionService {
	return &sessionService{
		sessionRepo: sessionRepo,
	}
}

// CreateSession creates a new session
func (s sessionService) CreateSession(userId int64) (api.AuthResponse, error) {
	session := repository.NewSession(userId)
	err := s.sessionRepo.Create(session)
	if err != nil {
		return api.AuthResponse{}, err
	}
	return api.AuthResponse{Token: session.SessionKey}, nil
}

// ValidateSession validates a session. If the session is invalid, an error is returned
func (s sessionService) ValidateSession(token string) error {
	_, err := s.sessionRepo.Get(token)
	return err
}

// GetUserID returns the user ID associated with a session
func (s sessionService) GetUserID(token string) (int64, error) {
	session, err := s.sessionRepo.Get(token)
	if err != nil {
		return 0, err
	}
	return session.UserID, nil
}

// DeleteSession deletes a session
func (s sessionService) DeleteSession(token string) error {
	return s.sessionRepo.Delete(token)
}

// HardDeleteSession deletes all expired sessions
func (s sessionService) HardDeleteSession() error {
	sessions, err := s.sessionRepo.GetAll()
	if err != nil {
		return err
	}
	for _, session := range sessions {
		if session.ExpiresAt.Before(time.Now()) {
			err := s.sessionRepo.HardDelete(session.SessionKey)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
