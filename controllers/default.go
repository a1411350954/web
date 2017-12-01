package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.Output.Header("Location","/login")
	c.Ctx.ResponseWriter.WriteHeader(302)
}
