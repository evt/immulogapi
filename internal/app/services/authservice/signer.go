package authservice

import (
	"github.com/evt/immulogapi/internal/pkg/jwt"
)

type Signer interface {
	Sign(j *jwt.CustomClaims) (string, error)
}
