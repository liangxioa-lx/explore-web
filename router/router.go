package router

import (
	"contract/controller/file"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	r := router.Group("/api")

	fileController.RegisterRouter(r)

	router.Static("/web", "./public")
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})
}
