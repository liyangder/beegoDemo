package models

import "fmt"

type Auth struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Username string `json:"username" xorm:"default '' comment('账号') VARCHAR(50)"`
	Password string `json:"password" xorm:"default '' comment('密码') VARCHAR(50)"`
}

func (m *Auth) Get() []Auth {
	users := make([]Auth, 0)
	err := dbm.Where("id > ?", 100).Desc("id").Find(&users)
	if err != nil {
		fmt.Println(err)
	}
	return users
}
