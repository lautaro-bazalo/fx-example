package config

import (
	"github.com/spf13/viper"
)

func ParseConfig() *Application {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")
	_ = viper.BindEnv("Storage.Host", "TORWART_DATABASE_HOST")
	_ = viper.BindEnv("Storage.User", "TORWART_DATABASE_USER")
	_ = viper.BindEnv("Storage.Password", "TORWART_DATABASE_PASSWORD")
	_ = viper.BindEnv("Storage.Database", "TORWART_DATABASE_NAME")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var conf Application

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	conf.Environment = Local

	return &conf
}
