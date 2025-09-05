package model

import (
	"gorm.io/gorm"
	"time"
)

type Comments struct {
	Id         int        `gorm:"primary_key;AUTO_INCREMENT"`
	Content    string     `gorm:"type:text;comment:评论内容"`
	UserId     int        `gorm:"type:int;comment:评论用户ID"`
	PostId     int64      `gorm:"type:bigint"`
	CreateTime *time.Time `gorm:"type:datetime"`
	State      int        `gorm:"type:tinyint"`
}

// 评论
func (c *Comments) TableName() string {
	return "blog_comments"
}

// 创建钩子
func (c *Comments) AfterCreate(tx *gorm.DB) error {
	tx.Model(&Post{}).Where("id =?", c.PostId).Update("rate_number", gorm.Expr("rate_number+?", 1))
	return nil
}

// 更新后看状态，如果无效Posts则数量-1 有效数量+1(这只是为了练习功能，本身逻辑会存在问题)
func (c *Comments) AfterUpdate(tx *gorm.DB) (err error) {
	if c.State == 0 {
		//-1
		tx.Model(&Post{}).Where("id =? and rate_number>0", c.PostId).Update("rate_number", gorm.Expr("rate_number-?", 1))
	} else if c.State == 1 {
		//+1
		tx.Model(&Post{}).Where("id =?", c.PostId).Update("rate_number", gorm.Expr("rate_number+?", 1))
	}
	return nil
}

// 删除
func (c *Comments) AfterDelete(tx *gorm.DB) error {
	tx.Model(&Post{}).Where("id =? and rate_number>0", c.PostId).Update("rate_number", gorm.Expr("rate_number-?", 1))
	return nil
}
