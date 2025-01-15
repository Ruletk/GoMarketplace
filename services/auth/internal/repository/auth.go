package repository

import (
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// Auth represents an authentication in the database
type Auth struct {
	ID           int64     `json:"id" gorm:"primaryKey" gorm:"column:id"`
	Email        string    `json:"email" gorm:"unique" gorm:"index" gorm:"column:email"`
	PasswordHash string    `json:"-" gorm:"column:password_hash"`
	Active       bool      `json:"active" gorm:"column:active" gorm:"default:true"`
	IsSeller     bool      `json:"is_seller" gorm:"column:is_seller" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at" gorm:"autoUpdateTime"`
	DeletedAt    time.Time `json:"delete_at" gorm:"column:delete_at"`
}

func (Auth) TableName() string {
	return "auth"
}

func (a Auth) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(password)) == nil
}

func (a Auth) GeneratePasswordHash(password string) (passwordHash string) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	logging.Logger.Debug("Generating password hash for user with email: ", a.Email, " Hash: ", string(pass))
	if err != nil {
		logging.Logger.Error("Failed to generate password hash: ", err)
		a.PasswordHash = ""
		return
	}
	return string(pass)
}

// AuthRepository represents the repository for the authentication
type AuthRepository interface {
	Create(auth *Auth) error
	GetByEmail(email string) (*Auth, error)
	GetByID(id int64) (*Auth, error)
	Update(auth *Auth) error
	Delete(id int64) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (a authRepository) Create(auth *Auth) error {
	logging.Logger.Debug("Creating user with email: ", auth.Email)
	return a.db.Create(auth).Error
}

func (a authRepository) GetByEmail(email string) (*Auth, error) {
	logging.Logger.Debug("Getting user by email: ", email)

	var auth Auth
	err := a.db.Where("email = ?", email).First(&auth).Error
	if err != nil {
		logging.Logger.Error("Failed to get user by email. ", err)
		return nil, err
	}
	logging.Logger.Debug("User found with email: ", email)
	return &auth, nil
}

func (a authRepository) GetByID(id int64) (*Auth, error) {
	logging.Logger.Debug("Getting user by ID: ", id)
	var auth Auth
	err := a.db.Where("id = ?", id).First(&auth).Error
	if err != nil {
		logging.Logger.Error("Failed to get user by ID: ", err)
		return nil, err
	}
	logging.Logger.Debug("User found with ID: ", id)
	return &auth, nil
}

func (a authRepository) Update(auth *Auth) error {
	logging.Logger.Debug("Updating user with ID: ", auth.ID)
	return a.db.Save(auth).Error
}

func (a authRepository) Delete(id int64) error {
	logging.Logger.Debug("Deleting user with ID: ", id)
	return a.db.Delete(&Auth{}, "id = ?", id).Error
}
