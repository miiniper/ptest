package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ptest/loges"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
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

func Register(c *gin.Context) {
	var user UserInfo
	user.Name = c.Query("name")
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		loges.Loges.Error("", zap.Error(err))
	}
	fmt.Println(user.Name)
}

func UserLogin(c *gin.Context) {
	//fmt.Println("todo ")
	h1 := HttpStatus{StatusCode: 200, Msg: "ok"}
	w2, _ := json.Marshal(h1)
	c.Writer.Write(w2)

}

func Chk(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, message)

}
