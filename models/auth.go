package models

type Auth struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Username string `json:"username" xorm:"default '' comment('账号') VARCHAR(50)"`
	Password string `json:"password" xorm:"default '' comment('密码') VARCHAR(50)"`
}
