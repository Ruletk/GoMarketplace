package service

import (
	"auth/internal/messages"
	"auth/internal/repository"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"gorm.io/gorm"
	"time"
)

type SessionService interface {
	// CreateSession creates a new session. Returns prepared response with token.
	CreateSession(userId int64) (messages.AuthResponse, error)

	// GetUserID returns the user ID associated with a session
	GetUserID(token string) (int64, error)

	// DeleteSession deletes a session
	DeleteSession(token string) error

	// HardDeleteSessions deletes all expired sessions. Admin method
	HardDeleteSessions() error

	// DeleteInactiveSessions deletes all sessions that are expired. Admin method
	DeleteInactiveSessions() error
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
func (s sessionService) CreateSession(userId int64) (messages.AuthResponse, error) {
	logging.Logger.Debug("Creating session for user with ID: ", userId)

	session := repository.NewSession(userId)
	err := s.sessionRepo.Create(session)

	if err != nil {
		logging.Logger.Error("Failed to create session: ", err)
		return messages.AuthResponse{}, err
	}
	logging.Logger.Debug("Session created with token: ", session.SessionKey[:5])
	return messages.AuthResponse{Token: session.SessionKey}, nil
}

func (s sessionService) GetSession(token string) (repository.Session, error) {
	session, err := s.sessionRepo.Get(token)
	if err != nil {
		return repository.Session{}, err
	}
	return *session, err
}

// GetUserID returns the user ID associated with a session
func (s sessionService) GetUserID(token string) (int64, error) {
	logging.Logger.Debug("Getting user ID for session with token: ", token[:5], "...")

	session, err := s.sessionRepo.Get(token)
	if err != nil {
		return 0, err
	}

	logging.Logger.Debug("User ID for session with token: ", token[:5], " is: ", session.UserID)
	return session.UserID, nil
}

// UpdateLastUsed updates the last used time of a session

// DeleteSession deletes a session
func (s sessionService) DeleteSession(token string) error {
	logging.Logger.Info("Deleting session with token: ", token[:5], "...")
	session, err := s.sessionRepo.Get(token)

	if err != nil {
		return err
	}

	if session.ExpiresAt.Before(time.Now()) {
		return gorm.ErrRecordNotFound
	}

	err = s.sessionRepo.Delete(token)
	if err != nil {
		logging.Logger.Error("Failed to delete session with token: ", token[:5], " - ", err)
	}
	return err
}

// HardDeleteSessions deletes all expired sessions
func (s sessionService) HardDeleteSessions() error {
	logging.Logger.Info("Deleting expired sessions...")

	return s.sessionRepo.HardDeleteAllExpired()
}

// DeleteInactiveSessions deletes all sessions that are expired
func (s sessionService) DeleteInactiveSessions() error {
	logging.Logger.Info("Deleting inactive sessions...")

	return s.sessionRepo.HardDeleteAllInactive()
}
