package mapper

import (
	"GoBlogServer/internal/dtos"
	"GoBlogServer/internal/repository/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

type PostsMapper struct {
	db *gorm.DB
}

func NewPostsMapper(db *gorm.DB) *PostsMapper {
	return &PostsMapper{db: db}
}

/*
*
查询列表
*/
func (p *PostsMapper) SelectList(userId int, state int, title string, pageOffset int, pageSize int) (*[]dtos.PostsListDto, int64, error) {
	table := p.db.Debug().Model(&model.Post{})
	table.Select("id,title,SUBSTRING(content,1,100) as content,create_time,state")
	table.Where("user_id = ?", userId)
	if state >= 0 {
		table.Where("state=?", state)
	}
	if len(title) > 0 {
		table.Where("title LIKE ?", "%"+title+"%")
	}
	var totalNum int64
	table.Count(&totalNum)
	table.Limit(pageSize).Offset(pageOffset)
	var list []dtos.PostsListDto
	table.Order("create_time desc").Find(&list)
	table.Find(&list)
	return &list, totalNum, nil
}

func (p *PostsMapper) SelectById(postId int64) (*model.Post, error) {
	var result model.Post
	r := p.db.Where("id=?", postId).First(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	if &result == nil {
		return nil, errors.New("post not found")
	}
	return &result, nil
}

func (p *PostsMapper) IsAuthorTo(id int64, userId int) bool {
	var count int64 = 0
	p.db.Model(&model.Post{}).Where("id=? and user_id=?", id, userId).Count(&count)
	return count > 0
}

// 创建
func (p *PostsMapper) SaveNew(id int64, userId int, title string, content string) (int64, error) {
	var t = time.Now()
	var tempBean = model.Post{Id: id, UserId: userId, Title: title, Content: &content, State: 1, RateNumber: 0, CreateTime: &t}
	r := p.db.Create(&tempBean)
	if r.Error != nil {
		return -1, r.Error
	}
	return id, nil
}

// 更新
func (p *PostsMapper) UpdateById(id int64, title string, content string, state int) (bool, error) {
	var t = time.Now()
	r := p.db.Debug().Model(&model.Post{}).Where("id=?", id).Updates(map[string]interface{}{
		"title":       title,
		"content":     content,
		"state":       state,
		"update_time": t,
	})
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

// 删除
func (p *PostsMapper) RemoveById(postId int64) (bool, error) {
	r := p.db.Where("id=?", postId).Delete(&model.Post{})
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}
