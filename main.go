package main //import ptest

import (
	"ptest/conf"
	"ptest/loges"
	"ptest/router"
)

func main() {
	loges.Loges.Info("starting server....")

	router.InitRouter()

}

func init() {
	conf.InitConf()
	//db.InitMgo()
}
