package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName  string
	LogLevel string

	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}

func main() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	var c Config
	v.Unmarshal(&c)
	fmt.Println(v.Get("app_name"))
	fmt.Println(v.Get("log_level"))

	fmt.Println("mysql ip: ", v.Get("mysql.ip"))
	fmt.Println("mysql port: ", v.Get("mysql.port"))
	fmt.Println("mysql user: ", v.Get("mysql.user"))
	fmt.Println("mysql password: ", v.Get("mysql.password"))
	fmt.Println("mysql database: ", v.Get("mysql.database"))
	fmt.Println(c)
}
