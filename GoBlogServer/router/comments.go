package router

import (
	"GoBlogServer/internal/container"
	"GoBlogServer/internal/controller"
	"GoBlogServer/internal/middleware"
	"GoBlogServer/internal/repository/mapper"
	"GoBlogServer/internal/service"
	"github.com/gin-gonic/gin"
)

/*
初始化  其实同类路由可以在一个文件初始化
*/
func InitCommentsRouter(Container *container.Container, Router *gin.RouterGroup) {
	//实例化service
	commentMapper := mapper.NewCommentMapper(Container.GetGormDb())
	postMapper := mapper.NewPostsMapper(Container.GetGormDb())
	commentService := service.NewCommentService(commentMapper, postMapper)
	commentController := controller.NewCommentController(commentService)

	//路由
	publicRouter := Router.Group("comments")
	publicRouter.POST("/list", commentController.GetPostComments)

	privateRouter := Router.Group("comments")
	privateRouter.Use(middleware.HandlerAuthWare())
	privateRouter.POST("/new", commentController.AddComment)
	privateRouter.POST("/disable", commentController.Remove)
}
