package common

// Configuration 总配置
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" json:"system"`
}

// SystemConfiguration 系统配置
type SystemConfiguration struct {
	Port       int    `mapstructure:"port" json:"port"`
	Mode       string `mapstructure:"mode" json:"mode"`
	ApiPrefix  string `mapstructure:"api-prefix" json:"api-prefix"`
	ApiVersion string `mapstructure:"api-version" json:"api-version"`
}
