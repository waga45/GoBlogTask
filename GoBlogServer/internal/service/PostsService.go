package service

import (
	"GoBlogServer/internal/component"
	"GoBlogServer/internal/repository/mapper"
)

type PostService struct {
	postsMapper *mapper.PostsMapper
	redisClient *component.RedisClient
}

func NewPostService(u *mapper.PostsMapper, rd *component.RedisClient) *PostService {
	return &PostService{postsMapper: u, redisClient: rd}
}
