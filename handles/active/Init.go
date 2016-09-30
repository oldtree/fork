package active

import (
	"github.com/gin-gonic/gin"
)

func InitActiveRouters(g *gin.Engine) {

}

func DefaultHandle(c *gin.Context) {
	c.Writer.Write([]byte("user "))
}
