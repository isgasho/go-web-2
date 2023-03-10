package dto

import (
	"embed"
	"fmt"
	"log"
	"os"
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
