package repository

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShoppingCartRespository struct {
	DB *gorm.DB
}

func NewShoppingCartRespository(db *gorm.DB) *ShoppingCartRespository {
	return &ShoppingCartRespository{
		DB: db,
	}
}

func (s *ShoppingCartRespository) AddItem(id uuid.UUID, input dto.InputAddItemShoppingCartDto) (*dto.OutputCartItemDto, error) {
	shoppingCartModel, err := s.isValid(id)
	if err != nil {
		return nil, err
	}
	parse, err := uuid.Parse(input.CoffeeID)
	if err != nil {
		return nil, err
	}
	cartItem := model.NewCartItem(shoppingCartModel.Id, parse, input.Quantity, input.Price)
	s.DB.Create(&cartItem)
	return nil, nil
}

func (s *ShoppingCartRespository) isValid(uuid uuid.UUID) (*model.ShoppingCart, error) {
	return nil, nil
}

func (s *ShoppingCartRespository) getAllItems(id uuid.UUID) error {
	return nil
}
