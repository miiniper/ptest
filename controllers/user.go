package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"ptest/db"
	"ptest/loges"

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

	db.Inserts("user", user)

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
