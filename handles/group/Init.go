package group

import (
	"github.com/gin-gonic/gin"
)

func InitGroupRoutes(g *gin.Engine) {

}
func DefaultHandle(c *gin.Context) {
	c.Writer.Write([]byte("user "))
}
