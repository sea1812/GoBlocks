package Components

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"strings"
)

/*
	数据库连接管理对象
*/

type TDbItem struct {
	Host   string
	Port   int
	User   string
	Pass   string
	Name   string
	Type   string
	Role   string
	Weight int
}

type TDbItemGroup struct {
	Nodes []*TDbItem
	Name  string
}

type TDbManager struct {
	Groups  []*TDbItemGroup
	ConfDir string //保存配置文件的目录
}

// SaveGroupToFile 保存Group到Conf文件
func (p *TDbItemGroup) SaveGroupToFile() {
	a := gjson.New(p.Nodes)
	//组合文件名
	mFilename := gfile.Join(gfile.Join(Default_Conf_Dir, Db_Conf_Subdir), p.Name+".conf")
	mContent := a.Export()
	_ = gfile.PutContents(mFilename, mContent)
}

// LoadItemsFromFile 从Conf文件中加载Item，文件名是Group.name+".conf"
func (p *TDbItemGroup) LoadItemsFromFile() error {
	if strings.TrimSpace(p.Name) != "" {
		mFilename := gfile.Join(gfile.Join(Default_Conf_Dir, Db_Conf_Subdir), p.Name+".conf")
		if gfile.Exists(mFilename) == true {
			p.Nodes = nil
			mC := gfile.GetContents(mFilename)
			mJson := gjson.New(mC)
			for _, v := range mJson.Array() {
				mv := gjson.New(v)
				p.AddItem(TDbItem{
					Host:   mv.GetString("Host"),
					Port:   mv.GetInt("Port"),
					User:   mv.GetString("User"),
					Pass:   mv.GetString("Pass"),
					Name:   mv.GetString("Name"),
					Type:   mv.GetString("Type"),
					Role:   mv.GetString("Role"),
					Weight: mv.GetInt("Weight"),
				})
			}
			return nil
		} else {
			return errors.New("设置文件" + mFilename + "不存在")
		}
	} else {
		return errors.New("Group名为空")
	}
}

// AddItem 在Group中添加节点
func (p *TDbItemGroup) AddItem(AItem TDbItem) {
	mItem := TDbItem{
		Host:   AItem.Host,
		Port:   AItem.Port,
		User:   AItem.User,
		Pass:   AItem.Pass,
		Name:   AItem.Name,
		Type:   AItem.Type,
		Role:   AItem.Role,
		Weight: AItem.Weight,
	}
	p.Nodes = append(p.Nodes, &mItem)
}

// AddGroup 在TDbManager中增加Group
func (p *TDbManager) AddGroup(AName string) {
	mGroup := TDbItemGroup{
		Nodes: nil,
		Name:  AName,
	}
	p.Groups = append(p.Groups, &mGroup)
}

// Group 按名字查找Group
func (p *TDbManager) Group(AName string) *TDbItemGroup {
	var mR *TDbItemGroup = nil
	for _, v := range p.Groups {
		if v.Name == AName {
			mR = v
		}
	}
	return mR
}

// ToConfig 转换为gdb.Config对象
func (p *TDbManager) ToConfig() gdb.Config {
	var mConfig map[string]gdb.ConfigGroup = make(map[string]gdb.ConfigGroup)
	fmt.Println("---------1111---------")
	fmt.Println(p.Groups)
	fmt.Println("---------1111---------")
	for _, v_group := range p.Groups {
		fmt.Println("Loop group", v_group)
		var mC []gdb.ConfigNode
		for _, v_node := range v_group.Nodes {
			fmt.Println("Loop v_node", v_node)
			mC = append(mC, gdb.ConfigNode{
				Host:   v_node.Host,
				Port:   fmt.Sprint(v_node.Port),
				User:   v_node.User,
				Pass:   v_node.Pass,
				Name:   v_node.Name,
				Type:   v_node.Type,
				Role:   v_node.Role,
				Weight: v_node.Weight,
			})
			fmt.Println("Loop", mC)
		}
		mConfig[v_group.Name] = mC
	}
	return mConfig
}
