package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Index() {

	this.SetSession("username","brady")
	username := this.GetSession("username")
	logs.Info(username)
	this.Ctx.WriteString("3333")
}
