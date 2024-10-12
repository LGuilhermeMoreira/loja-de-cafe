package handler

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/auth"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	_interface "github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/interface"
)

type UserHandler struct {
	_interface.UserInterface
	auth.JWTInterface[dto.OutputUser]
}

func NewUserHandler(userInterface _interface.UserInterface, jwtInterface auth.JWTInterface[dto.OutputUser]) *UserHandler {
	return &UserHandler{
		UserInterface: userInterface,
		JWTInterface:  jwtInterface,
	}
}

func (u *UserHandler) CreateUser(t any) {
	// todo
}
