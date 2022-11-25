package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	AppName          string `mapstructure:"app_name"`
	Server           `mapstructure:"server"`
	DataSourceConfig `mapstructure:"database"`
	RedisConfig      `mapstructure:"redis"`
}

type DataSourceConfig struct {
	DriverName     string `mapstructure:"driver_name"`
	Host           string `mapstructure:"host"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	DB             string `mapstructure:"db_name"`
	LogMode        bool   `mapstructure:"log-mode"`
	Port           int    `mapstructure:"port"`
	MaxOpenConnect int    `mapstructure:"max_open_connect"`
	MaxIdleConnect int    `mapstructure:"max_idle_connect"`
	Charset        string `mapstructure:"charset"`
	Collation      string `mapstructure:"collation"`
	Query          string `mapstructure:"query"`
}

type RedisConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `yaml:"port"`
	Password    string `mapstructure:"password"`
	DB          int    `mapstructure:"db"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxActive   int    `mapstructure:"max_active"`
	IdleTimeout int    `mapstructure:"idle_timeout"`
	PoolSize    int    `mapstructure:"pool_size"`
}

type Server struct {
	HttpPort     string `mapstructure:"http-port"`
	RunMode      string `mapstructure:"run-mode"`
	ReadTimeout  int    `mapstructure:"read-timeout"`
	WriteTimeout int    `mapstructure:"write-timeout"`
}

var Conf *AppConfig = nil

func InitConfigFile() {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("config")
	v.AddConfigPath(".")

	// 读取配置信息
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("viper读取配置信息失败：%v\n", err)
		panic(err)
		return
	}

	//将配置文件反序列化到_config
	BindConfig(v)

	//监控配置
	v.WatchConfig()

	//配置文件实时刷新
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Println("-------------配置文件修改了-------------")
		BindConfig(v)
	})
}

func BindConfig(v *viper.Viper) {
	//绑定配置文件中的所有配置项
	if err := v.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("绑定配置文件失败：%s \n", err))
	} else {
		log.Println("绑定配置文件成功！", *Conf)
	}
}
