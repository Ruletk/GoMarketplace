package service

import (
	"auth/pkg/utils"
	"errors"
)

// TODO: Add redis to cache the token

const (
	TokenTypeVerification  = "verification"
	TokenTypePasswordReset = "password_reset"
)

var (
	ErrInvalidTokenType = errors.New("invalid token type")
)

type Token struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

type TokenService interface {
	// GenerateToken generates a new token for the user
	GenerateToken(userID int64, type_ string) error

	// ValidateToken validates a token and returns the user ID
	ValidateToken(token string, tokenType string) (int64, error)

	// DeleteToken deletes a token
	DeleteToken(token string) error
}

type tokenService struct {
}

func NewTokenService() TokenService {
	return &tokenService{}
}

func (t tokenService) GenerateToken(userID int64, type_ string) error {
	// TODO+REFACTOR: Make a better token type validation
	if type_ != TokenTypeVerification && type_ != TokenTypePasswordReset {
		return ErrInvalidTokenType
	}

	token := utils.GenerateRandomString(64)
	// TODO: Add redis to cache the token
	// Token will be mapped to the user ID and type
	tok := Token{
		Token: token,
		Type:  type_,
	}
	_ = tok
	return nil
}

func (t tokenService) ValidateToken(token string, tokenType string) (int64, error) {
	return 0, nil
}

func (t tokenService) DeleteToken(token string) error {
	return nil
}
