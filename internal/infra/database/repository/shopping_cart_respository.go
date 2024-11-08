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

func (s *ShoppingCartRespository) Create(id uuid.UUID) (uuid.UUID, error) {
	shoppingCart := model.NewShoppingCart(id)
	return shoppingCart.Id, s.DB.Create(&shoppingCart).Error
}

func (s *ShoppingCartRespository) AddItem(id uuid.UUID, input dto.InputAddItemShoppingCartDto) (*dto.OutputShoppingCartDto, error) {
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
	cartItems, err := s.getAllItems(id)
	if err != nil {
		return nil, err
	}
	var response map[string]dto.OutputCartItemDto
	var totalPrice float64 = 0
	for _, cartItem := range cartItems {
		response[cartItem.Id.String()] = dto.OutputCartItemDto{
			CoffeeID: cartItem.CoffeeId.String(),
			Quantity: cartItem.Quantity,
			Price:    cartItem.Price,
		}
		totalPrice += cartItem.Price
	}
	return &dto.OutputShoppingCartDto{
		CartItemList: response,
		TotalPrice:   totalPrice,
	}, nil
}

func (s *ShoppingCartRespository) isValid(uuid uuid.UUID) (*model.ShoppingCart, error) {
	var shoppingCart model.ShoppingCart
	err := s.DB.Where("id = ?", uuid).First(&shoppingCart).Error
	if err != nil {
		return nil, err
	}
	return &shoppingCart, nil
}

func (s *ShoppingCartRespository) getAllItems(shoppingCartId uuid.UUID) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	if err := s.DB.Find(cartItems, shoppingCartId).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}
