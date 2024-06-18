package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCESS int = 0
	FAILED int = 1
)

//普通成功返回

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": SUCESS,
		"msg":  "成功",
		"data": v,
	})
}

//普通操作失败

func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": FAILED,
		"msg":  v,
	})
}
