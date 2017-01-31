package log


import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/wolfitem/gcommon/module/config"
	"strings"

	"github.com/wolfitem/gcommon/module/base"
)

func Debug(format string, a ...interface{}) {
	log.Debugf(format, a...)
}

func Info(format string, a ...interface{}) {
	log.Infof(format, a...)
}
func Warn(format string, a ...interface{}) {
	log.Warnf(format, a...)
}
func Error(format string, a ...interface{}) {
	log.Errorf(format, a...)
}

func Panic(format string, a ...interface{}) {
	log.Panicf(format, a...)
}



func Init() {


	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stdout)

	logLevel,_:=config.GetConfig().Get("log.log_level")
	if logLevel==""{
		logLevel="error"
	}
	set_log_level(logLevel)

	//添加Hook，用于本地存储日志文件或调用远程存储日志接口
	add_hook()

}

func add_hook(){
	log_file_dic,_:=config.GetConfig().Get("log.log_file_dic")
	if log_file_dic==""{
		return
	}
	base.CreateFile(log_file_dic +"/error.log")
	base.CreateFile(log_file_dic +"/info.log")
	base.CreateFile(log_file_dic +"/panic.log")
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.ErrorLevel : log_file_dic +"/error.log",
		log.InfoLevel : log_file_dic +"/info.log",
		log.PanicLevel:log_file_dic +"/panic.log",
	}))
}

func set_log_level(level string){
	switch strings.ToLower(level) {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}
}

