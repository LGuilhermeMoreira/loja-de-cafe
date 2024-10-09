package connection

import (
	"errors"
	"fmt"
	"github.com/LGuilhermeMoreira/loja-de-cafe/configuration"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func NewConnection(envData configuration.EnvData) (*gorm.DB, error) {
	log.Println("trying to connect to database")
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
		envData.DBHost,
		envData.DBPort,
		envData.DBUser,
		envData.DBPass,
		envData.DBName,
	)
	log.Println("dns: ", dns)
	count := 0
	for {
		if count >= 5 {
			return nil, errors.New("try 5 times to connect to database and failed")
		}
		log.Println("time: ", count)
		db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err == nil {
			return db, nil
		}
		time.Sleep(2 * time.Second)
		count++
	}

	return nil, errors.New("could not connect to database")
}
