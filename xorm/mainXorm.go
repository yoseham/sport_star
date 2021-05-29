package main

const DirverName = "mysql"
const MasterDataSourceName = ""

type UserInfo struct {
	Id	int `xorm:"not null pk autoincr"`
	Name	string
	SysCreated	int
	SysUpdated	int
}


func main() {

}