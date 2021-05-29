package datasource

import (
	"app/project/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

var(
	masterEngine *xorm.Engine
	slaveEngine *xorm.Engine
	lock sync.Mutex
)
//单例模式
func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)

	//优化性能，加上本机SQL缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	if err != nil {
		log.Fatalf("dbhelper.DbInstanceMaster err:%s", err)
	} else {
		masterEngine = engine
	}
	return masterEngine
}

func InstanceSlave()(slaveEngine *xorm.Engine) {
	if slaveEngine != nil {
		return
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatalf("dbhelper.DbInstanceSlave err:%s", err)
	} else {
		slaveEngine = engine
	}
	return
}