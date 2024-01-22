package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"wolfitem.com/gcommon/module/config"

	"wolfitem.com/gcommon/module/base"
)

func Debug(format string, a ...interface{}) {
	logrus.Debugf(format, a...)
}

func Info(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	logrus.Infof(format, a...)
}
func Warn(format string, a ...interface{}) {
	logrus.Warnf(format, a...)
}
func Error(format string, a ...interface{}) {
	logrus.Errorf(format, a...)
}

func Panic(format string, a ...interface{}) {
	logrus.Panicf(format, a...)
}

func Init() {

	// Output to stderr instead of stdout, could also be a file.
	logrus.SetOutput(os.Stdout)

	logLevel, err := config.GetConfig().Get("log.log_level")
	if err != nil {
		fmt.Printf(err.Error())
	}
	if logLevel == "" {
		logLevel = "error"
	}
	set_log_level(logLevel)

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})


	//添加Hook，用于本地存储日志文件或调用远程存储日志接口
	add_hook()

}

func add_hook() {
	log_file_dic, _ := config.GetConfig().Get("log.log_file_dic")

	if log_file_dic == "" {
		return
	}
	base.CreateFile(log_file_dic + "/error.log")
	base.CreateFile(log_file_dic + "/info.log")
	base.CreateFile(log_file_dic + "/panic.log")
	base.CreateFile(log_file_dic + "/debug.log")
	logrus.AddHook(lfshook.NewHook(lfshook.PathMap{
		logrus.ErrorLevel: log_file_dic + "/error.log",
		logrus.InfoLevel:  log_file_dic + "/info.log",
		logrus.PanicLevel: log_file_dic + "/panic.log",
		logrus.DebugLevel: log_file_dic + "/debug.log",
	}))

}

func set_log_level(level string) {
	switch strings.ToLower(level) {
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}
}
