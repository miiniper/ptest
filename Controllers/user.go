package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpStatus struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
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
