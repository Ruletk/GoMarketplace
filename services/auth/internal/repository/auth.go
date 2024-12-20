package repository

import (
	"gorm.io/gorm"
	"time"
)

// Auth represents an authentication in the database
type Auth struct {
	ID           int       `json:"id" gorm:"primaryKey" gorm:"column:id"`
	Email        string    `json:"email" gorm:"unique" gorm:"index" gorm:"column:email"`
	PasswordHash string    `json:"password_hash" gorm:"column:password_hash"`
	Active       bool      `json:"active" gorm:"column:active" gorm:"default:true"`
	IsSeller     bool      `json:"is_seller" gorm:"column:is_seller" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at" gorm:"autoUpdateTime"`
	DeletedAt    time.Time `json:"delete_at" gorm:"column:delete_at"`
}

// AuthRepository represents the repository for the authentication
type AuthRepository interface {
	Create(auth *Auth) error
	GetByEmail(email string) (*Auth, error)
	GetByID(id int) (*Auth, error)
	Update(auth *Auth) error
	Delete(id int) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (a authRepository) Create(auth *Auth) error {
	return a.db.Create(auth).Error
}

func (a authRepository) GetByEmail(email string) (*Auth, error) {
	var auth Auth
	err := a.db.Where("email = ?", email).First(&auth).Error
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func (a authRepository) GetByID(id int) (*Auth, error) {
	var auth Auth
	err := a.db.Where("id = ?", id).First(&auth).Error
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func (a authRepository) Update(auth *Auth) error {
	return a.db.Save(auth).Error
}

func (a authRepository) Delete(id int) error {
	return a.db.Delete(&Auth{}, "id = ?", id).Error
}
