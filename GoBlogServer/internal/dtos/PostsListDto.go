package dtos

import "time"

type PostsListDto struct {
	Id         int64     `gorm:"column:id" json:"id"`
	Title      string    `gorm:"column:title" json:"title"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`
	State      int       `gorm:"state" json:"state"`
}

type PagePostsListDto struct {
	TotalCount int64           `json:"totalCount"`
	List       *[]PostsListDto `json:"list"`
}
