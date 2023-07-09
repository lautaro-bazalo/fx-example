package dbfx

import (
	"fmt"
	"fxdemo/config"
	"fxdemo/internal/pkg/log"
	"fxdemo/internal/pkg/model/user"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Module = fx.Provide(provideGormDB)

func provideGormDB(log log.Logger, config *config.Application) (*gorm.DB, error) {

	loggerConfig := logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      NewLogLevel(config.Logger),
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:         logger.New(log, loggerConfig),
		TranslateError: true,
	}

	dns := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True", config.Storage.User, config.Storage.Password, config.Storage.Host, config.Storage.Database)
	fmt.Println("Connecting to dsn:", dns)

	db, err := gorm.Open(mysql.Open(dns), &gormConfig)

	if err != nil {
		return nil, fmt.Errorf("falied to connect database, err:%v", err)
	}

	if err := db.AutoMigrate(&user.User{}); err != nil {
		return nil, err
	}

	return db, nil

}
func NewLogLevel(loggerConfig config.Logger) (logLevel logger.LogLevel) {
	level, err := logrus.ParseLevel(loggerConfig.LogLevel)
	if err != nil {
		return logger.Info
	}
	switch level {
	case logrus.ErrorLevel:
		logLevel = logger.Error
	case logrus.WarnLevel:
		logLevel = logger.Warn
	default:
		logLevel = logger.Info
	}
	return
}
