package routes

import (
	controllers "app/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	bookController := controllers.NewBookController()

	v1 := router.Group("/v1")
	{
		v1.GET("/books", bookController.FindAll)
		v1.POST("/book", bookController.Create)
		v1.DELETE("/book/:id", bookController.Delete)
	}

	return v1
}
