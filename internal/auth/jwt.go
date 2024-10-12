package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jWT[T Models] struct {
	secret  string
	timeout int
	Model   T
}

func NewJWT[T Models](secret string, timeout int, model T) *jWT[T] {
	return &jWT[T]{
		secret:  secret,
		timeout: timeout,
		Model:   model,
	}
}

func (j *jWT[T]) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": j.Model,
		"exp":  time.Now().Add(time.Minute * time.Duration(j.timeout)).Unix(),
	})

	return token.SignedString([]byte(j.secret))
}

func (j *jWT[T]) getSecret(t *jwt.Token) (any, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(j.secret), nil
}
func (j *jWT[T]) ValidateToken(tokenString string) error {
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
