package models

import (
	_ "github.com/astaxie/beego/orm"
)

type Auth struct {
	Id       int
	Username string
	Password string
}

func (*Auth) TableName() string {
	TableName("auth")
}
