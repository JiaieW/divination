package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	Qianwen struct {
		ApiUrl string
		ApiKey string
	} `json:"qianwen"`
	Xinghuo struct {
		HostUrl   string `json:"hostUrl"`
		Appid     string `json:"appid"`
		ApiSecret string `json:"apiSecret"`
		ApiKey    string `json:"apiKey"`
	} `json:"xinghuo"`
}

var (
	once sync.Once
	C    *Config
)

func init() {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath("./config")
		if err := viper.ReadInConfig(); err != nil {
			panic("read in config err: " + err.Error())
		}

		C = &Config{}
		if err := viper.Unmarshal(C); err != nil {
			panic("unmarshal config err: " + err.Error())
		}
	})
}
