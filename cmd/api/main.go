package main

import (
	"errors"
	"fmt"
	"github.com/LGuilhermeMoreira/loja-de-cafe/configuration"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/connection"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/migration"
	"github.com/LGuilhermeMoreira/loja-de-cafe/route"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	// image bucket
	err := os.MkdirAll("./images", 0755)
	if err != nil {
		panic(errors.New("error creating image bucket"))
	}
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}
	envData := configuration.NewEnvData()
	db, err := connection.NewConnection(envData)
	if err != nil {
		panic(err)
	}
	err = migration.Migration(db)
	if err != nil {
		panic(errors.New("error running migrations"))
	}
	mux := route.NewMux(db, envData.JWTSecret, envData.JWTTime)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", envData.Port),
		Handler: mux,
	}
	err = srv.ListenAndServe()
	panic(err)
}
