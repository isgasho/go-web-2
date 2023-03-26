package initialize

import (
	"embed"
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go-web/common"
	"log"
)

// Casbin 初始化
func Casbin(fs embed.FS) {
	// 初始化数据库适配器
	adapter, err := gormadapter.NewAdapterByDB(common.DB)
	if err != nil {
		log.Fatal("初始化 Casbin gorm 适配器失败：", err.Error())
		return
	}

	// 获取 model.conf 配置
	var box ConfigBox
	box.Fs = fs
	box.Dir = configDir
	bs := box.Get("model.conf")
	config := string(bs[:])

	// 从字符串中加载配置
	m, err := casbinmodel.NewModelFromString(config)
	if err != nil {
		log.Fatal("初始化 Casbin 配置文件失败：", err.Error())
		return
	}

	// 读取配置文件
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		log.Fatal("创建 Casbin Enforcer 失败：", err.Error())
		return
	}

	// 加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Fatal("加载 Casbin 策略失败：", err.Error())
		return
	}

	// 配置全局
	common.CasbinEnforcer = enforcer
	log.Println("Casbin 权限控制初始化完成！")
}
