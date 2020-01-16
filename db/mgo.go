package db

import (
	"ptest/loges"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

var Session = InitMgo()

func InitMgo() *mgo.Session {
	mgoAddr := viper.GetString("mgo.addr")
	session, err := mgo.Dial(mgoAddr)
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	err = session.DB("admin").Login(viper.GetString("mgo.user"), viper.GetString("mgo.pwd"))
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}

	return session

}
