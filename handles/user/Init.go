package user

import (
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.Engine) {
	/*
	  https://fork.com/people/123123
	  https://fork.com/people/123123/setting
	  https://fork.com/people/123123/about
	  https://fork.com/people/123123/profile
	*/
	userGroup := g.Group("/user")
	{
		userGroup.GET("/:uid", UserMainEntry)
		userGroup.GET("/:uid/setting", DefaultHandle)
		userGroup.POST("/:uid/setting", DefaultHandle)
		activeGroup := userGroup.Group("/:uid/active")
		{
			activeGroup.GET("/:activetype")
			activeGroup.POST("/:activetype")
		}
	}
}
