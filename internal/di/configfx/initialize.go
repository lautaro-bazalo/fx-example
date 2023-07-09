package configfx

import (
	"fxdemo/config"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(Initialize)

func Initialize() *config.Application {

	viper.SetConfigFile("config/config.json")
	/*
		_ = viper.BindEnv("Storage.Host", "TORWART_DATABASE_HOST")
		_ = viper.BindEnv("Storage.User", "TORWART_DATABASE_USER")
		_ = viper.BindEnv("Storage.Password", "TORWART_DATABASE_PASSWORD")
		_ = viper.BindEnv("Storage.Database", "TORWART_DATABASE_NAME")

	*/
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var conf config.Application

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	conf.Environment = config.Local

	return &conf
}
