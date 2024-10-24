package model

import (
	"github.com/google/uuid"
	"time"
)

type ShoppingCart struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	User      User      `gorm:"references:Id"`
	UserId    uuid.UUID
	CartItems []CartItem
	CreatedAt time.Time
	UpdatedAt time.Time
	Price     float64
}

func NewShoppingCart(userId uuid.UUID) *ShoppingCart {
	return &ShoppingCart{
		Id:     uuid.New(),
		UserId: userId,
	}
}
