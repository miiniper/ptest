package main //import ptest

import (
	"ptest/Conf"
	"ptest/Loges"
	"ptest/Router"
)

func main() {
	Loges.Loges.Info("starting server....")
	Conf.InitConf()

	Router.InitRouter()

}
