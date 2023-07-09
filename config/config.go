package config

import (
	"time"
)

type Environment string

const (
	Local Environment = "local"
)

type Application struct {
	App         app
	Logger      Logger
	Storage     storage
	Environment Environment
}

type app struct {
	Port string
}

type storage struct {
	Host                  string
	Database              string
	User                  string
	Password              string
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionIdleTime time.Duration
	MaxConnectionLifetime time.Duration
}

type Logger struct {
	LogLevel     string
	PrettyLog    bool
	EnableFile   bool
	FileSettings fileSettings
}

type fileSettings struct {
	FileLocation string
	MaxSize      int
	MaxAge       int
}
