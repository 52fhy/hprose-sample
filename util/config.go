package util

import "github.com/go-ini/ini"

type MysqlCfg struct{
	Host string
	Port int32
	User string
	Password string
	Database string
}

type RedisCfg struct{
	Host string
	Port int32
	Auth string
}

type Config struct {
	ListenAddr string
	Mysql MysqlCfg
	Redis RedisCfg
}

//全局变量
var Cfg = &Config{}

//加载配置
func InitConfig(ConfigFile string) error {
	return ini.MapTo(Cfg, ConfigFile)
}
