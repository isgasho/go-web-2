package initialize

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go-web/common"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

// ZapLocalTimeEncoder 定义日志格式时间
func ZapLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// 格式化时间成字符串
	enc.AppendString(t.Format(common.MsecLocalTimeFormat))
}

func Logger() {
	// 生成日志文件完整路径
	now := time.Now()
	logFile := fmt.Sprintf("%s/%s.%04d-%02d-%02d.log", common.Config.Log.Path, common.Config.System.ServiceName, now.Year(), now.Month(), now.Day())

	// 配置日志切割属性
	hook := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    common.Config.Log.MaxSize,
		MaxAge:     common.Config.Log.MaxAge,
		MaxBackups: common.Config.Log.MaxBackups,
		Compress:   common.Config.Log.Compress,
	}
	defer hook.Close()

	// 配置 Zap 的时间输出
	enConfig := zap.NewProductionEncoderConfig()
	enConfig.EncodeTime = ZapLocalTimeEncoder

	// 日志颜色处理
	if common.Config.Log.Colorful {
		enConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		enConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	// 日志输出位置设置
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
		common.Config.Log.Level,
	)

	// 处理日志输出中函数文件问题
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// 配置全局
	common.Logger = logger.Sugar()

	// 输出信息
	log.Println("日志配置初始化完成！")
}
