package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

type Server struct {
	Port int
}

type Jwt struct {
	Secret string
	Expire int // 过期时间、秒为单位
}

type Conf struct {
	Server   Server
	Database Database
	Jwt      Jwt
}

var Config Conf

func Setup() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	fmt.Printf("env: %s\n", env)
	flag.Parse()
	viper.SetEnvPrefix(env)
	viper.AutomaticEnv()
	viper.SetConfigName("config." + env)
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if viper.Unmarshal(&Config) != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
