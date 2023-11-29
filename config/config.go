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
