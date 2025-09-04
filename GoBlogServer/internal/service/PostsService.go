package service

import (
	"GoBlogServer/internal/component"
	"GoBlogServer/internal/dtos"
	"GoBlogServer/internal/repository/mapper"
	"GoBlogServer/internal/repository/model"
	"GoBlogServer/internal/vos"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"strconv"
)

type PostService struct {
	postsMapper *mapper.PostsMapper
	redisClient *component.RedisClient
	nodeId      *snowflake.Node
}

func NewPostService(u *mapper.PostsMapper, rd *component.RedisClient) *PostService {
	n, _ := snowflake.NewNode(1)
	return &PostService{postsMapper: u, redisClient: rd, nodeId: n}
}

var node *snowflake.Node

/*
**
列表查询
*/
func (p *PostService) List(userId int, vo *vos.GetPostListVO) (*dtos.PagePostsListDto, error) {
	if userId <= 0 {
		return nil, errors.New("请登入后操作")
	}
	if vo.PageIndex <= 0 {
		vo.PageIndex = 1
	}
	if vo.PageSize <= 0 {
		vo.PageSize = 10
	}
	offset := (vo.PageIndex - 1) * vo.PageSize
	list, count, _ := p.postsMapper.SelectList(userId, vo.State, vo.Title, offset, vo.PageSize)
	return &dtos.PagePostsListDto{List: list, TotalCount: count}, nil
}

/*
*
新增
*/
func (p *PostService) NewPosts(userId int, vo *vos.UpdatePostsVO) (int64, error) {
	if userId <= 0 {
		return -1, errors.New("请登入后操作")
	}

	//雪花id
	var id = p.nodeId.Generate().Int64()
	pId, err := p.postsMapper.SaveNew(id, userId, vo.Title, vo.Content)
	return pId, err
}

/*
*
更新
*/
func (p *PostService) UpdatePosts(userId int, vo *vos.UpdatePostsVO) (bool, error) {
	if userId <= 0 {
		return false, errors.New("请登入后操作")
	}
	fmt.Println(vo)
	id, err := strconv.ParseInt(vo.Id, 10, 64)
	if id <= 0 {
		return false, errors.New("请选择更新文章")
	}
	hasPermission := p.postsMapper.IsAuthorTo(id, userId)
	if !hasPermission {
		return false, errors.New("您无更新权限")
	}
	b, err := p.postsMapper.UpdateById(id, vo.Title, vo.Content, vo.State)
	return b, err
}

/*
*
详细信息
*/
func (p *PostService) DetailInfo(id int64) (*model.Post, error) {
	if id <= 0 {
		return nil, errors.New("请选择对应文章")
	}
	return p.postsMapper.SelectById(id)
}

/*
*
删除
*/
func (p *PostService) RemoveById(userId int, postId int64) (bool, error) {
	if userId <= 0 {
		return false, errors.New("请登入后操作")
	}
	if postId <= 0 {
		return false, errors.New("请选择删除文章")
	}
	hasPermission := p.postsMapper.IsAuthorTo(postId, userId)
	if !hasPermission {
		return false, errors.New("您无删除权限")
	}
	return p.postsMapper.RemoveById(postId)
}
