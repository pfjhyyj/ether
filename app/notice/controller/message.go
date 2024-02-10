package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/notice/define"
	"github.com/pfjhyyj/ether/app/notice/service"
	"github.com/pfjhyyj/ether/app/notice/utils"
	"github.com/pfjhyyj/ether/common"
	utils2 "github.com/pfjhyyj/ether/utils"
	"net/http"
)

type MessageController struct {
	service *service.MessageService
}

func NewMessageController(service *service.MessageService) *MessageController {
	return &MessageController{
		service: service,
	}
}

// ListMyMessages godoc
// @Summary List my messages
// @Description List my messages
// @Tags message
// @Accept json
// @Produce json
// @Security Bearer
// @Param request query define.ListMyMessagesRequest true "ListMyMessagesRequest"
// @Success 200 {object} common.Response{data=common.Page={List=[]define.ListMyMessageResponse}}
// @Router /messages [get]
func (c *MessageController) ListMyMessages(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "message", "list"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}
	userId := ctx.GetUint(common.CtxUserIDKey)

	var req define.ListMyMessagesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	queryParam := utils.ConvertListMyMessagesRequestToQueryParam(&req)
	queryParam.UserId = userId
	queryParam.WithText = true

	messages, total, err := c.service.ListMyMessages(ctx, queryParam)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	messagesInfo := utils.ConvertMessageListToListMyMessageResponse(messages)
	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
		Data: &common.Page{
			Current: req.Current,
			Total:   total,
			List:    messagesInfo,
		},
	})
}

// ReadMessage godoc
// @Summary Read message
// @Description Read message
// @Tags message
// @Accept json
// @Produce json
// @Security Bearer
// @Param messageId path uint true "Message ID"
// @Success 200 {object} common.Response
// @Router /messages/{messageId}/read [put]
func (c *MessageController) ReadMessage(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "message", "read"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.ReadMessageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.ReadMessage(ctx, req.MessageId); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}

// BatchReadMessage godoc
// @Summary Batch read message
// @Description Batch read message
// @Tags message
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body define.BatchReadMessageRequest true "BatchReadMessageRequest"
// @Success 200 {object} common.Response
// @Router /messages/batchRead [put]
func (c *MessageController) BatchReadMessage(ctx *gin.Context) {
	if ok := utils2.CheckPermission(ctx, "message", "read"); !ok {
		ctx.JSON(http.StatusForbidden, &common.Response{
			Code: common.NoPermissionError,
			Msg:  "no permission",
		})
		return
	}

	var req define.BatchReadMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := c.service.BatchReadMessage(ctx, req.MessageIds); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, &common.Response{
		Code: common.Ok,
	})
}
