package main

import (
	"github.com/henrylee2cn/pholcus/exec"
	_ "hyperspider/rules"
	"github.com/henrylee2cn/pholcus/config"
	"github.com/henrylee2cn/pholcus/runtime/cache"
)

func main() {
	// 设置运行时默认操作界面，并开始运行
	// 运行软件前，可设置 -a_ui 参数为"web"、"gui"或"cmd"，指定本次运行的操作界面
	// 其中"gui"仅支持Windows系统
	exec.DefaultRun("web")
}

// 自定义相关配置，将覆盖默认值
func init() {
	// 标记当前init()已执行完毕
	defer cache.ExecInit(0)

	// 允许日志打印行号
	// logs.ShowLineNum()

	//mongodb链接字符串
	config.MGO_CONN_STR = "127.0.0.1:27017"
	//mongodb数据库
	config.DB_NAME = "pholcus"
	//mongodb连接池容量
	config.MGO_CONN_CAP = 1024

	//mysql服务器地址
	config.MYSQL_CONN_STR = "root:root@tcp(127.0.0.1:3306)"
	//mysql连接池容量
	config.MYSQL_CONN_CAP = 1024

	//
	//// 历史记录文件名前缀
	//config.HISTORY.FILE_NAME_PREFIX = "history"
	//
	//// 代理IP完整文件名
	//config.PROXY_FULL_FILE_NAME = "proxy.pholcus"
	//
	//// Surfer-Phantom下载器配置
	//config.SURFER_PHANTOM.FULL_APP_NAME = "phantomjs" //phantomjs软件相对路径与名称
}