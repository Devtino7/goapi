package router

import (
	"github.com/gin-gonic/gin"
)

func Inicialize() {
	// Inicialize router
	router := gin.Default()
	// Inicialize routes
	inicializeRoutes(router)
	// Run the server
	router.Run()
}
