package test

import (
	cc "GoBlocks/Components"
	"fmt"
)

func TestRM() {
	var mRM cc.TRouteManager
	mRM.Init()
	mRM.ConfDir = cc.Default_Conf_Dir
	fmt.Println("添加路由项")
	mRM.AddItem(cc.Route_Lua, "/hello1", "hello1")
	mRM.AddItem(cc.Route_Plugin, "/hello2", "hello2")
	mRM.AddItem(cc.Route_Proxy, "/hello3", "http://www.baidu.com")
	fmt.Println("所有路由项:::::", mRM.Items)
	fmt.Println("保存到文件")
	mRM.SaveItemsToFile(cc.Routes_Conf_Subdir) //保存到子目录下
	fmt.Println("保存完毕，看conf目录")
	fmt.Println("重新初始化")
	mRM.Init()
	fmt.Println("列出所有Conf文件")
	mFiles := mRM.AllConfFiles()
	fmt.Println(mFiles)
	fmt.Println("从文件加载路由项")
	for _, v := range mFiles {
		_ = mRM.AddItemFromFile(v)
	}
	fmt.Println(mRM.Items)
}
