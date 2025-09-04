package controller

import (
	"GoBlogServer/internal/constants"
	"GoBlogServer/internal/response"
	"GoBlogServer/internal/service"
	"GoBlogServer/internal/vos"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PostsController struct {
	postsService *service.PostService
}

func NewPostsController(p *service.PostService) *PostsController {
	return &PostsController{postsService: p}
}

/*
*
查询列表
*/
func (p *PostsController) GetList(ctx *gin.Context) {
	userId := ctx.GetInt(constants.UserIdKey)
	var vo vos.GetPostListVO
	ctx.ShouldBindJSON(&vo)
	pageList, err := p.postsService.List(userId, &vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}
	response.Success(ctx, pageList)
}

/*
*
写文章
*/
func (p *PostsController) AddNewPosts(ctx *gin.Context) {
	userId := ctx.GetInt(constants.UserIdKey)
	var vo vos.UpdatePostsVO
	ctx.ShouldBindJSON(&vo)
	id, err := p.postsService.NewPosts(userId, &vo)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}
	response.Success(ctx, id)
}

/*
*
更新文章
*/
func (p *PostsController) UpdatePosts(ctx *gin.Context) {
	userId := ctx.GetInt(constants.UserIdKey)
	var vo vos.UpdatePostsVO
	ctx.ShouldBindJSON(&vo)
	b, err := p.postsService.UpdatePosts(userId, &vo)
	if err != nil || !b {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}
	response.Success(ctx, b)
}

/*
*
获取详细
*/
func (p *PostsController) GetPostDetail(ctx *gin.Context) {
	strId := ctx.Query("id")
	if len(strId) <= 0 {
		response.Error(ctx, constants.ResponseError, "请选择相应文章")
		return
	}
	fmt.Println("strId=", strId)
	id, _ := strconv.ParseInt(strId, 10, 64)
	bean, err := p.postsService.DetailInfo(id)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
	}
	response.Success(ctx, *bean)
}

/*
*
更新
*/
func (p *PostsController) RemovePosts(ctx *gin.Context) {
	userId := ctx.GetInt(constants.UserIdKey)
	if userId <= 0 {
		response.Error(ctx, constants.ResponseError, "请选择登入后操作")
		return
	}
	var vo vos.RemoveDataVO
	ctx.ShouldBind(&vo)
	pId, _ := strconv.ParseInt(vo.Id, 10, 64)
	b, err := p.postsService.RemoveById(userId, pId)
	if err != nil {
		response.Error(ctx, constants.ResponseError, err.Error())
		return
	}
	response.Success(ctx, b)

}
