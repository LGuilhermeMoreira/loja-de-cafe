package route

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/auth"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/database/repository"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/infra/web/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewMux(db *gorm.DB, secret string, time int) *gin.Engine {
	//jwt
	JWT := auth.NewJWT(secret, time)

	// user
	userRepository := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepository, JWT)

	r := gin.Default()
	r.POST("/user", userHandler.CreateUser)
	return r
}
