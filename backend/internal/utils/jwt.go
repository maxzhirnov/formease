package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GenerateToken(userId string, email string) (string, error)
	ValidateToken(tokenString string) (string, string, error)
}

type JWTUtil struct {
	SecretKey []byte
}

func NewJWTUtil(secretKey string) *JWTUtil {
	return &JWTUtil{
		SecretKey: []byte(secretKey),
	}
}

func (j *JWTUtil) GenerateToken(userId string, email string, expirationHours int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Duration(expirationHours) * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	})
	return token.SignedString(j.SecretKey)
}

func (j *JWTUtil) ValidateToken(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})
	if err != nil {
		return "", "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", jwt.ErrSignatureInvalid
	}
	userId := claims["user_id"]
	email := claims["email"].(string)
	return userId.(string), email, nil
}

func (j *JWTUtil) GenerateRefreshToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(), // Refresh token expires in 7 days
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SecretKey)
}
