package Components

/*
	常量和数据类型定义
*/

// ProtocolType TApp中的访问协议类型
type ProtocolType int

const (
	Visit_via_http  ProtocolType = 1
	Visit_vis_https ProtocolType = 2
)

// TRouteType TRouteItem的路由类型
type TRouteType int

const (
	Route_Internal TRouteType = 1 //内部路由，基本用不到
	Route_Lua      TRouteType = 2 //指向Lua脚本的路由项
	Route_Plugin   TRouteType = 3 //指向插件的路由项
	Route_Proxy    TRouteType = 4 //指向反向代理的路由项
	Route_Static   TRouteType = 5 //指向静态目录的路由项
)

// 保存Conf文件的路径常量
const (
	Default_Conf_Dir   string = "./Conf"   //Conf子目录，所有Conf文件都分别保存到相应的子目录下
	Routes_Conf_Subdir string = "Routes"   //保存路由设置的Conf文件
	Redis_Conf_Subdir  string = "Redis"    //保存Redis设置的Conf文件
	Db_Conf_Subdir     string = "Database" //保存Redis设置的Conf文件
	System_Conf_Subdir string = "System"   //保存系统设置的Conf文件
)

// TDbType TDbManager的数据库type
type TDbType string

const (
	Db_MySQL    TDbType = "mysql"
	Db_Postgres TDbType = "pgsql"
	Db_Sqlite   TDbType = "sqlite"
	Db_MsSQL    TDbType = "mssql"
	Db_Oracle   TDbType = "oracle"
)
