package container

import "gorm.io/gorm"

// bean容器
type Container struct {
	db *gorm.DB
}
