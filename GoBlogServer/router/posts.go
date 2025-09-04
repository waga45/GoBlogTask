package router

import (
	"GoBlogServer/internal/container"
	"GoBlogServer/internal/controller"
	"GoBlogServer/internal/repository/mapper"
	"GoBlogServer/internal/service"
	"github.com/gin-gonic/gin"
)

/*
*
初始化文章相关路由和服务
*/
func InitPostsRouter(Container *container.Container, Router *gin.RouterGroup) {
	//数据
	postsMapper := mapper.NewPostsMapper(Container.GetGormDb())
	postService := service.NewPostService(postsMapper, Container.GetRedisClient())
	controller := controller.NewPostsController(postService)
	//路由
	userRouterPrivate := Router.Group("posts")
	userRouterPrivate.POST("new", controller.AddNewPosts)
	userRouterPrivate.POST("update", controller.UpdatePosts)
	userRouterPrivate.POST("detail", controller.GetPostDetail)
	userRouterPrivate.POST("remove", controller.RemovePosts)
}
