package model

import "time"

type Comments struct {
	Id         int        `gorm:"primary_key;AUTO_INCREMENT"`
	Content    string     `gorm:"type:text;comment:评论内容"`
	UserId     int        `gorm:"type:int;comment:评论用户ID"`
	PostId     int        `gorm:"type:bigint"`
	CreateTime *time.Time `gorm:"type:datetime"`
	State      int        `gorm:"type:tinyint"`
}

// 评论
func (c *Comments) TableName() string {
	return "blog_comments"
}
