package tmlog

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var glog *logrus.Logger

func init() {
	glog = logrus.New()
	glog.Formatter = &logrus.JSONFormatter{TimestampFormat: time.StampMilli}
}

func GetLg() *logrus.Logger {
	return glog
}

var lgMap = make(map[string]*logrus.Logger)
var lgMapLock sync.RWMutex

func GetLgWithPath(path string) *logrus.Logger {
	lgMapLock.RLock()
	if g, ok := lgMap[path]; ok {
		lgMapLock.RUnlock()
		return g
	}
	lgMapLock.RUnlock()

	lgMapLock.Lock()
	defer lgMapLock.Unlock()
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{TimestampFormat: time.StampMilli}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		glog.Errorln("tmlog use glog due to error opening file:", path, err)
		l = glog
	} else {
		l.SetOutput(f)
	}

	lgMap[path] = l
	return l
}
