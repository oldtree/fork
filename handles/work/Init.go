package work

import (
	"github.com/gin-gonic/gin"
)

func InitWorkRoutes(g *gin.Engine) {

}

func DefaultHandle(c *gin.Context) {
	c.Writer.Write([]byte("user "))
}
