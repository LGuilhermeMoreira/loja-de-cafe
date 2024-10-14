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
	f := middleware.NewFields(secret, time)
	// user
	userRepository := repository.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepository, JWT)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(f.LogRequest())
	r.GET("ping", handler.Status)
	r.POST("/user", userHandler.CreateUser)
	r.PUT("/user/:id", f.ValidateToken(), userHandler.UpdateUser)
	r.POST("/login", userHandler.Login)
	return r
}
