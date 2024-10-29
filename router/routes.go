package router

import (
	"github.com/Devtino7/goapi/handler"
	"github.com/gin-gonic/gin"
)

func inicializeRoutes(router *gin.Engine) {
	// Inicialize Handler
	handler.InicializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}

}
