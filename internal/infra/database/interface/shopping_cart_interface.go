package _interface

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	"github.com/google/uuid"
)

type ShoppingCartInterface interface {
	AddItem(uuid uuid.UUID, input dto.InputAddItemShoppingCartDto) (*dto.OutputShoppingCartDto, error)
}
