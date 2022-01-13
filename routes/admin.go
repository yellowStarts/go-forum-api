// Package routes 注册路由
package routes

import (
	"net/http"
	controllers "huango/app/http/controllers/admin"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAdminRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	admin := r.Group("/admin")
	{
		// ---- 系统状态 ----
		systemGroup := admin.Group("/system")
		{
			syc := new(controllers.SystemsController)
			// 系统状态
			systemGroup.GET("/health", syc.HealthCheck)
			// 硬盘状态
			systemGroup.GET("/disk", syc.DiskCheck)
			// CPU状态
			systemGroup.GET("/cpu", syc.CPUCheck)
			// 内存状态
			systemGroup.GET("/ram", syc.RAMCheck)
		}
		// 注册路由
		admin.GET("/ping", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"message": "pong!",
			})
		})
	}
}
