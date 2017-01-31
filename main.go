package main

import (
	"fmt"

	"github.com/wolfitem/gcommon/module/config"
	"github.com/wolfitem/gcommon/module/log"


)

func main() {

	config.Init()

	log.Init()

	log.Debug("debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")

	fmt.Println("end")

}
