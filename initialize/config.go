package initialize

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/viper"
	"go-web/common"
	"log"
	"os"
)

// 创建相关基础常量
const (
	configType    = "yml"
	configDir     = "config"
	configPrefix  = "application"
	configDefault = "dev"
)

// ConfigBox 存储配置文件的盒子
type ConfigBox struct {
	Fs  embed.FS // 读取多文件方式
	Dir string   // 保存目录，用于获取完整路径
}

// Get 传入文件名称，获取指定文件数据
func (c ConfigBox) Get(filename string) (bs []byte) {
	// 拼接获取完整路径
	f := fmt.Sprintf("%s%s%s", c.Dir, "/", filename)

	// 先从本地获取，用于本地覆盖代码中打包的配置
	bs, err := os.ReadFile(f)
	if err != nil {
		// 本地没有文件是正常的，这里只是一个提示作用，同时将 err 设置为 nil，避免影响后面判断
		log.Println("没有本地配置文件，将使用默认配置！")
		err = nil
	}

	// 确定本地配置没有覆盖的配置之后，读取embed打包中的默认配置
	if len(bs) == 0 {
		bs, err = c.Fs.ReadFile(f)
		if err != nil {
			// 此时就可能读取配置出现问题，如果出现问题则抛出异常
			log.Fatalln("从embed中读取配置文件失败：", err.Error())
		}
	}

	// 如果都没有问题，则把读取到的数据返回
	return
}

// ReadConfig viper 读取配置文件
func ReadConfig(box ConfigBox, v *viper.Viper, filename string) {
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
func Config(fs embed.FS, runEnv string) {
	// 初始化配置文件盒子
	var box ConfigBox
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
	// runEnv := strings.ToLower(os.Getenv(configEnvName))
	// 通过用户传递启动参数
	if runEnv != "" && runEnv != configDefault {
		// 再次读取需要的配置文件
		configName := fmt.Sprintf("%s-%s.%s", configPrefix, runEnv, configType)
		ReadConfig(box, v, configName)
	}

	// 将最终的配置文件传递给全局使用
	err := v.Unmarshal(&common.Config)
	if err != nil {
		log.Fatal("配置文件初始化失败：", err.Error())
	}
	log.Println("配置文件初始化完成，初始化环境为：", runEnv)
}
