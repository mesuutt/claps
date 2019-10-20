package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mesuutt/claps/config"
	"github.com/mesuutt/claps/controllers"
)

func NewRouter() *gin.Engine {
	config := config.GetConfig()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default()) // Allow all origins

	router.Use(static.Serve("/assets", static.LocalFile(config.GetString("asset_path"), false)))

	v1 := router.Group("v1")
	{
		claps := new(controllers.ClapsController)
		v1.POST("/claps/add", claps.Add)
		v1.POST("/claps/count", claps.Count)

		like := new(controllers.LikeController)
		v1.POST("/likes/increase", like.Increase)
		v1.POST("/likes/decrease", like.Decrease)
		v1.POST("/likes/count", like.Count)
	}

	return router
}
