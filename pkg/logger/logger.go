package logger

import (
	"os"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var (
	log       *logrus.Logger
	AppLog    *logrus.Entry
	NssLog    *logrus.Entry
	GinLog    *logrus.Entry
	DbLog     *logrus.Entry
	ConfigLog *logrus.Entry
	DmaapLog  *logrus.Entry
	NsLog     *logrus.Entry
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

	AppLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "App"})
	NssLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "NS Subnet Provider"})
	GinLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "Gin"})
	DbLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "DB"})
	DmaapLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "DMaaP Consumer"})
	ConfigLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "ConfigLoader"})
	NsLog = log.WithFields(logrus.Fields{"component": "RAN NSSMF", "category": "Network Service Provider"})

	AppLog.Info("Logger init")
	AppLog.Info("Get loglevel setting from env var (LogLevel = WARN|DEBUG|INFO)")
	log.SetLevel(getLoglevel())
	AppLog.Infoln("Logger init success")
}

func getLoglevel() (level logrus.Level) {
	lv := os.Getenv("LogLevel")
	if lv == "WARN" {
		AppLog.Infoln("Set loglevel to Warn")
		level = logrus.WarnLevel
	} else if lv == "DEBUG" {
		AppLog.Infoln("Set loglevel to Debug")
		level = logrus.DebugLevel
	} else {
		AppLog.Infoln("Set loglevel to Info (default)")
		level = logrus.InfoLevel
	}
	return
}
