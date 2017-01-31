package config

import (
	base "github.com/wolfitem/gcommon/module/base"

	"log"
	"github.com/kylelemons/go-gypsy/yaml"

	"strings"


)

const(
	config_file="app.config"


)
const config_default_context=
`
version: 1

log:
	log_level: error
	log_file_dic: $LOG_FILE_DIC

app_setting:

`


func Init(){
	config_file_path:=base.GetCurrentDirectory()+"/"+config_file

	if !base.CheckFileIsExist(config_file_path){

		log.Printf("create config file : %s ",config_file_path)
		config_context:=strings.Replace(config_default_context,"$LOG_FILE_DIC",base.GetCurrentDirectory(),-1)

		base.WriteFile(config_file_path,config_context)
	}
}

var configInstance *yaml.File

/**
获取配置对象
 */
func GetConfig() *yaml.File  {
	if configInstance !=nil {
		return configInstance
	}
	config_file_path:=base.GetCurrentDirectory()+"/"+config_file

	if !base.CheckFileIsExist(config_file_path){
		base.CreateFile(config_file_path)
	}

	config, err := yaml.ReadFile(config_file_path)
	if err != nil {
		log.Fatalf("read file(%q): %s", config_file_path, err)
	}
	version,_:=config.Get("version")
	log.Printf("read config file : %s ,version : %s \n",config_file_path,version)
	configInstance=config
	return config
}
