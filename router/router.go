package router

import (
	"ptest/controllers"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	v1 := router.Group("v1")
	v1.GET("/user/login", controllers.UserLogin)
	v1.POST("/user/register", controllers.Register)
	v1.GET("/chk", controllers.Chk)

	router.Run(viper.GetString("server.port"))
}
