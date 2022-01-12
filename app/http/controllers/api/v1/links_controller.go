package v1

import (
    "huango/app/models/link"
    "huango/pkg/response"

    "github.com/gin-gonic/gin"
)

type LinksController struct {
    BaseController
}

func (ctrl *LinksController) Index(c *gin.Context) {
    response.Data(c, link.AllCached())
}

// func (ctrl *LinksController) Show(c *gin.Context) {
//     linkModel := link.Get(c.Param("id"))
//     if linkModel.ID == 0 {
//         response.Abort404(c)
//         return
//     }
//     response.Data(c, linkModel)
// }

// func (ctrl *LinksController) Store(c *gin.Context) {

//     request := requests.LinkRequest{}
//     if ok := requests.Validate(c, &request, requests.LinkSave); !ok {
//         return
//     }

//     linkModel := link.Link{
//         FieldName:      request.FieldName,
//     }
//     linkModel.Create()
//     if linkModel.ID > 0 {
//         response.Created(c, linkModel)
//     } else {
//         response.Abort500(c, "创建失败，请稍后尝试~")
//     }
// }

// func (ctrl *LinksController) Update(c *gin.Context) {

//     linkModel := link.Get(c.Param("id"))
//     if linkModel.ID == 0 {
//         response.Abort404(c)
//         return
//     }

//     if ok := policies.CanModifyLink(c, linkModel); !ok {
//         response.Abort403(c)
//         return
//     }

//     request := requests.LinkRequest{}
//     bindOk, errs := requests.Validate(c, &request, requests.LinkSave)
//     if !bindOk {
//         return
//     }
//     if len(errs) > 0 {
//         response.ValidationError(c, 20101, errs)
//         return
//     }

//     linkModel.FieldName = request.FieldName
//     rowsAffected := linkModel.Save()
//     if rowsAffected > 0 {
//         response.Data(c, linkModel)
//     } else {
//         response.Abort500(c, "更新失败，请稍后尝试~")
//     }
// }

// func (ctrl *LinksController) Delete(c *gin.Context) {

//     linkModel := link.Get(c.Param("id"))
//     if linkModel.ID == 0 {
//         response.Abort404(c)
//         return
//     }

//     if ok := policies.CanModifyLink(c, linkModel); !ok {
//         response.Abort403(c)
//         return
//     }

//     rowsAffected := linkModel.Delete()
//     if rowsAffected > 0 {
//         response.Success(c)
//         return
//     }

//     response.Abort500(c, "删除失败，请稍后尝试~")
// }