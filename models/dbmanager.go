//数据库连接管理，缓存连接管理也会在这里
package models

import (
	"fork/tools/log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//DBaddress : used db connection
var DBaddress string

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", DBaddress, 5, 5)
	log.Info("start init database ")
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/fork", 10, 10)

	log.Info("end init database ")
	orm.Debug = true
}

func RegisterModels(model interface{}) {
	if model == nil {
		return
	}
	orm.RegisterModel(model)
	orm.RunSyncdb("default", true, true)
}
