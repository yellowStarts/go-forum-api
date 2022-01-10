// Package routes 注册路由
package routes

import (
	"huango/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 手机 + 验证码 注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			// Email + 验证码 注册
			authGroup.POST("/signup/using-email", suc.SingupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-code/captcha", vcc.ShowCaptcha)
			// 发送短信验证码
			authGroup.POST("/verify-code/phone", vcc.SendUsingPhone)
			// 发送 Email 验证码
			authGroup.POST("/verify-code/email", vcc.SendUsingEmail)

			// 用户授权
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名 登录
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
			// Token 刷新
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", pwc.ResetByPhone)
		}

		// 测试路由
		v1.GET("/ping", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"message": "pong!",
			})
		})
	}
}
