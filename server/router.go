package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mesuutt/claps/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		claps := new(controllers.ClapsController)
		v1.POST("/add", claps.Add)
	}

	return router
}
