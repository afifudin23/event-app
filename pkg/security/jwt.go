package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateToken(UserID string, secretKey string) *string {
	claims := JWTClaims{
		UID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secretKey))
	return &tokenString
}

func VerifyToken(tokenString, secret string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("Invalid or expired token")
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token claims")
}
