package utils

import "github.com/jinzhu/configor"

type Config struct {
	Bot struct {
		Token   string `required:"true" json:"token"`
		ApiHost string `default:"https://api.telegram.org/bot%s/%s" json:"apiHost"`
		Debug   bool   `default:"false" json:"debug"`
	} `json:"bot"`
	Database struct {
		Type string `default:"sqlite" json:"Type"`
		Host string `default:"localhost" json:"host"`
		Port int    `default:"3306" json:"port"`
	} `json:"database"`
	Proxy struct {
		Enable bool   `json:"enable"`
		Host   string `json:"host"`
		Port   int    `json:"port"`
	} `json:"proxy"`
}

var config Config

func init() {
	configor.Load(&config, "config.json")
}

func GetConfig() Config {
	return config
}
