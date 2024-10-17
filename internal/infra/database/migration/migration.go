package migration

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/model"
	"gorm.io/gorm"
)

var models = []interface{}{
	model.User{},
	model.Coffee{},
}

func Migration(db *gorm.DB) error {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}
