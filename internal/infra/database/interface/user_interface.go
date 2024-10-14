package _interface

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
)

type UserInterface interface {
	Create(user dto.InputCreateUser) (*dto.OutputUser, error)
	Login(email string, password string) (*dto.OutputUser, error)
	Update(id string, input dto.InputUpdateUser) (*dto.OutputUser, error)
}
