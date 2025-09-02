package mapper

import (
	"GoBlogServer/internal/repository/model"
	"gorm.io/gorm"
)

type PostsMapper struct {
	db *gorm.DB
}

func NewPostsMapper(db *gorm.DB) *PostsMapper {
	return &PostsMapper{db: db}
}

func (p *PostsMapper) SelectById(postId int) (*model.Post, error) {
	return nil, nil
}
