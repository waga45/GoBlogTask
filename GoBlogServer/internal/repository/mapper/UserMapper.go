package mapper

import (
	"GoBlogServer/internal/repository/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

type UserMapper struct {
	db *gorm.DB
}

func NewUserMapper(db *gorm.DB) *UserMapper {
	return &UserMapper{db: db}
}

/*
**
用户名是否存在
*/
func (m *UserMapper) ExistByUserName(username string) (bool, error) {
	count := int64(0)
	m.db.Model(&model.Users{}).Where(&model.Users{UserName: username, State: 1}).Count(&count)
	return count > 0, nil
}

/*
*
保存用户
*/
func (m *UserMapper) SaveNew(username string, email string, password string) (int, error) {
	t := time.Now()
	bean := model.Users{UserName: username, Email: email, Password: password, State: 1, CreateTime: &t, LastLogin: &t}
	r := m.db.Create(&bean)
	if r.Error != nil {
		return 0, r.Error
	}
	return int(r.RowsAffected), nil
}

/*
*
根据用户名查询用户
*/
func (m *UserMapper) SelectByUserName(value string) (*model.Users, error) {
	var user model.Users
	r := m.db.Model(&model.Users{}).Where(&model.Users{UserName: value, State: 1}).First(&user)
	if r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

/*
*
更新
*/
func (m *UserMapper) UpdateById(user *model.Users) (bool, error) {
	if user == nil {
		return false, errors.New("更新失败")
	}
	r := m.db.Where("id =?", user.Id).Updates(user)
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

/*
*
更新登入时间
*/
func (m *UserMapper) UpdateUserLoginTime(id int) (bool, error) {
	var t = time.Now()
	r := m.db.Model(&model.Users{}).Where("id =?", id).Update("last_login", t)
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

/*
*
根据ID查询
*/
func (m *UserMapper) SelectById(id int) (*model.Users, error) {
	var user model.Users
	r := m.db.Model(&model.Users{}).Where(&model.Users{Id: id}).First(&user)
	if r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}
