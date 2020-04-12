package models

import (
	"fmt"
	"reflect"
)

type Auth struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Username string `json:"username" xorm:"default '' comment('账号') VARCHAR(50)"`
	Password string `json:"password" xorm:"default '' comment('密码') VARCHAR(50)"`
}

func (m *Auth) Get() []Auth {
	data := make([]Auth, 0)
	err := dbm.Where("id > ?", 1).Desc("id").Find(&data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func (m *Auth) Add(data []Auth) interface{} {

	affected, err := dbm.Insert(&data)
	if err != nil {
		fmt.Println(err)
	}
	if affected < 1 {
		return "添加失败"
	}
	return true

}

func (m *Auth) Update(id int64) interface{} {
	affected, err := dbm.Exec("update auth set username = ? where id = ?", "wowsso", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(affected))

	return true
}

func (m *Auth) Delete(id int64) interface{} {
	affected, err := dbm.Id(id).Delete(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reflect.TypeOf(affected))

	return true
}
