package config

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

func InitConfig() {
	pwd, _ := os.Getwd()
	viper.AddConfigPath(pwd + "/config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func InitLogger(logger *logrus.Logger) {
	if logger != nil {
		return
	}

	path := "./catdfs_log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H",
		rotatelogs.WithLinkName(path),              // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(24*time.Hour),        // 文件最大保存时间
		rotatelogs.WithRotationTime(1*time.Minute), // 日志切割时间间隔
	)

	pathMap := lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
		logrus.FatalLevel: writer,
	}
	logger = logrus.New()
	logger.Hooks.Add(lfshook.NewHook(pathMap, &logrus.TextFormatter{}))
}
