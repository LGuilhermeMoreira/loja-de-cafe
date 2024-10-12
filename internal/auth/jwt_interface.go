package auth

import "github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"

type Models interface {
	dto.OutputUser
	// add outher implementation
}

type JWTInterface[T Models] interface {
	GenerateToken() (string, error)
	ValidateToken(token string) error
}
