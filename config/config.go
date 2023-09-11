package config

import (
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	Port string `yaml:"port"`
}

// 配置结构体
type Config struct {
	Server Server
}

var (
	Cfg *Config // 配置
)

// InitConfig 初始化配置等
func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		logrus.Panic(err)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logrus.Fatal("config init error.", err)
		} else {
			// Config file was found but another error was produced
			logrus.Info("config init.")
		}
	}
	// 序列化到结构体
	viper.Unmarshal(&Cfg)

	logrus.WithField("data", Cfg).Info("Init config success!")
	// 配置文件修改事件
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Info("Config file changed:", e.Name)
	})
	// viper 监听配置文件
	viper.WatchConfig()
}
