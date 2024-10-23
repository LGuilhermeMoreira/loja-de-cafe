package auth

import "github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"

type Models interface {
	dto.OutputUser
}

type JWTInterface[T Models] interface {
	GenerateToken(data T) (string, error)
	ValidateToken(token string) error
	ValidateTokenAdmin(token string) error
}
