package controller

import (
	"fmt"
	"github.com/KodaiD/rest/service"
	"github.com/gin-gonic/gin"
)

func (pc Controller) PostIndex(c *gin.Context) {
	var s service.Service
	p, err := s.GetPostsAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) PostCreate(c *gin.Context) {
	var s service.Service
	p, err := s.CreatePostModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc Controller) PostShow(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.GetPostByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) PostUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.UpdatePostByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) PostDelete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	if err := s.DeletePostByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}