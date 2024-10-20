package middleware

import (
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/auth"
	"github.com/LGuilhermeMoreira/loja-de-cafe/internal/dto"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type filds struct {
	jwt auth.JWTInterface[dto.OutputUser]
}

func NewFields(secret string, time int) *filds {
	return &filds{
		jwt: auth.NewJWT(secret, time),
	}
}

func (f filds) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token is missing"})
			return
		}
		tokenString := strings.Split(token, " ")[1]
		err := f.jwt.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}

func (f filds) ValidateTokenAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token is missing"})
			return
		}
		tokenString := strings.Split(token, " ")[1]
		err := f.jwt.ValidateTokenAdmin(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}

func (f filds) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("🆕%s ➡️%s 🌎%s\n", c.Request.Method, c.Request.RequestURI, c.ClientIP())
		c.Next()
	}
}
