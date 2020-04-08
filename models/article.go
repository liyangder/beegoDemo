package models

type Article struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	TagId         int    `json:"tag_id" xorm:"default 0 comment('标签ID') INT(10)"`
	Title         string `json:"title" xorm:"default '' comment('文章标题') VARCHAR(100)"`
	Desc          string `json:"desc" xorm:"default '' comment('简述') VARCHAR(255)"`
	Content       string `json:"content" xorm:"comment('内容') TEXT"`
	CoverImageUrl string `json:"cover_image_url" xorm:"default '' comment('封面图片地址') VARCHAR(255)"`
	CreatedOn     int    `json:"created_on" xorm:"default 0 comment('新建时间') INT(10)"`
	CreatedBy     string `json:"created_by" xorm:"default '' comment('创建人') VARCHAR(100)"`
	ModifiedOn    int    `json:"modified_on" xorm:"default 0 comment('修改时间') INT(10)"`
	ModifiedBy    string `json:"modified_by" xorm:"default '' comment('修改人') VARCHAR(255)"`
	DeletedOn     int    `json:"deleted_on" xorm:"default 0 INT(10)"`
	State         int    `json:"state" xorm:"default 1 comment('删除时间') TINYINT(3)"`
}
