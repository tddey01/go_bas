package config

import (
	"encoding/json"
	"os"
)

// AppConfig  服务端配置
type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
}

var ServConfig AppConfig

// InitConfig 服务端初始化
func InitConfig() *AppConfig {
	file, err := os.Open("../config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
