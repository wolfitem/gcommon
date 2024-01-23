package config

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
	base "github.com/wolfitem/gcommon/module/base"

	"strconv"
	"strings"
)

const (
	config_file = "goapp.config"
)
const config_default_context = `
version: 1
# 注意二级Key前面是四个空格，不能使用Tab
log:
    log_level: error
    log_file_dic: $LOG_FILE_DIC/

app_setting:

`

func Init() {
	config_file_path := base.GetCurrentDirectory() + "/" + config_file

	if !base.CheckFileIsExist(config_file_path) {
		config_context := strings.Replace(config_default_context, "$LOG_FILE_DIC", base.GetCurrentDirectory(), -1)
		base.WriteFile(config_file_path, config_context)
	}
}

var configInstance *yaml.File

/*
*
获取配置对象
*/
func GetConfig() *yaml.File {
	if configInstance != nil {
		return configInstance
	}
	config_file_path := base.GetCurrentDirectory() + "/" + config_file

	if !base.CheckFileIsExist(config_file_path) {
		base.CreateFile(config_file_path)
	}

	config, err := yaml.ReadFile(config_file_path)
	if err != nil {
		fmt.Printf("read file error.(%q): %s", config_file_path, err)
	}

	configInstance = config
	return config
}

/*
*
获取配置对象
*/
func Get_config_default_int(key string, default_vaule int) int {

	value, _ := GetConfig().Get(key)
	valueInt, _ := strconv.Atoi(base.TrimSpace(value))

	if valueInt <= 0 {
		return default_vaule
	}
	return valueInt
}

/*
*
获取配置对象
*/
func Get_config_default_bool(key string, default_vaule bool) bool {

	value, _ := GetConfig().Get(key)
	valueBool, err := strconv.ParseBool(base.TrimSpace(value))
	if err != nil {
		return default_vaule
	}
	return valueBool
}
