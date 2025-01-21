package auth

import (
	"github.com/HarvinRaj/goldshop/internal/models"
	"github.com/HarvinRaj/goldshop/logger"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManager struct {
	SecretKey     string
	TokenDuration time.Duration
}

type Claims struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWT(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		SecretKey:     secretKey,
		TokenDuration: tokenDuration,
	}
}

func (j *JWTManager) GenerateToken(user *models.Users) (string, error) {

	claims := &Claims{
		UserID:   user.UserID,
		UserName: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.TokenDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		logger.ErrorLog.Auth.Printf("Signing token failed, %v", err)
		return "", err
	}

	return signedToken, nil
}

func (j *JWTManager) ValidateToken(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil || !token.Valid {
		logger.ErrorLog.Auth.Printf("jwt validate token error, %v", err)
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		logger.ErrorLog.Auth.Printf("claims error, %v", jwt.ErrTokenInvalidClaims)
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
