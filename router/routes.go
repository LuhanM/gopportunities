package router

import (
	docs "github.com/LuhanM/gopportunities/docs"
	"github.com/LuhanM/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	gswag "github.com/swaggo/gin-swagger"
)

func initilizeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	router.GET("/swagger/*any", gswag.WrapHandler(swagfiles.Handler))
}
