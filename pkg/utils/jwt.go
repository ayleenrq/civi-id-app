package utils

import (
	"civi-id-app/configs"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, roleID int) (string, error) {
	claims := JWTClaims{
		UserID: fmt.Sprint(userID),
		Role:   fmt.Sprint(roleID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.GetJWTSecret()))
}

func ParseToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(configs.GetJWTSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GetExpiryFromToken(tokenStr string) (time.Time, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(configs.GetJWTSecret()), nil
	})

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims.ExpiresAt.Time, nil
	}

	return time.Time{}, err
}
