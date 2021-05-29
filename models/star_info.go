package models

type StarInfo struct {
	Id           int    `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	NameZh       string `xorm:"not null default '' comment('中文名') VARCHAR(50)"`
	NameEn       string `xorm:"not null default '' comment('英文名') VARCHAR(50)"`
	Avatar       string `xorm:"not null default '' comment('头像') VARCHAR(255)"`
	Birthday     string `xorm:"not null default '' comment('出生日期') VARCHAR(50)"`
	Height       int    `xorm:"not null default 0 comment('身高，单位cm') INT(10)"`
	Weight       int    `xorm:"not null default 0 comment('体重，单位kg') INT(10)"`
	Club         string `xorm:"not null default '' comment('俱乐部') VARCHAR(50)"`
	Jersy        string `xorm:"not null default '' comment('球衣号码以及主打位置') VARCHAR(50)"`
	Country      string `xorm:"not null default '' comment('国籍') VARCHAR(50)"`
	Birthaddress string `xorm:"not null default '' comment('出生地') VARCHAR(255)"`
	Feature      string `xorm:"not null default '' comment('个人特点') VARCHAR(255)"`
	Moreinfo     string `xorm:"comment('更多介绍') TEXT"`
	SysStatus    int    `xorm:"not null default 0 comment('状态，默认值 0 正常，1 删除') TINYINT(4)"`
	SysCreated   int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated   int    `xorm:"not null default 0 comment('最后修改时间') INT(10)"`
}
