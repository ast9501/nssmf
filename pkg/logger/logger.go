package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var (
	log        *logrus.Logger
	InitLog    *logrus.Entry
	AppLog     *logrus.Entry
	HandlerLog *logrus.Entry
)

func init() {
	log = logrus.New()

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339Nano,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	InitLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "Init"})
	AppLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "App"})
	HandlerLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "Hdlr"})

	InitLog.Info("Logger init")
	AppLog.Info("Get loglevel setting from env var (LogLevel = WARN|DEBUG|INFO)")
	log.SetLevel(getLoglevel())
	InitLog.Infoln("Logger init success")
}

func getLoglevel() (level logrus.Level) {
	lv := os.Getenv("LogLevel")
	if lv == "WARN" {
		InitLog.Infoln("Set loglevel to Warn")
		level = logrus.WarnLevel
	} else if lv == "DEBUG" {
		InitLog.Infoln("Set loglevel to Debug")
		level = logrus.DebugLevel
	} else {
		InitLog.Infoln("Set loglevel to Info (default)")
		level = logrus.InfoLevel
	}
	return
}
