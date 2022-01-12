// Package routes 注册路由
package routes

import (
	controllers "huango/app/http/controllers/api/v1"
	"huango/app/http/controllers/api/v1/auth"
	"huango/app/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/api/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// ---- 登录 ----
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名 登录
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			// Token 刷新
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// ---- 重置密码 ----
			pwc := new(auth.PasswordController)
			// 手机号 + 验证码 重置密码
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			// Email + 验证码 重置密码
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)

			// ---- 注册用户 ----
			suc := new(auth.SignupController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)
			// 判断邮箱是否注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
			// 手机 + 验证码 注册
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			// Email + 验证码 注册
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SingupUsingEmail)

			// ---- 发送验证码 ----
			vcc := new(auth.VerifyCodeController)
			// 发送短信验证码
			authGroup.POST("/verify-code/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			// 发送 Email 验证码
			authGroup.POST("/verify-code/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-code/captcha", middlewares.LimitPerRoute("20-H"), vcc.ShowCaptcha)
		}

		// ---- 用户接口 ----
		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
		// 用户路由分组
		usersGroup := v1.Group("/users")
		{
			// 用户列表
			usersGroup.GET("", uc.Index)
		}

		// ---- 分类接口 ----
		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			// 分类列表
			cgcGroup.GET("", cgc.Index)
			// 新建分类
			cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
			// 更新分类
			cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
			// 删除分类
			cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
		}

		// ---- 分类接口 ----
		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			// 新建话题
			tpcGroup.POST("", middlewares.AuthJWT(), tpc.Store)
			// 更新话题
			tpcGroup.PUT("/:id", middlewares.AuthJWT(), tpc.Update)
			// 删除分类
			tpcGroup.DELETE("/:id", middlewares.AuthJWT(), tpc.Delete)
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
