package controllers

import (
	"beegoDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexContr struct {
	beego.Controller
}

func (c *IndexContr) Index() {
	c.Ctx.WriteString("hello world")
}

func (c *IndexContr) GetUser() {
	o := orm.NewOrm()
	auth := models.Auth{Id: 1}
	err := o.Read(&auth)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(auth.Id, auth.Username)
	}
}
