package app

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/e"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResult[T interface{}] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

func (g *Gin) Ok(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  "",
		Data: data,
	})
}

func (g *Gin) Error(errCode int, msg string) {
	g.C.JSON(200, Response{
		Code: errCode,
		Msg:  msg,
		Data: nil,
	})
}

func (g *Gin) ErrorCode(errCode int) {
	g.C.JSON(200, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: nil,
	})
}
