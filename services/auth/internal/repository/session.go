package repository

import (
	"auth/pkg/utils"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"gorm.io/gorm"
	"time"
)

const (
	// SessionTTL represents the time to live for the session in seconds. By default, it is set to 1 year.
	SessionTTL = 60 * 60 * 24 * 365
)

// Session represents a session in the database
type Session struct {
	SessionKey string    `json:"session_key" gorm:"primaryKey" gorm:"column:session_key"`
	UserID     int64     `json:"user_id" gorm:"column:user_id"`
	LastUsed   time.Time `json:"last_used" gorm:"column:last_used"`
	ExpiresAt  time.Time `json:"expires_at" gorm:"column:expires_at"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" gorm:"autoUpdateTime"`
}

func (Session) TableName() string {
	return "sessions"
}

func NewSession(userID int64) *Session {

	return &Session{
		SessionKey: utils.GenerateRandomString(64),
		UserID:     userID,
		LastUsed:   time.Unix(0, 0),
		ExpiresAt:  time.Now().Add(time.Second * SessionTTL),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// SessionRepository represents the repository for the session
type SessionRepository interface {
	Create(session *Session) error
	GetAll() ([]*Session, error)
	Get(sessionKey string) (*Session, error)
	UpdateLastUsed(sessionKey string) error
	Delete(sessionKey string) error
	HardDelete(sessionKey string) error
	HardDeleteAllExpired() error
	HardDeleteAllInactive() error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (s sessionRepository) GetAll() ([]*Session, error) {
	logging.Logger.Debug("Getting all sessions")
	var sessions []*Session
	err := s.db.Find(&sessions).Error
	if err != nil {
		logging.Logger.Error("Failed to get all sessions: ", err)
		return nil, err
	}
	logging.Logger.Debug("Found ", len(sessions), " sessions")
	return sessions, nil
}

func (s sessionRepository) Create(session *Session) error {
	logging.Logger.Debug("Creating session with key: ", session.SessionKey[:5], "...")
	return s.db.Create(session).Error
}

func (s sessionRepository) Get(sessionKey string) (*Session, error) {
	logging.Logger.Debug("Getting session with key: ", sessionKey[:5], "...")
	var session Session
	err := s.db.Where("session_key = ?", sessionKey).Where("expires_at > ?", time.Now()).First(&session).Error
	if err != nil {
		logging.Logger.Error("Failed to get session with key: ", sessionKey[:5], "... - ", err)
		return nil, err
	}
	logging.Logger.Debug("Session found with key: ", sessionKey[:5], "...")
	return &session, nil
}

func (s sessionRepository) UpdateLastUsed(session string) error {
	logging.Logger.Debug("Updating last used time for session with key: ", session[:5], "...")
	return s.db.Model(&Session{}).Where("session_key = ?", session).Update("last_used", time.Now()).Error
}

func (s sessionRepository) Delete(sessionKey string) error {
	logging.Logger.Debug("Expiring session with key: ", sessionKey[:5], "...")
	return s.db.Model(&Session{}).Where("session_key = ?", sessionKey).Update("expires_at", time.Now()).Error
}

func (s sessionRepository) HardDelete(sessionKey string) error {
	logging.Logger.Debug("Deleting session with key: ", sessionKey[:5], "...")
	return s.db.Delete(&Session{}, "session_key = ?", sessionKey).Error
}

func (s sessionRepository) HardDeleteAllExpired() error {
	logging.Logger.Debug("Deleting all expired sessions...")
	return s.db.Delete(&Session{}, "expires_at < ?", time.Now()).Error
}

func (s sessionRepository) HardDeleteAllInactive() error {
	logging.Logger.Debug("Deleting all inactive sessions...")
	return s.db.Delete(&Session{}, "last_used < ?", time.Now().Add(-time.Second*SessionTTL)).Error
}
