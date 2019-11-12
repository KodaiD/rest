package controller

import (
	"fmt"
	"github.com/KodaiD/rest/service"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (pc Controller) UserIndex(c *gin.Context) {
	var s service.Service
	p, err := s.GetUsersAll()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) UserCreate(c *gin.Context) {
	var s service.Service
	p, err := s.CreateUserModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

func (pc Controller) UserShow(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.GetUserByID(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) UserUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.UpdateUserByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) UserDelete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	if err := s.DeleteUserByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}