package tmlog

import (
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
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

	hook := lumberjack.Logger{
		Filename:   path,
		MaxSize:    10,
		MaxBackups: 10,
	}
	l.SetOutput(&hook)

	lgMap[path] = l
	return l
}
