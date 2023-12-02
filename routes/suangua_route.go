package routes

import (
	"divination/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.Index)
	r.GET("/qigua", controllers.QiGua)
	return r
}
