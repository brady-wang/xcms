package main

import (
	"github.com/astaxie/beego/logs"
	_ "xcms/routers"
	_ "xcms/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	logs.Async()
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/test.log"}`)
	logs.SetLevel(logs.LevelInformational)     // 设置日志写入缓冲区的等级
	beego.Run()
}