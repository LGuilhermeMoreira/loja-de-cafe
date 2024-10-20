package repository

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/model"
	"github.com/LGuilhermeMoreira/loja-de-cafe/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coffee struct {
	DB *gorm.DB
}

func NewCoffeeRepository(db *gorm.DB) *Coffee {
	return &Coffee{DB: db}
}

func (c *Coffee) isValid(id uuid.UUID) (*model.Coffee, error) {
	var coffee model.Coffee
	err := c.DB.Find(&coffee, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &coffee, nil
}

func (c *Coffee) Create(input dto.InputCreateCoffee, path string) (*dto.OutputCoffee, error) {
	coffee := model.NewCoffee(input.Prince, input.Description, path, input.Name)
	err := c.DB.Create(&coffee).Error
	if err != nil {
		return nil, err
	}
	response := dto.OutputCoffee{
		Name:        coffee.Name,
		Id:          coffee.ID.String(),
		Description: coffee.Description,
		Prince:      coffee.Price,
		Data:        input.Data,
	}
	return &response, nil
}

func (c *Coffee) Update(input dto.InputUpdateCoffee, id uuid.UUID) (*dto.OutputCoffee, error) {
	coffee, err := c.isValid(id)
	if err != nil {
		return nil, err
	}

	coffee.Price = input.Prince
	coffee.Name = input.Name
	coffee.Description = input.Description

	err = c.DB.Save(coffee).Error
	if err != nil {
		return nil, err
	}
	response := dto.OutputCoffee{
		Name:        coffee.Name,
		Id:          coffee.ID.String(),
		Description: coffee.Description,
		Prince:      coffee.Price,
		Data:        utils.GetBase64(coffee.ImagePath),
	}
	return &response, nil
}

func (c *Coffee) Delete(id uuid.UUID) error {
	if _, err := c.isValid(id); err != nil {
		return err
	}
	return c.DB.Delete(&model.Coffee{}, id).Error
}

func (c *Coffee) FindAll(pagination, limit int, sort string) ([]dto.OutputCoffee, error) {
	var coffees []model.Coffee
	err := c.DB.Find(&coffees).Offset((pagination - 1) * limit).Order("created_at" + sort).Error
	if err != nil {
		return nil, err
	}
	response := make([]dto.OutputCoffee, len(coffees))
	for i, c := range coffees {
		response[i] = dto.OutputCoffee{
			Name:        c.Name,
			Id:          c.ID.String(),
			Description: c.Description,
			Prince:      c.Price,
			Data:        utils.GetBase64(c.ImagePath),
		}
	}
	return response, nil
}
func (c *Coffee) FindById(id uuid.UUID) (*dto.OutputCoffee, error) {
	coffee, err := c.isValid(id)
	if err != nil {
		return nil, err
	}
	response := dto.OutputCoffee{
		Name:        coffee.Name,
		Id:          coffee.ID.String(),
		Description: coffee.Description,
		Prince:      coffee.Price,
		Data:        utils.GetBase64(coffee.ImagePath),
	}
	return &response, nil
}
