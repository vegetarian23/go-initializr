package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vegetarian23/go-initializr/internal/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user/1", controller.NewUserController().GetUserById)
	}

	return r
}
