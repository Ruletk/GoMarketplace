package service

// Tested and passed. 15.01.2025 Ruletk
// TODO: Move common functions to a separate shared package/library. Like GenerateToken, ValidateToken and so on.

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtService interface {
	GenerateToken(payload jwt.MapClaims, expires int64) (string, error)
	GenerateVerificationToken(userId int64) (token string, err error)
	GeneratePasswordResetToken(userId int64) (token string, err error)
	ParseToken(token string) (map[string]interface{}, error)
	IsVerificationToken(token string) (isValid bool, userId int64)
	IsPasswordResetToken(token string) (isValid bool, userId int64)
	DeleteToken(token string) error
	//GenerateAccessToken(payload map[string]interface{}) (string, error) // Maybe add this later
}

type jwtService struct {
	algo   jwt.SigningMethod
	secret string
}

func NewJwtService(algo jwt.SigningMethod, secret string) JwtService {
	return &jwtService{
		algo:   algo,
		secret: secret,
	}
}

// GenerateToken generates a new token for the user.
// The token will expire in time specified by the expires parameter.
// Expire is added to the payload, if current time is 1000 and expires is 100, the token will expire at 1100.
func (j jwtService) GenerateToken(payload jwt.MapClaims, expires int64) (string, error) {
	payload["exp"] = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expires)))
	payload["iat"] = jwt.NewNumericDate(time.Now())
	payload["nbf"] = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(j.algo, payload)
	return token.SignedString([]byte(j.secret))
}

// GenerateVerificationToken generates a new verification token for the user.
// The token will expire in 7 days.
// Can be checked with IsVerificationToken.
func (j jwtService) GenerateVerificationToken(userId int64) (token string, err error) {
	payload := jwt.MapClaims{
		"userId": userId,
		"type":   "verification",
	}
	token, err = j.GenerateToken(payload, 3600*24*7) // 7 days
	if err != nil {
		return "", err
	}
	return token, nil
}

// GeneratePasswordResetToken generates a new password reset token for the user.
// The token will expire in 1 day.
// Can be checked with IsPasswordResetToken.
func (j jwtService) GeneratePasswordResetToken(userId int64) (token string, err error) {
	payload := map[string]interface{}{
		"userId": userId,
		"type":   "password_reset",
	}
	token, err = j.GenerateToken(payload, 3600*24) // 1 day, for security reasons
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken parses a token and returns the claims.
// If the token is invalid, an error is returned.
func (j jwtService) ParseToken(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !j.CheckTokenNotDeleted(token) {
		return nil, jwt.ErrTokenExpired
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

// IsVerificationToken checks if a token is a verification token.
// If the token is invalid, false is returned.
// If the token is a verification token, true is returned and the userId is returned.
func (j jwtService) IsVerificationToken(token string) (isValid bool, userId int64) {
	claims, err := j.ParseToken(token)
	if err != nil {
		return false, 0
	}
	return claims["type"] == "verification", int64(claims["userId"].(float64))
}

// IsPasswordResetToken checks if a token is a password reset token.
// If the token is invalid, false is returned.
// If the token is a password reset token, true is returned and the userId is returned.
func (j jwtService) IsPasswordResetToken(token string) (isValid bool, userId int64) {
	claims, err := j.ParseToken(token)
	if err != nil {
		return false, 0
	}
	return claims["type"] == "password_reset", int64(claims["userId"].(float64))
}

// DeleteToken deletes a token from the system.
// This is useful when a token is no longer needed.
// In other words, marking a token as invalid.
func (j jwtService) DeleteToken(token string) error {
	// This is a dummy function, as we don't store tokens.
	// TODO: Make every token store in key-value storage, like Redis.
	//  Token that needs to be deleted will store in a blacklist.
	//  Token will be added with TTL, so it will be deleted automatically.
	//  TL;DR: Implement token blacklist with TTL.
	return nil
}

// CheckTokenNotDeleted checks if a token is not deleted.
// True is returned if the token is not deleted and is valid.
// False is returned if the token is deleted and cannot be used.
func (j jwtService) CheckTokenNotDeleted(token string) bool {
	// This is a dummy function, as we don't store tokens.
	// TODO: Read the comment in DeleteToken.
	return true
}
