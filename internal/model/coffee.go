package model

import (
	"github.com/google/uuid"
	"time"
)

type Coffee struct {
	ID          uuid.UUID `gorm:"primaryKey;unique"`
	Name        string    `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	Description string    `gorm:"not null"`
	ImagePath   string    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCoffee(price float64, description, path, name string) *Coffee {
	return &Coffee{
		ID:          uuid.New(),
		Name:        name,
		Price:       price,
		Description: description,
		ImagePath:   path,
	}
}
