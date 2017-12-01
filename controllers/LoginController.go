package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"web/models"
	"github.com/hzwy23/hauth/utils/hret"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get(){

	c.Data["CaptchaId"] = captcha.New()
	c.TplName = "login.html"
}


func (c *LoginController) Post(){

	username := c.GetString("username","")

	password := c.GetString("password","")

	user := models.FindSysUserByUsername(username)

	if user.Status!=0{
		hret.Error(c.Ctx.ResponseWriter,500,"账号被禁用!")
	}else if user.Password != password{
		hret.Error(c.Ctx.ResponseWriter,500,"账号或密码错误！")
		user.ContinueErrorCount++
		if user.ContinueErrorCount>=5{
			user.Status=1
		}
		models.UpdateSysUser(&user)
	}else{
		if user.ContinueErrorCount!=0{
			user.ContinueErrorCount=0
			models.UpdateSysUser(&user)
		}

		hret.Success(c.Ctx.ResponseWriter,"登录成功")
	}

}