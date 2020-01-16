package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ptest/loges"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"

	"go.uber.org/zap"
)

type HttpStatus struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

type UserInfo struct {
	Name     string `json:"name"`
	StaffId  string `json:"staff_id"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user UserInfo

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}
	fmt.Println(user.Name, user.Mail, user.Password, user.StaffId)
	//db.Inserts("user", user)

	mgoAddr := viper.GetString("mgo.addr")
	Session, err := mgo.Dial(mgoAddr)
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}
	defer Session.Close()
	Session.SetMode(mgo.Monotonic, true)
	err = Session.DB("admin").Login(viper.GetString("mgo.user"), viper.GetString("mgo.pwd"))
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}

	cc := Session.DB("ptest").C("user")
	err = cc.Insert(user)
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}
}

func UserLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Println("todo ")
	h1 := HttpStatus{StatusCode: 200, Msg: "ok"}
	w2, _ := json.Marshal(h1)
	w.Write(w2)

}

func Chk(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h1 := HttpStatus{StatusCode: 200, Msg: "ok"}
	w2, _ := json.Marshal(h1)
	w.Write(w2)

}
