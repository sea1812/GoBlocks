package test

import (
	cc "GoBlocks/Components"
	"fmt"
)

// 测试TRedisManager
func TestRedisM() {
	mRM := cc.TRedisManager{
		Items:   nil,
		ConfDir: cc.Default_Conf_Dir,
	}
	fmt.Println("增加服务器-1")
	mRM.AddItem(cc.TRedisItem{
		Name:            "Demo1",
		Host:            "127.0.0.1",
		Port:            6379,
		Db:              0,
		Pass:            "",
		MaxIdle:         1000,
		MaxActive:       1000,
		IdleTimeout:     1000,
		MaxConnLifetime: 1000,
		ConnectTimeout:  30000,
		TLS:             false,
		TLSSkipVerify:   false,
		Instance:        nil,
	})
	fmt.Println("增加服务器-2")
	mRM.AddItem(cc.TRedisItem{
		Name:            "Demo2",
		Host:            "127.0.0.1",
		Port:            6379,
		Db:              1,
		Pass:            "",
		MaxIdle:         1000,
		MaxActive:       1000,
		IdleTimeout:     1000,
		MaxConnLifetime: 1000,
		ConnectTimeout:  30000,
		TLS:             false,
		TLSSkipVerify:   false,
		Instance:        nil,
	})
	fmt.Println("--------", len(mRM.Items))
	fmt.Println("保存到文件")
	mRM.SaveItemsToFile(cc.Redis_Conf_Subdir)
	fmt.Println("Conf文件列表")
	mFiles := mRM.AllConfFiles()
	fmt.Println(mFiles)
	fmt.Println("创建实例")
	mRM.Start()
}
