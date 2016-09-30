//默认的处理handle，暂时使用gin-http请求处理框架
package handles

import (
	"time"

	"github.com/gin-gonic/gin"
)

func DefaultHandle(c *gin.Context) {
	c.JSON(200, gin.H{"errorcode": -1, "data": nil, "description": "default hanldle"})
	return
}

func BadRequestHandle(c *gin.Context) {
	c.JSON(200, gin.H{"errorcode": -1, "data": nil, "description": "bad request"})
	return
}

func NotSupportHandle(c *gin.Context) {
	c.JSON(200, gin.H{"errorcode": -1, "data": nil, "description": "not support"})
	return
}

func AboutHandle(c *gin.Context) {
	data := make(map[string]string, 4)
	data["version"] = "0.0.1"
	data["name"] = "fork"
	data["owner"] = "grapes"
	data["location"] = "shenzhen"
	c.JSON(200, gin.H{"errorcode": 0, "data": nil, "description": "ok"})
	return
}

func MainHandle(c *gin.Context) {
	data := make(map[string]string, 4)
	data["title"] = "fork"
	data["time"] = time.Now().String()
	c.JSON(200, gin.H{"errorcode": 0, "data": data})
	return
}

func InitToolsRouters(g *gin.Engine) {
	enityRouters := g.Group("/")
	enityRouters.GET("about", AboutHandle)
	enityRouters.GET("notsupport", NotSupportHandle)
	enityRouters.GET("badrequest", BadRequestHandle)

}
