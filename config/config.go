package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	Base struct {
		Debug bool `mapstructure:"debug"`
		Port  int  `mapstructure:"port"`
	} `mapstructure:"base"`

	DB struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
		User string `mapstructure:"user"`
		Pass string `mapstructure:"password"`
		Name string `mapstructure:"dbname"`
	} `mapstructure:"database"`

	Redis struct {
		Addr string `mapstructure:"addr"`
		Pass string `mapstructure:"password"`
		DB   int    `mapstructure:"db"`
	} `mapstructure:"redis"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() *Config {
	return config
}
