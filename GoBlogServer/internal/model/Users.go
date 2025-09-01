package model

import "time"

type Users struct {
	Id         int       `gorm:"primary_key;AUTO_INCREMENT"`
	UserName   string    `gorm:"type:varchar(100);"`
	Password   string    `gorm:"type:varchar(255);"`
	Email      string    `gorm:"type:varchar(50);"`
	State      int       `gorm:"type:tinyint(1);"`
	CreateTime time.Time `gorm:"type:datetime;"`
	LastLogin  time.Time `gorm:"type:datetime;"`
}

// 自定义表明
func (u *Users) TableName() string {
	return "blog_users"
}
