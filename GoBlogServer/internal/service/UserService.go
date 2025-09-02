package service

import (
	"GoBlogServer/internal/dtos"
	"GoBlogServer/internal/repository/mapper"
	"GoBlogServer/internal/vos"
)

type UserService struct {
	userMapper *mapper.UserMapper
}

func NewUserService(userMapper *mapper.UserMapper) *UserService {
	return &UserService{userMapper: userMapper}
}

/*
*
登入
*/
func UserLogin(vo *vos.UserLoginVO) (*dtos.LoginDto, error) {
	return nil, nil
}
