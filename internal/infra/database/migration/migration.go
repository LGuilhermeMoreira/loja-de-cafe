package migration

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/model"
	"gorm.io/gorm"
)

var models = []interface{}{
	model.User{},
	model.Coffee{},
	model.CartItem{},
	model.ShoppingCart{},
}

func Migration(db *gorm.DB) error {
	for _, m := range models {
		err := db.AutoMigrate(m)
		if err != nil {
			return err
		}
	}
	return nil
}
