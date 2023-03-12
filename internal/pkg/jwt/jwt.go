package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID uint64
	jwt.RegisteredClaims
}

type Manager struct {
	key string
}

func NewManager(key string) *Manager {
	return &Manager{
		key: key,
	}
}

func (m *Manager) Sign(claims *CustomClaims) (string, error) {
	claims.ID = fmt.Sprintf("%s-%d", claims.UserID, time.Now().UnixNano())
	claims.IssuedAt = jwt.NewNumericDate(time.Now().UTC())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * 12))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(m.key))
}

func (m *Manager) Validate(raw string) error {
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(raw, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.key), nil
	})
	if err != nil {
		return fmt.Errorf("jwt.ParseWithClaims failed: %w", err)
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
