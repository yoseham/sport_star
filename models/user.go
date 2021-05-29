package models

type User struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	Name       string `xorm:"not null VARCHAR(128)"`
	Syscreated int    `xorm:"not null INT(11)"`
	Sysupdated int    `xorm:"not null INT(11)"`
}
