package routes

import (
	"github.com/avner.oliveira/pdi-api/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		user := main.Group("user")
		{
			user.GET("/:id", controllers.FindById)
			user.POST("/", controllers.Createuser)
		}
	}

	return router
}
