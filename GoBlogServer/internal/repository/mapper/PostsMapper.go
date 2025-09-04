package mapper

import (
	"GoBlogServer/internal/repository/model"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

type PostsMapper struct {
	db *gorm.DB
}

var node *snowflake.Node

func NewPostsMapper(db *gorm.DB) *PostsMapper {
	return &PostsMapper{db: db}
}

func (p *PostsMapper) SelectById(postId int) (*model.Post, error) {
	var result model.Post
	r := p.db.Where("id=?", postId).First(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return &result, nil
}

// 创建
func (p *PostsMapper) SaveNew(userId int, title string, content string) (int64, error) {
	var t = time.Now()
	var id = node.Generate().Int64()
	var tempBean = model.Post{Id: id, UserId: userId, Title: title, Content: &content, State: 1, RateNumber: 0, CreateTime: &t}
	r := p.db.Create(&tempBean)
	if r.Error != nil {
		return -1, r.Error
	}
	return id, nil
}

func (p *PostsMapper) RemoveById(postId int64) (bool, error) {
	r := p.db.Where("id=?", postId).Delete(&model.Post{})
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}
