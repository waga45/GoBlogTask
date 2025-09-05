package dtos

import "time"

type PostCommentDto struct {
	Id         int       `gorm:"column:id" json:"id"`
	UserName   string    `gorm:"column:userName" json:"userName"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:createTime;type:datetime" json:"createTime"`
}
