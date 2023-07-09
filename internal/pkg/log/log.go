package log

import (
	"fxdemo/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Logger interface {
	Error(args ...interface{})
	Print(args ...interface{})
	Printf(fmt string, args ...interface{})
	Errorf(format string, v ...interface{})
}

type logger struct {
	entry logrus.Entry
}

func NewLogger(config config.Logger) Logger {

	hostname, _ := os.Hostname()

	l := logrus.New()
	level := config.LogLevel
	if level != "" {
		parseLevel, e := logrus.ParseLevel(level)
		if e == nil {
			l.SetLevel(parseLevel)
		}
	}

	l.Formatter = &Formatter{host: hostname}

	if config.EnableFile {
		lumberjackLogger := &lumberjack.Logger{
			Filename: config.FileSettings.FileLocation,
			MaxSize:  config.FileSettings.MaxSize,
			MaxAge:   config.FileSettings.MaxAge,
			Compress: true,
		}
		l.SetOutput(lumberjackLogger)
	} else {
		l.SetOutput(os.Stdout)
	}

	return &logger{
		entry: *logrus.NewEntry(l),
	}
}

func (l *logger) Printf(s string, fmt ...interface{}) {
	l.entry.Printf(s, fmt...)
}

func (l *logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *logger) Print(args ...interface{}) {
	l.entry.Print(args...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.entry.Errorf(format, v...)
}
