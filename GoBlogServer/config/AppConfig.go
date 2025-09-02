package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	/**
	服务配置
	*/
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	/**
	数据库
	*/
	Database struct {
		Mysql struct {
			Url      string `yaml:"url"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}
	}
	/**
	redis 配置
	*/
	Redis struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Auth         string `yaml:"auth"`
		ReadTimeOut  int    `yaml:"read_timeout"`
		WriteTimeOut int    `yaml:"write_timeout"`
		MinIdleConn  int    `yaml:"min_idle_conn"`
	}
	//wt
	Jwt struct {
		SecretKey  string `yaml:"secret_key"`  //秘钥
		ExpireHour int    `yaml:"expire_hour"` //有效期 小时
	}
}

/*
*
加载配置
*/
func LoadConfig(configPath string) (*AppConfig, error) {
	appConfig := &AppConfig{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	err = d.Decode(&appConfig)
	if err != nil {
		return nil, err
	}
	return appConfig, nil
}
