package route

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/auth"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/repository"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/web/handler"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/web/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewMux(db *gorm.DB, secret string, time int) *gin.Engine {
	//jwt
	JWT := auth.NewJWT(secret, time)
	// middlware
	mw := middleware.NewFields(secret, time)
	// user
	userRepository := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepository, JWT)
	// coffee
	coffeeRepository := repository.NewCoffeeRepository(db)
	coffeeHandler := handler.NewCoffeeHandler(coffeeRepository, JWT)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//user
	r.GET("ping", handler.Status)
	r.POST("/user", userHandler.CreateUser)
	r.PUT("/user/:id", mw.ValidateToken(), userHandler.UpdateUser)
	r.POST("/login", userHandler.Login)
	//coffee
	r.POST("/coffee/", mw.ValidateTokenAdmin(), coffeeHandler.CreateCoffee)
	r.PUT("/coffee/:id", mw.ValidateTokenAdmin(), coffeeHandler.UpdateCoffee)
	r.DELETE("/coffee/", mw.ValidateTokenAdmin(), coffeeHandler.DeleteCoffee)
	r.GET("/coffee/:id", coffeeHandler.GetCoffee)
	r.GET("/coffee/", coffeeHandler.ListCoffees)
	return r
}
