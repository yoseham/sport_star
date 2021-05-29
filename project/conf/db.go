package conf

const DriverName  = "mysql"

type DbConf struct {
	Host string
	Port int
	User string
	Pwd string
	DbName string
}

var MasterDbConfig DbConf = DbConf{
	"106.54.91.157",
	3306,
	"root",
	"07597321",
	"sportstar",
}

var SlaveDbConfig DbConf = DbConf{
	"106.54.91.157",
	3306,
	"root",
	"07597321",
	"sportstar",
}