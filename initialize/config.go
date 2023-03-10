package initialize

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"go-web/common"
	"go-web/pkg/dto"
	"log"
	"os"
	"strings"
)

// 创建相关基础常量
const (
	configEnvName = "RUN_ENV"
	configType    = "yml"
	configDir     = "config"
	configPrefix  = "application"
	configDefault = "dev"
)

// ReadConfig viper 读取配置文件
func ReadConfig(box dto.ConfigBox, v *viper.Viper, filename string) {
	// 设置需要读取的配置文件类型
	v.SetConfigType(configType)

	// 获取指定文件的数据
	config := box.Get(filename)

	// 判断是否获取到数据，如果读取到的内容为空，则抛出异常
	if len(config) == 0 {
		log.Fatal("配置文件没有获取到相应的内容：", filename)
	}

	// viper 加载配置
	err := v.ReadConfig(bytes.NewReader(config))
	if err != nil {
		log.Fatal("使用viper配置文件加载失败：", err.Error())
	}
}

// Config 初始化配置文件
func Config(fs embed.FS) {
	// 初始化配置文件盒子
	var box dto.ConfigBox
	// 需要把 config 下面的所有配置都传过来
	box.Fs = fs
	// 保存定义的配置目录，用于拼接完整路径
	box.Dir = configDir

	// 拼接默认配置文件名称，默认为 dev
	defaultConfigName := fmt.Sprintf("%s-%s.%s", configPrefix, configDefault, configType)

	// 使用 viper 读取配置文件，将默认配置先加载进去
	v := viper.New()
	ReadConfig(box, v, defaultConfigName)
	settings := v.AllSettings()
	for index, setting := range settings {
		v.SetDefault(index, setting)
	}

	// 获取运行模式，通过运行方式判断实际需要加载的配置文件
	runEnv := strings.ToLower(os.Getenv(configEnvName))
	if runEnv != "" && runEnv != configDefault {
		// 再次读取需要的配置文件
		configName := fmt.Sprintf("%s-%s.%s", configPrefix, runEnv, configType)
		ReadConfig(box, v, configName)
	} else {
		runEnv = configDefault
	}

	// 将最终的配置文件传递给全局使用
	err := v.Unmarshal(&common.Config)
	if err != nil {
		log.Fatal("配置文件初始化失败：", err.Error())
	}
	log.Println("配置文件初始化成功，初始化环境为：", runEnv)
}
