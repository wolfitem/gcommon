package log


import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/wolfitem/gcommon/module/config"
	"time"
	"strings"

)


const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta //洋红
	info          = "[INFO]"
	trac          = "[TRAC]"
	erro          = "[ERRO]"
	warn          = "[WARN]"
	succ          = "[SUCC]"
)

// see complete color rules in document in https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13
func Debug(format string, a ...interface{}) {
	msg:=fmt.Sprintf(format, a...)

	prefix := yellow(trac)
	fmt.Println(formatLog(prefix), msg)

	log.Debugln(msg)

}
func Info(format string, a ...interface{}) {
	msg:=fmt.Sprintf(format, a...)

	prefix := blue(info)
	fmt.Println(formatLog(prefix), msg)

	log.Infoln(msg)
}
func Success(format string, a ...interface{}) {
	msg:=fmt.Sprintf(format, a...)

	prefix := green(succ)
	fmt.Println(formatLog(prefix), msg)

	log.Infoln(msg)
}
func Warn(format string, a ...interface{}) {
	msg:=fmt.Sprintf(format, a...)

	prefix := magenta(warn)
	fmt.Println(formatLog(prefix), msg)

	log.Warnln(msg)

}
func Error(format string, a ...interface{}) {
	msg:=fmt.Sprintf(format, a...)

	prefix := red(erro)
	fmt.Println(formatLog(prefix), msg)

	log.Errorln(msg)
}
func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, s)
}
func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}
func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}
func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}
func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}
func formatLog(prefix string) string {
	return time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " "
}

func Init() {

	configInstance:=config.GetConfig()

	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	logLevel,_:=configInstance.Get("log_level")
	set_log_level(logLevel)
	//添加Hook，用于本地存储日志文件或调用远程存储日志接口
	add_hook()

}

func add_hook(){
	configInstance:=config.GetConfig()
	log_file_dic,_:=configInstance.Get("log_file_dic")
	if log_file_dic==""{
		return
	}
	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.ErrorLevel : log_file_dic +"/error.log",
		log.InfoLevel : log_file_dic +"/info.log",
	}))
}

func set_log_level(level string){
	switch strings.ToLower(level) {
	case "error":
		log.SetLevel(log.ErrorLevel)
		break
	case "warn":
		log.SetLevel(log.WarnLevel)
		break
	case "info":
		log.SetLevel(log.InfoLevel)
		break
	case "debug":
		log.SetLevel(log.DebugLevel)
		break
	default:
		log.SetLevel(log.ErrorLevel)
		break
	}
}

