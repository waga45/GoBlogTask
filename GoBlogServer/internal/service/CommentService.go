package service

import (
	"GoBlogServer/internal/dtos"
	"GoBlogServer/internal/repository/mapper"
	"GoBlogServer/internal/vos"
	"errors"
	"strconv"
)

type CommentService struct {
	commentMapper *mapper.CommentMapper
	postsMapper   *mapper.PostsMapper
}

func NewCommentService(_commentMapper *mapper.CommentMapper, _postsMapper *mapper.PostsMapper) *CommentService {
	return &CommentService{commentMapper: _commentMapper, postsMapper: _postsMapper}
}

/*
*
发起评论
*/
func (c *CommentService) CreateComment(userId int, vo *vos.PushCommentVO) (bool, error) {
	if userId <= 0 {
		return false, errors.New("请登入后操作")
	}
	if len(vo.PostId) <= 0 {
		return false, errors.New("请选择评论文章")
	}
	if len(vo.Content) <= 0 {
		return false, errors.New("请输入您的评论")
	}
	postId, err := strconv.ParseInt(vo.PostId, 10, 64)
	if err != nil {
		return false, err
	}
	_, err = c.commentMapper.SaveNew(userId, postId, vo.Content)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
*
评论列表
*/
func (c *CommentService) GetPostComments(postId string, pageIndex int, pageSize int) (*dtos.BasePageDto[dtos.PostCommentDto], error) {
	if len(postId) <= 0 {
		return nil, errors.New("请选择评论文章")
	}
	id, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		return nil, err
	}
	pageOffset := CalculateOffsetNum(&pageIndex, &pageSize)
	return c.commentMapper.GetListByPosts(id, pageOffset, pageSize)
}

/*
*
移除评论
*/
func (c *CommentService) RemoveComments(userId int, postId string, id int) (bool, error) {
	if userId <= 0 {
		return false, errors.New("请登入后操作")
	}
	if len(postId) <= 0 {
		return false, nil
	}
	if id <= 0 {
		return false, nil
	}
	pId, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		return false, err
	}
	hasPermission := c.postsMapper.IsAuthorTo(pId, userId)
	if !hasPermission {
		return false, errors.New("只有作者才有权限操作")
	}
	b, err := c.commentMapper.UpdateDisable(id)
	if err != nil {
		return false, err
	}
	return b, nil
}
