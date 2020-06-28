package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//0.创建路由
	r := gin.Default()
	//1.绑定路由规则，执行函数
	//gin.Context封装了request和response, 直接调用c获取其他属性或方法
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.Run()
}
