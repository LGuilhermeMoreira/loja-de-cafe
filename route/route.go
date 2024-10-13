package route

import "github.com/gin-gonic/gin"

func NewMux() *gin.Engine {
	r := gin.Default()
	r.POST()
}
