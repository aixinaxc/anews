package base

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
)

var RedisUrl string
var RedisPassword string
var MysqlUrl string
var AppId string
var AppSecret string


func Config()  {
	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println("config:",err)
		return
	}
	RedisUrl,err = config.Get("RedisUrl")
	RedisPassword,err = config.Get("RedisPassword")
	MysqlUrl,err = config.Get("MysqlUrl")
	AppId,err = config.Get("AppId")
	AppSecret,err = config.Get("AppSecret")
	if err != nil{
		fmt.Println("config-data:",err)
	}
	fmt.Println("redis_config:",RedisUrl,RedisPassword)
	fmt.Println("mysql_config:",MysqlUrl)
	fmt.Println("redis_config:",RedisUrl,RedisPassword)
	fmt.Println("weixin:",AppId,AppSecret)
}
