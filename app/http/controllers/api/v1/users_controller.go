package v1

import (
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
