package config

import (
	"github.com/spf13/viper"
	"github.com/tristin2024/logs"
)

var Viper *viper.Viper

func init() {
	Viper = viper.GetViper()
	Viper.SetConfigFile("./config/config.yaml") //指定配置文件
	// Viper.AddConfigPath("./config") //指定查找配置文件的路径
	if err := Viper.ReadInConfig(); err != nil {
		logs.Std.Errorf("fatal error config file: %s", err)
		panic(err)
	}
}
