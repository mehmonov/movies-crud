package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mehmonov/movies-crud/internal/models"
)

type JWTService struct {
	accessSecret  string
	refreshSecret string
}

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func NewJWTService(accessSecret, refreshSecret string) *JWTService {
	return &JWTService{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}

func (s *JWTService) GenerateTokenPair(userID uint) (*models.TokenPair, error) {
	accessToken, err := s.generateToken(userID, AccessToken, time.Minute*15)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateToken(userID, RefreshToken, time.Hour*24*7)
	if err != nil {
		return nil, err
	}

	return &models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *JWTService) generateToken(userID uint, tokenType TokenType, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    string(tokenType),
		"exp":     time.Now().Add(expiration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := s.accessSecret
	if tokenType == RefreshToken {
		secret = s.refreshSecret
	}

	return token.SignedString([]byte(secret))
}

func (s *JWTService) ValidateAccessToken(tokenString string) (uint, error) {
	return s.validateToken(tokenString, AccessToken)
}

func (s *JWTService) ValidateRefreshToken(tokenString string) (uint, error) {
	return s.validateToken(tokenString, RefreshToken)
}

func (s *JWTService) validateToken(tokenString string, tokenType TokenType) (uint, error) {
	secret := s.accessSecret
	if tokenType == RefreshToken {
		secret = s.refreshSecret
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check token type
		if tokenTypeStr, ok := claims["type"].(string); !ok || TokenType(tokenTypeStr) != tokenType {
			return 0, errors.New("invalid token type")
		}

		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, errors.New("invalid token")
}
