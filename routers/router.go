package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
	"github.com/dchest/captcha"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Handler("/captcha/*.png",captcha.Server(100, 50))
}
