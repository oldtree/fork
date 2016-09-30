package auth

import (
	"fork/service/accountService"

	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func RegisterGetHandle(c *gin.Context) {
	c.JSON(200, gin.H{"errorcode": 200, "desc": "ok", "data": nil})
	return
}

func RegisterPostHandle(c *gin.Context) {
	auth := new(accountService.Auth)
	password := c.PostForm("password")
	account := c.PostForm("account")
	nickname := c.PostForm("nickname")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{"errorcode": 200, "desc": "account not set", "data": nil})
		return
	}
	if password == "" {
		c.JSON(http.StatusOK, gin.H{"errorcode": 200, "desc": "password not set", "data": nil})
		return
	}
	if nickname == "" {
		c.JSON(http.StatusOK, gin.H{"errorcode": 200, "desc": "nickname not set", "data": nil})
	}

	ok, status, err := auth.Register(account, password, nickname)
	if ok {
		c.SetCookie("fork", nickname, int(time.Second*60*60*24*2), "/fork/"+nickname, "www.forker.com", true, true)
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.JSON(http.StatusOK, gin.H{"errorcode": status, "desc": err.Error(), "data": nil})
	return
}
