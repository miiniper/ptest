//package db
//
//import (
//	"ptest/loges"
//
//	"github.com/spf13/viper"
//	"go.uber.org/zap"
//	"gopkg.in/mgo.v2"
//)
//
//var Session *mgo.Session
//
//func InitMgo() {
//	mgoAddr := viper.GetString("mgo.addr")
//	Session, err := mgo.Dial(mgoAddr)
//	if err != nil {
//		loges.Loges.Error("", zap.Error(err))
//	}
//	defer Session.Close()
//	Session.SetMode(mgo.Monotonic, true)
//	err = Session.DB("admin").Login(viper.GetString("mgo.user"), viper.GetString("mgo.pwd"))
//	if err != nil {
//		loges.Loges.Error("", zap.Error(err))
//	}
//
//}
//
////
////func Inserts(dbtable string, data interface{}) error {
////	c := Session.DB("ptest").C(dbtable)
////	err := c.Insert(&data)
////	if err != nil {
////		loges.Loges.Error("", zap.Error(err))
////		return err
////	}
////	return nil
////
////}
