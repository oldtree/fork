// fork project main.go
package main

import (
	"fork/handles/active"
	"fork/handles/auth"
	"fork/handles/group"
	"fork/handles/statics"
	"fork/handles/user"
	"fork/handles/work"

	"fork/tools/log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	g := gin.Default()
	log.Info("start init static files")
	//g.Static("/static", "./public/static")
	//g.LoadHTMLGlob("static/html/*")
	log.Info("end init static files")
	log.Info("start init router")
	user.InitUserRouter(g)
	auth.InitAuthRouters(g)
	active.InitActiveRouters(g)
	group.InitGroupRoutes(g)
	work.InitWorkRoutes(g)
	statics.InitStaticsRoutes(g)
	log.Info("end init router")
	g.Run("localhost:9000")
}
