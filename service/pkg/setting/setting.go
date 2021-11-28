package setting

import (
	"github.com/go-ini/ini"
	"go.uber.org/zap"
	"log"
)

// Server 项目服务参数
type Server struct {
	RunMode      string // 启动模式
	HttpPort     string // 服务端口
	ReadTimeout  int    // 读取时间
	WriteTimeout int    // 写入时间
}

// ServerOptions 项目服务配置对象
var ServerOptions = &Server{}

// DataBase 数据库配置项参数
type DataBase struct {
	Type     string // 数据库类型
	User     string // 用户名
	Password string // 密码
	Host     string // 数据库连接
	Db       string // 数据库名
}

// DataBaseOptions 数据库配置对象
var DataBaseOptions = &DataBase{}

// Redis redis配置项
type Redis struct {
	PORT     string
	HOST     string
	UserName string
	PassWord string
}

// RedisOptions redis配置对象
var RedisOptions = &Redis{}

// go-ini 实例
var cfg *ini.File

// Setup 初始化配置项
func Setup() {
	// 预先定义错误
	var err error
	// 读取项目配置
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		zap.S().Error("setting.Setup, fail to parse 'config/app.ini': %v", err)
	}
	// 服务配置绑定结构体
	mapTo("server", ServerOptions)
	// 数据库配置绑定结构体
	mapTo("database", DataBaseOptions)
	// redis配置绑定结构体
	mapTo("redis", RedisOptions)
	zap.S().Info("setting.Setup, success to parse 'config/app.ini'")
}

// mapTo 映射参数至结构体
// sction 配置文件区域配置 [server]
func mapTo(section string, v interface{}) {
	// 锁定配置文件的区域 如配置区为 [server]
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
