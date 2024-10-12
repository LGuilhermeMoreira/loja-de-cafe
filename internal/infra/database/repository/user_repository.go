package repository

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(input dto.InputCreateUser) (*dto.OutputUser, error) {
	user, err := model.NewUser(input.Name, input.Email, input.Password, false)
	if err != nil {
		return nil, err
	}
	err = u.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	var output dto.OutputUser
	output.Email = user.Email
	output.Name = user.Name
	output.ID = user.Id
	return &output, nil
}

func (u *User) Login(email, password string) (bool, error) {
	var user model.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil, nil
}
