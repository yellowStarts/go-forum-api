package main

import (
	"fmt"
	"huango/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务，默认为 8080，我们指定端口为 8000
	err := router.Run(":8000")
	if err != nil {
		// 错误处理，端口被占用或者其他错误
		fmt.Println(err.Error())
	}
}
