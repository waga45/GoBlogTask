package model

import "time"

type Post struct {
	Id         int        `gorm:"type:bigint;primary_key;AUTO_INCREMENT:false"`
	Title      string     `gorm:"type:varchar(255);comment:标题"`
	Content    *string    `gorm:"type:text;comment:内容"`
	UserId     int        `gorm:"type:int;comment:所属用户"`
	RateNumber int        `gorm:"type:int;comment:评论数"`
	CreateTime *time.Time `gorm:"type:datetime"`
	UpdateTime *time.Time `gorm:"type:datetime"`
	State      int        `gorm:"type:tinyint;comment:状态"`
}

// 自定义表名
func (p *Post) TableName() string {
	return "blog_posts"
}
