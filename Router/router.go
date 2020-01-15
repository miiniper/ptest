package Router

import (
	"ptest/Controllers"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	v1 := router.Group("v1")
	v1.GET("/login", Controllers.UserLogin)
	v1.GET("/chk", Controllers.Chk)

	router.Run(viper.GetString("server.port"))
}
