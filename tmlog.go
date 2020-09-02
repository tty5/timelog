package tmlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log *logrus.Entry

const modName = "[tmlog]"

func init() {
	xlog := logrus.New()
	xlog.Formatter = &logrus.TextFormatter{TimestampFormat: time.StampMilli, FullTimestamp: true}
	log = xlog.WithField("mod", modName)

	f, err := os.OpenFile("/var/tmlog", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
		return
    }
    xlog.SetOutput(f)
}

func GetLogger() *logrus.Entry {
	return log
}
