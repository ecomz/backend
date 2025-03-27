package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	jwt.RegisteredClaims
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewClaims(id, name, email, issuer string, duration time.Time) MyClaims {
	return MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(duration),
		},
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func GenerateToken(claim MyClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(secret))
}

func ParseToken(accessToken, secret string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, err
	}

	claims, ok := token.Claims.(*MyClaims)
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
