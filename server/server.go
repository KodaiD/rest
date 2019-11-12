package server

import (
	"github.com/KodaiD/rest/controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controller.Controller{}
		u.GET("", ctrl.UserIndex)
		u.GET("/:id", ctrl.UserShow)
		u.POST("", ctrl.UserCreate)
		u.PUT("/:id", ctrl.UserUpdate)
		u.DELETE("/:id", ctrl.UserDelete)
	}

	p := r.Group("/posts")
	{
		ctrl := controller.Controller{}
		p.GET("", ctrl.PostIndex)
		p.GET("/:id", ctrl.PostShow)
		p.POST("", ctrl.PostCreate)
		p.PUT("/:id", ctrl.PostUpdate)
		p.DELETE("/:id", ctrl.PostDelete)
	}

	return r
}