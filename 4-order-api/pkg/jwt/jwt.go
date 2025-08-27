package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret string
}

type JWTData struct {
	Phone string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		secret: secret,
	}
}

func (j *JWT) Create(phone string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": phone,
	})

	s, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *JWT) Parse(tokenStr string) (bool, *JWTData) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return false, nil
	}

	phone := token.Claims.(jwt.MapClaims)["phone"]
	return token.Valid, &JWTData{
		Phone: phone.(string),
	}
}
