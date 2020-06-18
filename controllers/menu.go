package controllers

import (
	"github.com/astaxie/beego"
)

type MenuController struct {
	beego.Controller
}

func (this *MenuController) Index() {
	this.Ctx.WriteString("登录")
}

func (this *MenuController) Add() {
	this.Ctx.WriteString("登录")
}

func (this *MenuController) AddDo() {
	this.Ctx.WriteString("登录")
}

func (this *MenuController) Edit() {
	this.Ctx.WriteString("登录")
}

func (this *MenuController) EditDo() {
	this.Ctx.WriteString("登录")
}

func (this *MenuController) DeleteDo() {
	this.Ctx.WriteString("登录")
}

func (this *MenuController) List() {
	this.Ctx.WriteString("登录")
}
