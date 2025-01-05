package server

import (
	"github.com/RosaLeaveMeAlone/gin-crud/src/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
