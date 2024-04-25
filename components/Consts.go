package components

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
