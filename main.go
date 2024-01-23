package main

import (
	"fmt"

	"github.com/wolfitem/gcommon/module/config"
	"github.com/wolfitem/gcommon/module/log"


)

func init(){
	config.Init()

	log.Init()
}

func main() {



	fmt.Println("end")

}
