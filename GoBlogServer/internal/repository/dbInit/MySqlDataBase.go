package dbInit

import (
	"GoBlogServer/config"
	"GoBlogServer/internal/repository/model"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"strings"
	"time"
)

/*
*
初始化数据库
*/
func InitMysql(config *config.AppConfig) (*gorm.DB, error) {
	databaseConfig := config.Database
	url := databaseConfig.Mysql.Url
	if len(url) == 0 {
		return nil, fmt.Errorf("database url is empty")
	}
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		errStr := fmt.Sprintf("%s", err.Error())
		if strings.Contains(errStr, "Unknown database") {
			db = autoCreateDatabase(databaseConfig.Mysql.Url)
		}
		if db == nil {
			return nil, err
		}
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(databaseConfig.Mysql.MaxIdleConn)
	sqlDB.SetMaxOpenConns(databaseConfig.Mysql.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(databaseConfig.Mysql.MaxIdleTime))
	err = sqlDB.Ping()
	if err != nil {
		zap.S().Error("database connection failed", zap.Error(err))
		return nil, err
	}
	//自动建表
	autoCreateTable(db)
	return db, nil
}

func autoCreateDatabase(url string) *gorm.DB {
	zap.S().Infof("数据库不存在，准创建...")
	re := regexp.MustCompile(`/([^/?]+)\?`)
	matches := re.FindStringSubmatch(url)
	if len(matches) <= 0 {
		panic("创建数据库失败，未找到数据库名称")
	}
	dbName := matches[1]
	copyUrl := string(url)
	copyUrl = strings.ReplaceAll(copyUrl, dbName, "")

	db, err := gorm.Open(mysql.Open(copyUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//执行数据库穿件
	err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;").Error
	if err != nil {
		panic(err)
	}
	//再次链接
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	return db
}

// 自动建表
func autoCreateTable(db *gorm.DB) {
	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Comments{})
}
