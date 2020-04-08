package main

import (
	_ "beegoDemo/models"
	_ "beegoDemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
