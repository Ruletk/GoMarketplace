package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	secret := "testsecret"
	algo := jwt.SigningMethodHS256
	service := NewJwtService(algo, secret)

	payload := jwt.MapClaims{
		"userId": 12345,
		"role":   "user",
	}
	expires := int64(3600) // 1 hour

	token, err := service.GenerateToken(payload, expires)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, float64(12345), claims["userId"])
	assert.Equal(t, "user", claims["role"])
	assert.WithinDuration(t, time.Now().Add(time.Second*time.Duration(expires)), time.Unix(int64(claims["exp"].(float64)), 0), time.Minute)
	assert.WithinDuration(t, time.Now(), time.Unix(int64(claims["iat"].(float64)), 0), time.Minute)
	assert.WithinDuration(t, time.Now(), time.Unix(int64(claims["nbf"].(float64)), 0), time.Minute)
}

func TestGenerateVerificationToken(t *testing.T) {
	secret := "testsecret"
	algo := jwt.SigningMethodHS256
	service := NewJwtService(algo, secret)

	userId := int64(12345)

	token, err := service.GenerateVerificationToken(userId)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, float64(userId), claims["userId"])
	assert.Equal(t, "verification", claims["type"])
	assert.WithinDuration(t, time.Now().Add(time.Hour*24*7), time.Unix(int64(claims["exp"].(float64)), 0), time.Minute)
	assert.WithinDuration(t, time.Now(), time.Unix(int64(claims["iat"].(float64)), 0), time.Minute)
	assert.WithinDuration(t, time.Now(), time.Unix(int64(claims["nbf"].(float64)), 0), time.Minute)
}

func TestGeneratePasswordResetToken(t *testing.T) {
	secret := "testsecret"
	algo := jwt.SigningMethodHS256
	service := NewJwtService(algo, secret)

	userId := int64(12345)

	token, err := service.GeneratePasswordResetToken(userId)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	fmt.Println(token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, float64(userId), claims["userId"])
	assert.Equal(t, "password_reset", claims["type"])
	assert.WithinDuration(t, time.Now().Add(time.Hour*24), time.Unix(int64(claims["exp"].(float64)), 0), time.Minute)
	assert.WithinDuration(t, time.Now(), time.Unix(int64(claims["iat"].(float64)), 0), time.Minute)
	assert.WithinDuration(t, time.Now(), time.Unix(int64(claims["nbf"].(float64)), 0), time.Minute)
}

func TestParseToken(t *testing.T) {
	secret := "testsecret"
	algo := jwt.SigningMethodHS256
	service := NewJwtService(algo, secret)

	payload := jwt.MapClaims{
		"userId": 12345,
		"role":   "user",
	}
	expires := int64(3600) // 1 hour

	token, err := service.GenerateToken(payload, expires)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := service.ParseToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, float64(12345), claims["userId"])
	assert.Equal(t, "user", claims["role"])
}

func TestParseToken_InvalidToken(t *testing.T) {
	secret := "testsecret"
	algo := jwt.SigningMethodHS256
	service := NewJwtService(algo, secret)

	invalidToken := "invalid.token.string"

	claims, err := service.ParseToken(invalidToken)
	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestParseToken_InvalidClaims(t *testing.T) {
	secret := "testsecret"
	algo := jwt.SigningMethodHS256
	service := NewJwtService(algo, secret)

	// Create a token with invalid claims
	token := jwt.NewWithClaims(algo, jwt.MapClaims{})
	tokenString, err := token.SignedString([]byte(secret))
	assert.NoError(t, err)

	// No credentials jwt is still jwt but with no claims
	claims, err := service.ParseToken(tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
}

func TestIsVerificationToken(t *testing.T) {
	secret := "testsecret"
	algo := jwt.SigningMethodHS256
	service := NewJwtService(algo, secret)

	userId := int64(12345)

	// Generate a valid verification token
	token, err := service.GenerateVerificationToken(userId)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Test with a valid verification token
	isValid, parsedUserId := service.IsVerificationToken(token)
	assert.True(t, isValid)
	assert.Equal(t, userId, parsedUserId)

	// Test with an invalid token
	invalidToken := "invalid.token.string"
	isValid, parsedUserId = service.IsVerificationToken(invalidToken)
	assert.False(t, isValid)
	assert.Equal(t, int64(0), parsedUserId)

	// Generate a token with a different type
	payload := jwt.MapClaims{
		"userId": userId,
		"type":   "different_type",
	}
	differentToken, err := service.GenerateToken(payload, 3600)
	assert.NoError(t, err)
	assert.NotEmpty(t, differentToken)

	// Test with a token of a different type
	isValid, parsedUserId = service.IsVerificationToken(differentToken)
	assert.False(t, isValid)
	assert.Equal(t, int64(12345), parsedUserId)
}

func TestIsPasswordResetToken(t *testing.T) {
	secret := "testsecret"
	algo := jwt.SigningMethodHS256
	service := NewJwtService(algo, secret)

	userId := int64(12345)

	// Generate a valid password reset token
	token, err := service.GeneratePasswordResetToken(userId)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Test with a valid password reset token
	isValid, parsedUserId := service.IsPasswordResetToken(token)
	assert.True(t, isValid)
	assert.Equal(t, userId, parsedUserId)

	// Test with an invalid token
	invalidToken := "invalid.token.string"
	isValid, parsedUserId = service.IsPasswordResetToken(invalidToken)
	assert.False(t, isValid)
	assert.Equal(t, int64(0), parsedUserId)

	// Generate a token with a different type
	payload := jwt.MapClaims{
		"userId": userId,
		"type":   "different_type",
	}
	differentToken, err := service.GenerateToken(payload, 3600)
	assert.NoError(t, err)
	assert.NotEmpty(t, differentToken)

	// Test with a token of a different type
	isValid, parsedUserId = service.IsPasswordResetToken(differentToken)
	assert.False(t, isValid)
	assert.Equal(t, int64(12345), parsedUserId)
}
