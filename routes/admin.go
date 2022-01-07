// Package routes 注册路由
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAdminRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	admin := r.Group("/admin")
	{
		// 注册路由
		admin.GET("/ping", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"message": "pong!",
			})
		})
	}
}
