package main //import ptest

import (
	"ptest/conf"
	"ptest/loges"
	"ptest/router"
)

func main() {
	loges.Loges.Info("starting server....")
	conf.InitConf()
	//	db.InitMgo()

	router.InitRouter()

}
