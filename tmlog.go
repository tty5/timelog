package tmlog

import (
	"github.com/sirupsen/logrus"
	"time"
)

var log *logrus.Entry

const modName = "[tmlog]"

func init() {
	xlog := logrus.New()
	xlog.Formatter = &logrus.TextFormatter{TimestampFormat: time.StampMilli}
	log = xlog.WithField("mod", modName)
}

func GetLogger() *logrus.Entry {
	return log
}
