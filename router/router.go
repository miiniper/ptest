package router

import (
	"net/http"
	"ptest/controllers"

	"github.com/spf13/viper"

	"github.com/julienschmidt/httprouter"
)

func InitRouter() {
	router := httprouter.New()

	router.GET("/user/login", controllers.UserLogin)
	router.POST("/user/register", controllers.Register)
	router.GET("/chk", controllers.Chk)
	http.ListenAndServe(viper.GetString("server.port"), router)

	//router.Run(viper.GetString("server.port"))
}
