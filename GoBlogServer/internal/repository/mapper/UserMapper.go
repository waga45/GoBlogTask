package mapper

import "gorm.io/gorm"

type UserMapper struct {
	db *gorm.DB
}

func NewUserMapper(db *gorm.DB) *UserMapper {
	return &UserMapper{db: db}
}
