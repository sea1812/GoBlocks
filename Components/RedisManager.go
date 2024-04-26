package Components

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"time"
)

/*
	Redis服务器机群管理器对象
*/

type TRedisItem struct {
	Name            string        //服务器名
	Host            string        `json:"Host"`
	Port            int           `json:"Port"`
	Db              int           `json:"Db"`
	Pass            string        `json:"Pass"`            // Password for AUTH.
	MaxIdle         int           `json:"MaxIdle"`         // Maximum number of connections allowed to be idle (default is 10)
	MaxActive       int           `json:"MaxActive"`       // Maximum number of connections limit (default is 0 means no limit).
	IdleTimeout     time.Duration `json:"IdleTimeout"`     // Maximum idle time for connection (default is 10 seconds, not allowed to be set to 0)
	MaxConnLifetime time.Duration `json:"MaxConnLifetime"` // Maximum lifetime of the connection (default is 30 seconds, not allowed to be set to 0)
	ConnectTimeout  time.Duration `json:"ConnectTimeout"`  // Dial connection timeout.
	TLS             bool          `json:"Tls"`             // Specifies the config to use when a TLS connection is dialed.
	TLSSkipVerify   bool          `json:"TlsSkipVerify"`   // Disables server name verification when connecting over TLS.
	Instance        *gredis.Redis //服务器实例
}

type TRedisManager struct {
	Items   []*TRedisItem //机群列表
	ConfDir string        //保存conf文件的文件夹
}

// AddItem 添加服务器记录
func (p *TRedisManager) AddItem(AItem TRedisItem) {
	mItem := TRedisItem{
		Name:           AItem.Name,
		Host:           AItem.Host,
		Port:           AItem.Port,
		Db:             AItem.Db,
		Pass:           AItem.Pass,
		MaxIdle:        AItem.MaxIdle,
		MaxActive:      AItem.MaxActive,
		IdleTimeout:    AItem.IdleTimeout,
		ConnectTimeout: AItem.ConnectTimeout,
		TLS:            AItem.TLS,
		TLSSkipVerify:  AItem.TLSSkipVerify,
		Instance:       nil,
	}
	p.Items = append(p.Items, &mItem)
}

// AddItemFromFile 从Conf文件中添加服务器记录
func (p *TRedisManager) AddItemFromFile(AFilename string) error {
	if gfile.Exists(AFilename) == true {
		mText := gfile.GetContents(AFilename)
		mJson := gjson.New(mText)
		p.AddItem(TRedisItem{
			Name:            mJson.GetString("Name"),
			Host:            mJson.GetString("Host"),
			Port:            mJson.GetInt("Port"),
			Db:              mJson.GetInt("Db"),
			Pass:            mJson.GetString("Pass"),
			MaxIdle:         mJson.GetInt("MaxIdle"),
			MaxActive:       mJson.GetInt("MaxActive"),
			IdleTimeout:     mJson.GetDuration("IdleTimeout"),
			MaxConnLifetime: mJson.GetDuration("MaxConnLifetime"),
			ConnectTimeout:  mJson.GetDuration("ConnectTimeout"),
			TLS:             mJson.GetBool("TLS"),
			TLSSkipVerify:   mJson.GetBool("TLSSkipVerify"),
			Instance:        nil,
		})
		return nil
	} else {
		return errors.New(AFilename + "文件不存在")
	}
}

// Start 初始化，创建各个连接的实例
func (p *TRedisManager) Start() {
	//循环创建Redis连接对象
	for _, v := range p.Items {
		v.Instance = gredis.New(&gredis.Config{
			Host:            v.Host,
			Port:            v.Port,
			Db:              v.Db,
			Pass:            v.Pass,
			MaxIdle:         v.MaxIdle,
			MaxActive:       v.MaxActive,
			IdleTimeout:     v.IdleTimeout,
			MaxConnLifetime: v.MaxConnLifetime,
			ConnectTimeout:  v.ConnectTimeout,
			TLS:             v.TLS,
			TLSSkipVerify:   v.TLSSkipVerify,
		})
	}
}

// Server 根据服务器名字查找RedisItem
func (p *TRedisManager) Server(AName string) *TRedisItem {
	var mR *TRedisItem = nil
	for _, v := range p.Items {
		if v.Name == AName {
			mR = v
			break
		}
	}
	return mR
}

// SaveItemsToFile 将Items分别保存到不同的文件中
func (p *TRedisManager) SaveItemsToFile(ASubDir string) {
	var mFilename string
	var mContent string
	for _, v := range p.Items {
		//用下划线替换URL中的/分隔符作为文件名
		mFilename = v.Name
		mFilename = gfile.Join(p.ConfDir, gfile.Join(ASubDir, mFilename)+".conf")
		fmt.Println("文件名：", mFilename)
		mJson := gjson.New(v)
		mContent = mJson.Export()
		_ = gfile.PutContents(mFilename, mContent)
	}
}

// AllConfFiles 列举指定目录下所有conf文件
func (p *TRedisManager) AllConfFiles() []string {
	mDir := p.ConfDir
	mDir = gfile.Join(p.ConfDir, Redis_Conf_Subdir)
	mFiles, _ := gfile.ScanDirFile(mDir, "*.conf", false)
	return mFiles
}
