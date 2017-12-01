package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

func init(){
	orm.RegisterModel(new(SysUser))
}

type SysUser struct{
	//id
	Id int					`orm:"auto;pk"`
	//用户名
	Username string			`orm:"column(user_name);size(30);unique"`
	//密码
	Password string			`orm:"column(pass_word);size(64)"`
	//昵称
	Name string				`orm:"column(name);size(50);null"`
	//email
	Email string			`orm:"column(email);size(50);null"`
	//性别
	Sex rune				`orm:"column(sex);null"`
	//连续输入密码错误次数
	ContinueErrorCount int	`orm:"column(continue_error_count);default(0)"`
	//状态（0：正常，1：禁用）
	Status int				`orm:"column(status);default(0)"`
	//创建时间
	CreateTime time.Time	`orm:"auto_now_add;column(create_time);type(datetime)"`
	//更新时间
	UpdateTime time.Time	`orm:"auto_now;column(update_time);type(datetime)"`
	//创建者
	CreateBy string			`orm:"column(create_by);size(50)"`
	//更新者
	UpdateBy string			`orm:"column(update_by);size(50)"`

}

func AddSysUser(u *SysUser) (int64,error){
	o := orm.NewOrm()

	id,err := o.Insert(u)
	if err !=nil{
		return 0,err
	}
	return id,nil
}

func ListSysUser(sort string,page int,pageSize int)(users []orm.Params,count int64){

	o := orm.NewOrm()

	var offset int
	if page<=1{
		offset = 0
	}else{
		offset = (page-1)*pageSize
	}

	count,err := o.QueryTable("sys_user").Limit(pageSize,offset).OrderBy(sort).Values(&users)

	if err!=nil{
		logs.Error(err)
	}

	return
}

func FindSysUserByUsername(username string)(user SysUser){
	o := orm.NewOrm()

	err := o.QueryTable(new(SysUser)).Filter("Username",username).One(&user)

	if err!=nil{
		logs.Error(err)
	}

	return
}

func UpdateSysUser(user *SysUser){
	o := orm.NewOrm()

	user.UpdateBy = "chengbiao"

	count,err := o.Update(user)

	if err!=nil{
		logs.Error(err)
	}
	logs.Info("更新SysUser完成，影响条数:%d",count)
}