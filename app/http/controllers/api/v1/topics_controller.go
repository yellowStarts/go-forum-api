package v1

import (
	"huango/app/http/requests"
	"huango/app/models/topic"
	"huango/pkg/auth"
	"huango/pkg/logger"
	"huango/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseController
}

// func (ctrl *TopicsController) Index(c *gin.Context) {
//     topics := topic.All()
//     response.Data(c, topics)
// }

// func (ctrl *TopicsController) Show(c *gin.Context) {
//     topicModel := topic.Get(c.Param("id"))
//     if topicModel.ID == 0 {
//         response.Abort404(c)
//         return
//     }
//     response.Data(c, topicModel)
// }

func (ctrl *TopicsController) Store(c *gin.Context) {

	logger.Dump(auth.CurrentUID(c), "当前用户ID")

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

// func (ctrl *TopicsController) Update(c *gin.Context) {

//     topicModel := topic.Get(c.Param("id"))
//     if topicModel.ID == 0 {
//         response.Abort404(c)
//         return
//     }

//     if ok := policies.CanModifyTopic(c, topicModel); !ok {
//         response.Abort403(c)
//         return
//     }

//     request := requests.TopicRequest{}
//     bindOk, errs := requests.Validate(c, &request, requests.TopicSave)
//     if !bindOk {
//         return
//     }
//     if len(errs) > 0 {
//         response.ValidationError(c, 20101, errs)
//         return
//     }

//     topicModel.FieldName = request.FieldName
//     rowsAffected := topicModel.Save()
//     if rowsAffected > 0 {
//         response.Data(c, topicModel)
//     } else {
//         response.Abort500(c, "更新失败，请稍后尝试~")
//     }
// }

// func (ctrl *TopicsController) Delete(c *gin.Context) {

//     topicModel := topic.Get(c.Param("id"))
//     if topicModel.ID == 0 {
//         response.Abort404(c)
//         return
//     }

//     if ok := policies.CanModifyTopic(c, topicModel); !ok {
//         response.Abort403(c)
//         return
//     }

//     rowsAffected := topicModel.Delete()
//     if rowsAffected > 0 {
//         response.Success(c)
//         return
//     }

//     response.Abort500(c, "删除失败，请稍后尝试~")
// }
