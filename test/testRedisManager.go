package test

import (
	cc "GoBlocks/Components"
	"fmt"
	"time"
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
		MaxIdle:         300000,
		MaxActive:       300000,
		IdleTimeout:     300000,
		MaxConnLifetime: 300000,
		ConnectTimeout:  300000,
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
		MaxIdle:         300000,
		MaxActive:       300000,
		IdleTimeout:     300000,
		MaxConnLifetime: 300000,
		ConnectTimeout:  300000,
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
	fmt.Println("创建实例，连接服务器")
	mRM.Start()
	fmt.Println("按名字查找")
	mServer := mRM.Server("Demo1")
	fmt.Println("操作下Demo1")
	_, er := mServer.Instance.Do("SET", "Key1", time.Now())
	fmt.Println("操作返回的ERROR", er)
	mServer2 := mRM.Server("Demo2")
	fmt.Println("操作下Demo2")
	_, er2 := mServer2.Instance.Do("SET", "Key2", time.Now())
	fmt.Println("操作返回的ERROR", er2)
}
