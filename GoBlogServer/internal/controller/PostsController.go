package controller

import (
	"GoBlogServer/internal/service"
	"github.com/gin-gonic/gin"
)

type PostsController struct {
	postsService *service.PostService
}

func NewPostsController(p *service.PostService) *PostsController {
	return &PostsController{postsService: p}
}

func (p *PostsController) AddNewPosts(ctx *gin.Context) {

}

func (p *PostsController) UpdatePosts(ctx *gin.Context) {

}

func (p *PostsController) GetPostDetail(ctx *gin.Context) {

}

func (p *PostsController) RemovePosts(ctx *gin.Context) {

}