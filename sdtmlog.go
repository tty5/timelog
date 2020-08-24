package sdtmlog

import (
	"github.com/sirupsen/logrus"
	"time"
)

var log *logrus.Entry

const pkgname = "[sdtm]"

func init() {
	xlog := logrus.New()
	xlog.Formatter = &logrus.TextFormatter{TimestampFormat: time.StampMilli}
	log = xlog.WithField("pkg", pkgname)
}

func GetLogger() *logrus.Entry {
	return log
}
