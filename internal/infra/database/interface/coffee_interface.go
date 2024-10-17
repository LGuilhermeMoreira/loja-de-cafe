package _interface

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	"github.com/google/uuid"
)

type CoffeeInterface interface {
	FindById(id uuid.UUID) (*dto.OutputCoffee, error)
	FindAll(pagination int, sort string) ([]dto.OutputCoffee, error)
	Create(input dto.InputCreateCoffee, path string) (*dto.OutputCoffee, error)
	Update(input dto.InputUpdateCoffee) (*dto.OutputCoffee, error)
	Delete(id uuid.UUID) error
}
