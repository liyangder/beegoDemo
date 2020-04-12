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
	res := oauth.Get()
	c.out_success(res)

}
func (c *AuthContr) AddUser() {
	oauth := new(models.Auth)
	data := make([]models.Auth, 1)

	data[0].Username = "李五"
	data[0].Password = "李adsfa"

	oauth.Add(data)
}

func (c *AuthContr) UpdateUser() {
	oauth := new(models.Auth)
	oauth.Update(1)
}

func (c *AuthContr) Delete_user() {
	oauth := new(models.Auth)
	oauth.Delete(5)
}
