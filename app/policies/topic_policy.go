// Package policies 用户授权
package policies

import (
	"huango/app/models/topic"
	"huango/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}