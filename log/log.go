package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

// InitLogger 初始化logrus格式级别
func InitLogger() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})
}
