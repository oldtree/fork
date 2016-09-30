//用户的登陆，注册，关于，修改用户的相关信息的handle
package user

import (
	"fork/tools/log"

	"fork/service/accountService"

	"github.com/gin-gonic/gin"
)

func UserMainEntry(c *gin.Context) {
	uid := c.Query("uid")
	authToken := c.Request.Header.Get("authtoken")
	if uid == "" {
		c.Redirect(302, "baidu.com")
		return
	}
	if authToken == "" {
		c.Redirect(302, "/login")
		return
	}
	authcheck := new(accountService.Auth)
	ok, status, err := authcheck.CheckAuth(authToken)
	if !ok {
		c.Redirect(302, "/login")
		log.Error(err)
		return
	}

	if status != 0 {
		c.Redirect(302, "ilegal user access!")
		log.Error(err)
		return
	}
	// login ok
	//get user profile
	//get user avator
	//get user recent work
	//get user topic

}

func DefaultHandle(c *gin.Context) {
	c.Writer.Write([]byte("user "))
}
