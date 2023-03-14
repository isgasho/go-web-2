# gin-web

#### 介绍
基于 Go + Gin + Gorm + Casbin 开发的 RBAC Web 后台实现方式。


#### 开发依赖

```bash
# gin：Web 框架
go get -u github.com/gin-gonic/gin
# viper：YAML 配置文件读取
go get -u github.com/spf13/viper
# zap：日志处理
go get -u go.uber.org/zap
# lumberjack：日志切割
go get -u github.com/natefinch/lumberjack 
# gorm：MySQL 数据库操作
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
# carbon：时间库
go get -u github.com/golang-module/carbon/v2
# gin-jwt：JWT
go get -u github.com/appleboy/gin-jwt/v2
# go-redis：Redis
go get -u github.com/redis/go-redis/v9
```
