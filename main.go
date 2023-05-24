package main

import (
	"backend-demo/controllers"
	"backend-demo/initialzers"

	"github.com/gin-gonic/gin"
)

func init() {
	initialzers.LoadEnv()
	initialzers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.GET("/posts", controllers.PostFindAll)
	r.GET("/posts/:id", controllers.PostFindOne)

	r.Run()
}
