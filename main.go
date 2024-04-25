package main

import (
	"github.com/gogf/gf/frame/g"
)

func main() {
	s := g.Server()
	//初始化本地数据库
	//读取系统设置对象
	//初始化路由管理器
	//初始化插件管理器
	//初始化脚本管理器
	s.Run()
}
