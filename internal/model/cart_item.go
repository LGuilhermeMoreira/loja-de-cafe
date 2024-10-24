package model

import (
	"github.com/google/uuid"
	"time"
)

type CartItem struct {
	Id             uuid.UUID    `gorm:"primaryKey;unique"`
	ShoppingCart   ShoppingCart `gorm:"references:Id"`
	ShoppingCartId uuid.UUID
	CoffeeId       uuid.UUID
	Coffee         Coffee `gorm:"references:Id"`
	Price          float64
	Quantity       int `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewCartItem(shoppingCartId, coffeeId uuid.UUID, quantity int, price float64) *CartItem {
	return &CartItem{
		Id:             uuid.New(),
		ShoppingCartId: shoppingCartId,
		CoffeeId:       coffeeId,
		Quantity:       quantity,
		Price:          price * float64(quantity),
	}
}
