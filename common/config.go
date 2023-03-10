package common

import "go.uber.org/zap/zapcore"

// Configuration 总配置
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" json:"system"`
	Log    LogConfiguration    `mapstructure:"log" json:"log"`
	Mysql  MysqlConfiguration  `mapstructure:"mysql" json:"mysql"`
}

// SystemConfiguration 系统配置
type SystemConfiguration struct {
	ServiceName string `mapstructure:"service-name" json:"service-name"`
	Port        int    `mapstructure:"port" json:"port"`
	Mode        string `mapstructure:"mode" json:"mode"`
	ApiPrefix   string `mapstructure:"api-prefix" json:"api-prefix"`
	ApiVersion  string `mapstructure:"api-version" json:"api-version"`
}

// LogConfiguration 日志相关配置
type LogConfiguration struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Colorful   bool          `mapstructure:"colorful" json:"colorful"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"max-size"`
	MaxBackups int           `mapstructure:"max-backups" json:"max-backups"`
	MaxAge     int           `mapstructure:"max-age" json:"max-age"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

// MysqlConfiguration MySQL 配置
type MysqlConfiguration struct {
	Host               string `mapstructure:"host" json:"host"`
	Port               string `mapstructure:"port" json:"port"`
	Database           string `mapstructure:"database" json:"database"`
	Username           string `mapstructure:"username" json:"username"`
	Password           string `mapstructure:"password" json:"password"`
	Charset            string `mapstructure:"charset" json:"charset"`
	Collation          string `mapstructure:"collation" json:"collation"`
	TablePrefix        string `mapstructure:"table-prefix" json:"table-prefix"`
	Query              string `mapstructure:"query" json:"query"`
	MaxIdleConnections int    `mapstructure:"max-idle-connections" json:"max-idle-connections"`
	MaxOpenConnections int    `mapstructure:"max-open-connections" json:"max-open-connections"`
	MaxIdleTime        int    `mapstructure:"max-idle-time" json:"max-idle-time"`
	LogMode            bool   `mapstructure:"log-mode" json:"log-mode"`
	LogLevel           int    `mapstructure:"log-level" json:"log-level"`
}
