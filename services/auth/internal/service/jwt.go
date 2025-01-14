package service

import "github.com/golang-jwt/jwt/v5"

type JwtService interface {
	GenerateToken(payload map[string]interface{}) (string, error)
	GenerateVerificationToken(userId int64) (token string, err error)
	GeneratePasswordResetToken(userId int64) (token string, err error)
	ParseToken(token string) (map[string]interface{}, error)
	IsVerificationToken(token string) (isValid bool, userId int64)
	IsPasswordResetToken(token string) (isValid bool, userId int64)
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

func (j jwtService) GenerateToken(payload map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(j.algo, jwt.MapClaims(payload))
	return token.SignedString([]byte(j.secret))
}

func (j jwtService) GenerateVerificationToken(userId int64) (token string, err error) {
	payload := map[string]interface{}{
		"userId": userId,
		"type":   "verification",
	}
	token, err = j.GenerateToken(payload)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j jwtService) GeneratePasswordResetToken(userId int64) (token string, err error) {
	payload := map[string]interface{}{
		"userId": userId,
		"type":   "password_reset",
	}
	token, err = j.GenerateToken(payload)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j jwtService) ParseToken(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, nil)
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return claims, nil
}

func (j jwtService) IsVerificationToken(token string) (isValid bool, userId int64) {
	claims, err := j.ParseToken(token)
	if err != nil {
		return false, 0
	}
	return claims["type"] == "verification", int64(claims["userId"].(float64))
}

func (j jwtService) IsPasswordResetToken(token string) (isValid bool, userId int64) {
	claims, err := j.ParseToken(token)
	if err != nil {
		return false, 0
	}
	return claims["type"] == "password_reset", int64(claims["userId"].(float64))
}
