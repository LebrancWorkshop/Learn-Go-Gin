package main

import (
	"gogin-postgres/controller"
	"gogin-postgres/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, videoController.Save(c)) 
	})

	server.Run(":7901") 

}