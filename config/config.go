package config

import (
	"time"

	"github.com/spf13/viper"
)

type LoggerConfig struct {
	Dir    string `mapstructure:"dir"`
	Access string `mapstructure:"access"`
	Error  string `mapstructure:"error"`
	Debug  string `mapstructure:"debug"`
}

type DBConfig struct {
	Debug    bool   `mapstructure:"debug"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	Password     string `mapstructure:"password"`
	PoolSize     int    `mapstructure:"poolsize"`
	MinIdleConns int    `mapstructure:"minidleconns"`
}

type ContactConfig struct {
	Email string `mapstructure:"email"`
}

type SessionConfig struct {
	SecretKey string        `mapstructure:"secretKey"` // session secret key
	Expire    time.Duration `mapstructure:"expire"`
}

// CrawlerExecutor 爬虫服务器信息
type CrawlerExecutor struct {
	Token string `mapstructure:"token"`
	Path  string `mapstructure:"path"`
}

type config struct {
	Domain          string          `mapstructure:"domain"`  // 站点
	Port            uint            `mapstructure:"port"`    // 服务端口
	Debug           bool            `mapstructure:"debug"`   // debug 开关
	Redis           RedisConfig     `mapstructure:"redis"`   // redis配置
	Logger          LoggerConfig    `mapstructure:"logger"`  // logger配置
	DB              DBConfig        `mapstructure:"db"`      // db配置
	Contact         ContactConfig   `mapstructure:"contact"` // 联系方式
	Session         SessionConfig   `mapstructure:"session"`
	CrawlerExecutor CrawlerExecutor `mapstructure:"crawler_executor"` // 爬虫相关信息
}

var Global config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Global)
	if err != nil {
		panic(err)
	}
}
