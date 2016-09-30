package auth

import (
	"github.com/gin-gonic/gin"
)

func InitAuthRouters(g *gin.Engine) {
	mainEntryGroup := g.Group("")
	{
		loginGroup := mainEntryGroup.Group("login")
		{
			loginGroup.GET("", LoginGetHandle)
			loginGroup.POST("", LoginPostHandle)
		}
		registerGroup := mainEntryGroup.Group("register")
		{
			registerGroup.GET("", DefaultHandle)
			registerGroup.POST("", DefaultHandle)
		}
	}
}

func DefaultHandle(c *gin.Context) {
	c.Writer.WriteString("auth")
}
