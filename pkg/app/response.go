package app

import (
	"gin-blog/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{})  {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg": e.GetMsg(errCode),
		"data": data,
	})

	return
}

func (g *Gin) ResponseSuccess() {
	g.Response(http.StatusOK, e.SUCCESS, make(map[string]string))
	return
}