package auth

import (
	"fork/service/accountService"
	"fork/tools/log"

	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func LoginGetHandle(c *gin.Context) {
	c.JSON(200, gin.H{"errorcode": 0, "desc": "ok", "data": nil})
	return
}

func LoginPostHandle(c *gin.Context) {
	auth := new(accountService.Auth)
	password := c.PostForm("password")
	account := c.PostForm("email")
	if password == "" || account == "" {
		c.JSON(http.StatusOK, gin.H{"errorcode": 0, "desc": "params error", "data": nil})
		return
	}
	ok, status, err := auth.Login(account, password)
	if ok {
		c.SetCookie("fork", auth.Nickname, int(time.Second*60*60*24*2), "/fork/"+auth.Nickname, "www.forker.com", true, true)
		c.Redirect(302, "/user/"+auth.Uid)
		return
	}
	c.JSON(http.StatusOK, gin.H{"errorcode": status, "desc": err.Error(), "data": nil})
	return
}

func LogoutPostHanle(c *gin.Context) {
	cookie, err := c.Request.Cookie("fork")
	if err != nil {
		log.Info(err)
	}
	content := cookie.Value
	log.Info(content)
	//todo clear cache data
	c.JSON(http.StatusOK, gin.H{"errorcode": 0, "desc": "logout success", "data": nil})
	return
}
