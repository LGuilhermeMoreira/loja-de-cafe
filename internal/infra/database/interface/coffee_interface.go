package _interface

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	"github.com/google/uuid"
)

type CoffeeInterface interface {
	FindById(id uuid.UUID) (*dto.OutputCoffee, error)
	FindAll(pagination, limit int, sort string) ([]dto.OutputCoffee, error)
	Create(input dto.InputCreateCoffee) (*dto.OutputCoffee, error)
	Update(input dto.InputUpdateCoffee, id uuid.UUID) (*dto.OutputCoffee, error)
	Delete(id uuid.UUID) error
}
