GoBlogServer 博客管理系统

架构：基于gin实现的轻量级mvc架构

一.涉及框架与技术：
1. Gin服务 
2. Zap 日志记录
3. yaml.v3 配置解析器
4. go-redis/redis/v8 Redis缓存
5. golang-jwt/jwt/v5 JWT令牌
6. mojocn/base64Captcha 图像验证码
7. gin-contrib/cors Cros跨域

二.配置
1. config/app_config.yaml  (包括服务端口，数据库链接等配置)

三.项目启动

运行 go run Application.go 即可。启动成功会自动创建数据库与表，无需手动创建

常用指令：
1. go clean -modcache
2. go mod tidy
3. go mod download

