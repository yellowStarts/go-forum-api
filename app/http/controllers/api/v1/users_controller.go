package v1

import (
	"huango/app/models/user"
	"huango/pkg/auth"
	"huango/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
