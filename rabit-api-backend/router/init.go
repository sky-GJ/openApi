package router

import (
	"fmt"
	"rabit-api-backend/global"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

var indexGroup *gin.RouterGroup

func init() {
	r = gin.Default()

	// main router
	// r.Use(middleware.Cors, middleware.RefreshToken)
	indexGroup = r.Group("api")

	// children router
	UserRouter()
	InterfaceRouter()
}

func SetupGin() {
	addr := global.Config.Server.Http.Addr()

	if err := r.Run(addr); err != nil {
		fmt.Println("[ERROR] running err: " + err.Error())
		return
	}
}
