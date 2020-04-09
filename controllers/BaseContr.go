package controllers

import (
	"github.com/astaxie/beego"
)

type BaseContr struct {
	beego.Controller
}

func (c *BaseContr) out_success(res interface{}) {
	c.Data["json"] = res
	c.ServeJSON()
}
