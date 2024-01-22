package main

import (
	"fmt"

	"wolfitem.com/gcommon/module/config"
	"wolfitem.com/gcommon/module/log"


)

func init(){
	config.Init()

	log.Init()
}

func main() {



	fmt.Println("end")

}
