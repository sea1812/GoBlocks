package Components

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"strings"
)

/*
	微服务路由管理器对象 TRouteManager
	脚本和插件使用固定格式的路由，样式如下：
	脚本路由：http://domain.com/lua/aaa		aaa为脚本标识名
	插件路由：http://domain.com/plugin/aaa	aaa为插件标识名
*/

type TRouteItem struct {
	RouteType   TRouteType //路由项类型
	UrlString   string     //路由路径，对脚本和插件此处填充固定的路由，对代理则填充完整的路由
	RouteTarget string     //路由目标，根据RouteType不同，处理方式也不同
}

type TRouteManager struct {
	Items           []TRouteItem //路由项列表
	ScriptUrlPrefix string       //对脚本的路由前置路径
	PluginUrlPrefix string       //对插件的路由前置路径
	ConfDir         string       //保存conf文件的文件夹
}

// Init 初始化
func (p *TRouteManager) Init() {
	p.Items = nil
	p.ScriptUrlPrefix = "/lua"
	p.PluginUrlPrefix = "/plugin"
	p.ConfDir = "./conf"
	if gfile.Exists(p.ConfDir) == false {
		_ = gfile.Mkdir(p.ConfDir)
	}
}

// AddItem 添加新项目
func (p *TRouteManager) AddItem(AType TRouteType, AUrl string, ATarget string) {
	mItem := TRouteItem{
		RouteType:   AType,
		UrlString:   AUrl,
		RouteTarget: ATarget,
	}
	p.Items = append(p.Items, mItem)
}

// AddItemFromFile 从文件中加载并添加Item，文件格式为JSON文本
func (p *TRouteManager) AddItemFromFile(AFilename string) error {
	var mR error
	var mFilename string
	if strings.Contains(AFilename, ".conf") != true {
		mFilename = AFilename + ".conf"
	} else {
		mFilename = AFilename
	}
	if gfile.Exists(AFilename) == true {
		mC := gfile.GetContents(mFilename)
		mJson := gjson.New(mC)
		mRouteType := mJson.GetInt("RouteType")
		mUrl := mJson.GetString("UrlString")
		mTarget := mJson.GetString("RouteTarget")
		p.AddItem(TRouteType(mRouteType), mUrl, mTarget)
		mR = nil
	} else {
		mR = errors.New(AFilename + "文件不存在")
	}
	return mR
}

// SaveItemsToFile 将Items分别保存到不同的文件中
func (p *TRouteManager) SaveItemsToFile(ASubDir string) {
	var mFilename string
	var mContent string
	for _, v := range p.Items {
		//用下划线替换URL中的/分隔符作为文件名
		mFilename = strings.Replace(v.UrlString, "/", "_", -1)
		mFilename = gfile.Join(p.ConfDir, gfile.Join(ASubDir, mFilename)+".conf")
		fmt.Println("文件名：", mFilename)
		mJson := gjson.New(v)
		mContent = mJson.Export()
		_ = gfile.PutContents(mFilename, mContent)
	}
}

// AllConfFiles 列举指定目录下所有conf文件
func (p *TRouteManager) AllConfFiles() []string {
	mDir := p.ConfDir
	mDir = gfile.Join(p.ConfDir, "Routes")
	mFiles, _ := gfile.ScanDirFile(mDir, "*.conf", false)
	return mFiles
}
