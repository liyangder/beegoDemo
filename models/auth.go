package models

import "fmt"

type Auth struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Username string `json:"username" xorm:"default '' comment('账号') VARCHAR(50)"`
	Password string `json:"password" xorm:"default '' comment('密码') VARCHAR(50)"`
}

func (m *Auth) Find() map[int]Auth {
	users := make(map[int]Auth)
	err := dbm.Id(2).Where("id > ?", 1).Desc("id").Find(&users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
	return users
	//fmt.Println(err)
	//fmt.Println("sjsj")
}
