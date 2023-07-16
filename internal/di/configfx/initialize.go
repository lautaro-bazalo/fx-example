package configfx

import (
	"fxdemo/config"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(Initialize)

func Initialize() *config.Application {

	viper.SetConfigFile("config/config.json")

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
