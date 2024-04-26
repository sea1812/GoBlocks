package test

import (
	cc "GoBlocks/Components"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	_ "github.com/lib/pq"
)

func TestDM() {
	mDM := cc.TDbManager{
		ConfDir: cc.Db_Conf_Subdir,
	}
	fmt.Println("添加Group")
	mDM.AddGroup("default")
	//找到第一个Group
	mGroup := mDM.Group("default")
	if mGroup != nil {
		fmt.Println("在group中增加Item")
		mGroup.AddItem(cc.TDbItem{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "postgres",
			Pass:     "123456",
			Name:     "test",
			Type:     string(cc.Db_Postgres),
			Role:     "master",
			Weight:   0,
			Instance: nil,
		})
	}
	fmt.Println("保存到文件")
	mGroup.SaveGroupToFile()
	fmt.Println("从文件载入")
	mGroup.Nodes = nil
	_ = mGroup.LoadItemsFromFile()
	fmt.Println("当前的Groups")
	fmt.Println(mGroup.Nodes)
	fmt.Println("---------")
	fmt.Println("当前的DM")
	fmt.Println(mDM)
	fmt.Println("查询下数据库")
	mConfig := mDM.ToConfig()
	gdb.SetConfig(mConfig)
	mDb2, _ := gdb.New("default")
	res, _ := mDb2.Model("test").All()
	fmt.Println("查询结果：", res)
}
