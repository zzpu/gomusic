package main

import (
	_ "gomusic/routers"
	"github.com/astaxie/beego"
	"gomusic/daemon"
	"gomusic/models"
)

func main() {

	daemon.Init()
        models.Init()
	beego.Run()
}

