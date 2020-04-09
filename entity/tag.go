package entity

type Tag struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name       string `json:"name" xorm:"default '' comment('标签名称') VARCHAR(100)"`
	CreatedOn  int    `json:"created_on" xorm:"default 0 comment('创建时间') INT(10)"`
	CreatedBy  string `json:"created_by" xorm:"default '' comment('创建人') VARCHAR(100)"`
	ModifiedOn int    `json:"modified_on" xorm:"default 0 comment('修改时间') INT(10)"`
	ModifiedBy string `json:"modified_by" xorm:"default '' comment('修改人') VARCHAR(100)"`
	DeletedOn  int    `json:"deleted_on" xorm:"default 0 comment('删除时间') INT(10)"`
	State      int    `json:"state" xorm:"default 1 comment('状态 0为禁用、1为启用') TINYINT(3)"`
}
