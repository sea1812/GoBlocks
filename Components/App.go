package Components

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/net/ghttp"
)

/*
	微服务后端主对象TApp
*/

type TApp struct {
	Name               string        //应用名称
	Version            string        //应用版本号的文字表示
	VersionCode        float32       //版本号的整数表示
	Domain             string        //绑定的域名
	RootUrl            string        //根路径字符串，如/api之类或/表示根路径
	DefaultProtocol    ProtocolType  //默认访问协议，http或https二选一，为空则默认http
	DefaultPort        int           //HTTP对应的默认端口，如为0则监听80/443（视上一项选择http或https）
	TLSPort            int           //HTTPS对应的端口
	CertFile           string        //HTTPS用的cert文件路径
	KeyFile            string        //HTTPS用的key文件路径
	EnableLua          bool          //是否启用Lua脚本
	EnablePlugin       bool          //是否启用插件
	EnableProxy        bool          //是否启用反向代理
	LuaDir             string        //Lua脚本保存根目录，默认为当前目录下的lua子目录
	PluginDir          string        //插件文件保存根目录，默认为当前目录下的plugin子目录
	ProxyDir           string        //反向代理配置文件保存根目录，默认为当前目录下的proxy子目录
	StaticDir          string        //静态资源文件保存根目录，默认为当前目录下的static子目录
	ConfDir            string        //设置文件保存根目录，默认为当前目录下的conf子目录
	RouteManager       TRouteManager //动态路由管理器
	EnableShareSession bool          //是否启用共享Session，即将Session保存到Redis
	EnableRedis        bool          //是否启用Redis
	EnableCache        bool          // 是否启用Cache，如果启用CAHCE同时也启用Redis则缓存后端使用Redis，否则使用内存Cache
	MainServer         *ghttp.Server //主Web服务器，对应g.Server()
	RedisServers       TRedisManager //主Redis，对应g.Redis()
	DatabaseServers    []gdb.DB      //主数据库，对应g.DB()
}

// Init 初始化，
func (p *TApp) Init() {

}
