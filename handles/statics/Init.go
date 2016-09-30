package statics

import (
	"github.com/gin-gonic/gin"
)

func InitStaticsRoutes(g *gin.Engine) {

}
func DefaultHandle(c *gin.Context) {
	c.Writer.Write([]byte("user "))
}
