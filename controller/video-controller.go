package controller

import (
	"gogin-postgres/entity"
	"gogin-postgres/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(c *gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (controller *controller) Save(c *gin.Context) entity.Video {
	var video entity.Video
	c.BindJSON(&video)
	controller.service.Save(video) 
	return video
}

func (controller *controller) FindAll() []entity.Video {
	return controller.service.FindAll()	
}