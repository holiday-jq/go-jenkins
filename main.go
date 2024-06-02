package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"net/http"
	"os"
)

var logger *zap.Logger

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")
	logger.Info("hahah")
	logger.Error("err")
	logger.Warn("warn")
}

func main() {
	LogInit()
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/hello", index)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func LogInit() {
	logFileName := "app.log"
	logDir := "./logs"
	fullPath := fmt.Sprintf("%s/%s", logDir, logFileName)
	// 初始化文件IO，用于zap写入日志

	fileOut := &lumberjack.Logger{
		Filename:   fullPath, // 日志文件路径
		MaxSize:    1024,     // 单个文件最大大小（MB）
		MaxBackups: 100,      // 保留旧文件的最大数量
		MaxAge:     7,        // 旧文件保留天数
		Compress:   true,     // 是否压缩旧文件
	}
	mulWriter := io.MultiWriter(os.Stdout, fileOut)

	fileWriteSyncer := zapcore.AddSync(mulWriter)
	// 定义日志encoder配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// 构建core
	core := zapcore.NewCore(encoder, fileWriteSyncer, zap.InfoLevel)
	// 构建logger
	logger = zap.New(core)
	// 记录日志
	logger.Info("Starting the service...")
}
