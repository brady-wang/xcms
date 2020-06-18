package controllers

import (
	"github.com/astaxie/beego"
)

type FormatController struct {
	beego.Controller
}

func (this *v) Index() {
	this.Ctx.WriteString("登录")
}

func (this *v) Add() {
	this.Ctx.WriteString("登录")
}

func (this *v) AddDo() {
	this.Ctx.WriteString("登录")
}

func (this *v) Edit() {
	this.Ctx.WriteString("登录")
}

func (this *UserController) EditDo() {
	this.Ctx.WriteString("登录")
}

func (this *UserController) DeleteDo() {
	this.Ctx.WriteString("登录")
}

func (this *UserController) List() {
	this.Ctx.WriteString("登录")
}
