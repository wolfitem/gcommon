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
`version: 1
log:
	log_level: error
	log_file_dic: $LOG_FILE_DIC
app_setting:

`

//var config_instance config_yaml
//
//type config_yaml struct {
//	version string
//	log struct {
//		log_level string   `yaml:"log_level"`
//		log_file string   `yaml:"log_file"`
//	} `yaml:"log"`
//	app_settings struct {
//		app_setting[] struct {
//			key string
//			value string
//		}  `yaml:"app_setting"`
//	}
//}


func Init(){
	config_file_path:=base.GetCurrentDirectory()+"/"+config_file

	if !base.CheckFileIsExist(config_file_path){

		log.Printf("create config file : %s ",config_file_path)
		config_context:=strings.Replace(config_default_context,"$LOG_FILE_DIC",base.GetCurrentDirectory(),-1)

		base.WriteFile(config_file_path,config_context)
	}
}


/**
获取配置对象
 */
func GetConfig() *yaml.File  {

	config_file_path:=base.GetCurrentDirectory()+"/"+config_file
	log.Print("read file : "+config_file_path)
	config, err := yaml.ReadFile(config_file_path)
	if err != nil {
		log.Fatalf("read file(%q): %s", config_file_path, err)
	}
	//str,_:=config.Get("log_level")

	return config
}


/**
获取配置对象
 */
//func GetConfig() config_yaml {
//	if config_instance!=nil {
//		return config_instance
//	}
//	config_file_path:=base.GetCurrentDirectory()+"/"+config_file
//	log.Print("read file : "+config_file_path)
//
//	config := config_yaml{}
//
//	config_context:=base.ReadFile(config_file_path)
//
//	err := yaml.Unmarshal([]byte(config_context), &config)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	config_instance =config
//	return config_instance
//}