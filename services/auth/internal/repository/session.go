package repository

import (
	"auth/pkg/utils"
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
	UserID     int       `json:"user_id" gorm:"column:user_id"`
	LastUsed   time.Time `json:"last_used" gorm:"column:last_used"`
	ExpiresAt  time.Time `json:"expires_at" gorm:"column:expires_at"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" gorm:"autoUpdateTime"`
}

func NewSession(userID int) *Session {

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
	Get(sessionKey string) (*Session, error)
	Delete(sessionKey string) error
	HardDelete(sessionKey string) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository() SessionRepository {
	return &sessionRepository{}
}

func (s sessionRepository) Create(session *Session) error {
	return s.db.Create(session).Error
}

func (s sessionRepository) Get(sessionKey string) (*Session, error) {
	var session Session
	err := s.db.Where("session_key = ?", sessionKey).Where("expires_at > ?", time.Now()).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s sessionRepository) Delete(sessionKey string) error {
	return s.db.Update("expires_at", time.Now()).Where("session_key = ?", sessionKey).Error
}

func (s sessionRepository) HardDelete(sessionKey string) error {
	return s.db.Delete(&Session{}, "session_key = ?", sessionKey).Error
}
