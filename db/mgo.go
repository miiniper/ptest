package db

import (
	"ptest/loges"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

func InitMgo() *mgo.Session {
	Sess, err := mgo.Dial(viper.GetString("mgo.addr"))
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}

	//	defer Sess.Close()

	Sess.SetMode(mgo.Monotonic, true)
	err = Sess.DB("admin").Login(viper.GetString("mgo.user"), viper.GetString("mgo.pwd"))
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}
	return Sess
}

func Inserts(table string, data interface{}) error {
	sess := InitMgo()
	defer sess.Close()
	c := sess.DB("ptest").C(table)
	err := c.Insert(&data)
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
		return err
	}
	return nil

}
