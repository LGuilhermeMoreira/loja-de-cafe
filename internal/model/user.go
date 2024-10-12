package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Admin     bool
	Id        uuid.UUID `gorm:"primaryKey;unique"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email, password string, admin bool) (*User, error) {
	id := uuid.New()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		Admin:    admin,
		Id:       id,
		Name:     name,
		Email:    email,
		Password: string(hashPassword),
	}, nil
}
