package tmlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var gloge *logrus.Entry

const modName = "[tmlog]"

var glog *logrus.Logger

func init() {
	glog = logrus.New()
	glog.Formatter = &logrus.TextFormatter{TimestampFormat: time.StampMilli, FullTimestamp: true}
	gloge = glog.WithField("mod", modName)
}

func GetLge() *logrus.Entry {
	return gloge
}

var lgMap = make(map[string]*logrus.Logger)
var lgMapLock sync.RWMutex

func GetLgWithPath(path string) *logrus.Logger {
	lgMapLock.Lock()
	lgMapLock.Unlock()
	if g, ok := lgMap[path]; ok {
		return g
	}

	l := logrus.New()
	l.Formatter = &logrus.TextFormatter{TimestampFormat: time.StampMilli, FullTimestamp: true}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(fmt.Sprintln("tmlog error opening file: %v", err))
	}
	l.SetOutput(f)

	lgMap[path] = l

	return l
}
