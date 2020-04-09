package controllers

import (
	"beegoDemo/models"
)

type AuthContr struct {
	BaseContr
}

func (c *AuthContr) Index() {
	c.Ctx.WriteString("hello world")
}

func (c *AuthContr) GetUser() {
	oauth := new(models.Auth)
	res := oauth.Find()
	c.out_success(res)

}
