package model

import "time"

type Post struct {
	Id         int64      `gorm:"type:bigint;primary_key;AUTO_INCREMENT:false" json:"id"`
	Title      string     `gorm:"type:varchar(255);comment:标题" json:"title"`
	Content    *string    `gorm:"type:text;comment:内容" json:"content"`
	UserId     int        `gorm:"type:int;comment:所属用户" json:"userId"`
	RateNumber int        `gorm:"type:int;comment:评论数" json:"rateNumber"`
	CreateTime *time.Time `gorm:"type:datetime" json:"createTime"`
	UpdateTime *time.Time `gorm:"type:datetime" json:"updateTime"`
	State      int        `gorm:"type:tinyint;comment:状态" json:"state"`
}

// 自定义表名
func (p *Post) TableName() string {
	return "blog_posts"
}
