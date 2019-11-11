package main

import (
	"github.com/KodaiD/rest/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	db.Init()
	r.Run()

	db.Close()
}