package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWT[T Models] struct {
	Secret  string
	Timeout int
}

func NewJWT[T Models](secret string, timeout int) *JWT[T] {
	return &JWT[T]{
		Secret:  secret,
		Timeout: timeout,
	}
}

func (j *JWT[T]) GenerateToken(data T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Hour * time.Duration(j.Timeout)).Unix(),
	})

	return token.SignedString([]byte(j.Secret))
}

func (j *JWT[T]) getSecret(t *jwt.Token) (any, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(j.Secret), nil
}
func (j *JWT[T]) ValidateToken(tokenString string) error {
	tokenJWT, err := jwt.Parse(tokenString, j.getSecret)
	if err != nil {
		return err
	}
	if !tokenJWT.Valid {
		return jwt.ErrSignatureInvalid
	}
	claims, ok := tokenJWT.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("invalid expiration claim")
	}
	expTime := time.Unix(int64(exp), 0)
	if time.Now().After(expTime) {
		return errors.New("token expired")
	}
	return nil
}
