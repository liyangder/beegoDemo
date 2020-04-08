package routers

import (
	"beegoDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.IndexContr{})
	beego.AutoRouter(&controllers.AuthContr{})
}
