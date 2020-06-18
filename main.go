package main

import (
	_ "xcms/routers"
	_ "xcms/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}