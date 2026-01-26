package services

import (
	"fmt"
	"rest-api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtService struct {
	secretKey []byte
}

type JwtData struct {
	UserId   [16]byte
	UserRole domain.UserRole
}

func NewJwtService(secret string) *JwtService {
	return &JwtService{
		secretKey: []byte(secret),
	}
}

func (j *JwtService) GenerateCode(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"role":   user.Role,
		"ttl":    time.Now().Add(time.Hour * 24 * 100).Unix(),
	})

	tokenString, err := token.SignedString(j.secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (j *JwtService) ValidateCode(tokenStr string) (JwtData, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		return JwtData{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return JwtData{}, fmt.Errorf("Error getting JWT")
	}

	if claims["ttl"].(float64) < float64(time.Now().Unix()) {
		return JwtData{}, fmt.Errorf("Expired token")
	}

	idStr, ok := claims["userID"].(string)
	if !ok {
		return JwtData{}, fmt.Errorf("userID claim is not a string")
	}

	userID, err := uuid.Parse(idStr)
	if err != nil {
		return JwtData{}, fmt.Errorf("invalid uuid format in token: %w", err)
	}

	role, ok := claims["role"].(string)
	if !ok {
		return JwtData{}, fmt.Errorf("User role is not a string")
	}

	userRole, err := domain.CastUserRole(role)

	if err != nil {
		return JwtData{}, fmt.Errorf("Error casting user role")
	}

	return JwtData{UserId: userID, UserRole: userRole}, nil
}
