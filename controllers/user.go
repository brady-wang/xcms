package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.Ctx.WriteString("登录")
}

func (this *UserController) Add() {
	this.Ctx.WriteString("登录")
}

func (this *UserController) AddDo() {
	this.Ctx.WriteString("登录")
}

func (this *UserController) Edit() {
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
