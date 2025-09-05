package controller

import (
	"GoBlogServer/internal/constants"
	"GoBlogServer/internal/response"
	"GoBlogServer/internal/service"
	"GoBlogServer/internal/vos"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController(_commentService *service.CommentService) *CommentController {
	return &CommentController{commentService: _commentService}
}

/*
*
添加评论
*/
func (c *CommentController) AddComment(ctx *gin.Context) {
	var vo vos.PushCommentVO
	ctx.ShouldBindJSON(&vo)
	userId := ctx.GetInt(constants.UserIdKey)
	b, err := c.commentService.CreateComment(userId, &vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}
	if !b {
		response.Error(ctx, constants.ResponseError, "服务异常，请稍后重试")
		return
	}
	response.Success(ctx, b)
}

/*
*
列表
*/
func (c *CommentController) GetPostComments(ctx *gin.Context) {
	var vo vos.GetCommentListVO
	err := ctx.ShouldBindJSON(&vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}
	result, err := c.commentService.GetPostComments(vo.PostId, vo.PageIndex, vo.PageSize)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}
	response.Success(ctx, result)
}

/*
*
删除评论
*/
func (c *CommentController) Remove(ctx *gin.Context) {
	var vo vos.CommentVO
	err := ctx.ShouldBindJSON(&vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
	}
	userId := ctx.GetInt(constants.UserIdKey)
	if userId <= 0 {
		response.Error(ctx, constants.ResponseError, "请登入后操作！")
		return
	}
	b, err := c.commentService.RemoveComments(userId, vo.PostId, vo.Id)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
	}
	response.Success(ctx, b)
}
