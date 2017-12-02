package main

import (
	_ "github.com/a1411350954/web/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:123@/web?charset=utf8&loc=Local")
	orm.RunSyncdb("default", false, true)
	beego.SetLevel(beego.LevelDebug)

}

func main() {

	beego.Run()
}

