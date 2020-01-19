package db

import (
	"ptest/loges"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

func InitMgo() {
	sess, err := mgo.Dial(viper.GetString("mgo.addr"))
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}

	sess.SetMode(mgo.Monotonic, true)
	err = sess.DB("admin").Login(viper.GetString("mgo.user"), viper.GetString("mgo.pwd"))
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}
	Sess = sess
}

var Sess *mgo.Session

func Inserts(table string, data interface{}) error {
	//Sess = InitMgo()
	defer Sess.Close()
	c := Sess.DB("ptest").C(table)
	err := c.Insert(&data)
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
		return err
	}
	return nil
}

//
//func Selects() {
//	sess := InitMgo()
//	defer sess.Close()
//}
