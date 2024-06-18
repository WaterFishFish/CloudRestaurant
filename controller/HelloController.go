package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)
}
func (hello *HelloController) Hello(context *gin.Context) {
	context.JSON(200, gin.H{
		"msg": "hello cloudrestaurant",
	})

}
