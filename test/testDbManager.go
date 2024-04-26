package test

import (
	cc "GoBlocks/Components"
	"fmt"
)

func TestDM() {
	mDM := cc.TDbManager{
		ConfDir: cc.Db_Conf_Subdir,
	}
	fmt.Println("添加Group")
	mDM.AddGroup("default")
	fmt.Println("添加第二个Group")
	mDM.AddGroup("second")
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
	mGroup.SaveGroupToFile(cc.Db_Conf_Subdir)
	fmt.Println("从文件载入")
	mGroup.Nodes = nil
	_ = mGroup.LoadItemsFromFile()
	fmt.Println(mGroup.Nodes)
}
