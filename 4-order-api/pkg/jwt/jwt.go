package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: secret,
	}
}

func (j *JWT) Create(sessionId string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"session": sessionId,
	})

	s, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return s, nil
}
