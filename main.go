package main

import (
	"time"
	"fmt"
	"test"
)
func init() {
	/*orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123@tcp(localhost:3306)/web?charset=utf8&loc=Local")

	orm.RunSyncdb("default",false,true)

	orm.Debug = true

	logs.SetLogger("console")*/

}
func main() {
	//beego.Run()
	//models.AddSysUser(&models.SysUser{Username:"admin",Password:"123456"})
	time1 := time.Date(2017,11,30,11,15,0,0,time.Local)
	time2 := time.Date(2017,11,30,11,15,29,0,time.Local)

	time3 := time.Date(1970,1,1,0,0,0,0,time.Local)

	//time.Year()

	//fmt.Println(time.Format("2006-01-02 15:04:05"))

	//1512011461   1512011478





	fmt.Println((time1.Unix()-time3.Unix())/30)
	fmt.Println((time2.Unix()-time3.Unix())/30)
}
