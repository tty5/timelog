package tmlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log *logrus.Entry

const modName = "[tmlog]"

var xlog *logrus.Logger

func init() {
	xlog = logrus.New()
	xlog.Formatter = &logrus.TextFormatter{TimestampFormat: time.StampMilli, FullTimestamp: true}
	log = xlog.WithField("mod", modName)
}

func GetLogger() *logrus.Entry {
	return log
}

func SetLogPath(path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(fmt.Sprintln("tmlog error opening file: %v", err))
	}
	xlog.SetOutput(f)
}
