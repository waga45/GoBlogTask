package service

import (
	"GoBlogServer/internal/component"
	"GoBlogServer/internal/dtos"
	"GoBlogServer/internal/repository/mapper"
	"GoBlogServer/internal/repository/model"
	"GoBlogServer/internal/vos"
	"GoBlogServer/utils"
	"errors"
	"strings"
)

type UserService struct {
	userMapper *mapper.UserMapper
}

func NewUserService(userMapper *mapper.UserMapper) *UserService {
	return &UserService{userMapper: userMapper}
}

/*
*
注册，成功返回授权token
*/
func (u *UserService) RegisterUser(vo *vos.UserRegisterVO) (*dtos.LoginDto, error) {
	b, err := u.userMapper.ExistByUserName(vo.Username)
	if err != nil {
		return nil, err
	}
	if b {
		return nil, errors.New("该用户名已存在，请更换")
	}
	pwd := utils.Md5String(vo.Password)
	id, err := u.userMapper.SaveNew(vo.Username, vo.Email, pwd)
	if err != nil {
		return nil, err
	}
	return genLoginAuth(id, vo.Username), nil
}

/*
*
登入
*/
func (u *UserService) UserLogin(vo *vos.UserLoginVO) (*dtos.LoginDto, error) {
	user, err := u.userMapper.SelectByUserName(vo.Username)
	if err != nil {
		return nil, err
	}
	if user.State != 1 {
		return nil, errors.New("用户被禁止的登入")
	}
	pwd := utils.Md5String(vo.Password)
	if strings.Compare(user.Password, pwd) != 0 {
		return nil, errors.New("账号或密码错误")
	}
	u.userMapper.UpdateUserLoginTime(user.Id)
	return genLoginAuth(user.Id, user.UserName), nil
}

/*
*
获取用户信息
*/
func (u *UserService) GetUserInfo(id int) (*model.Users, error) {
	if id <= 0 {
		return nil, errors.New("error")
	}
	user, err := u.userMapper.SelectById(id)
	if err != nil {
		return nil, errors.New("该用户信息不存在")
	}
	return user, nil
}

func genLoginAuth(id int, name string) *dtos.LoginDto {
	token, err := component.GetJwtInstance().GenToken(id, name)
	if err != nil {
		panic(err)
	}
	return &dtos.LoginDto{Token: token, UserName: name}
}
