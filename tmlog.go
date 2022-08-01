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
	glog.Formatter = &logrus.JSONFormatter{TimestampFormat: time.StampMicro}
}

func GetLg() *logrus.Logger {
	return glog
}

func GetLgWithPathLogSize(path string, maxSize int, maxBackups int) *logrus.Logger {
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{TimestampFormat: time.StampMicro}

	hook := lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		LocalTime:  true,
	}
	l.SetOutput(&hook)

	return l
}

func GetLgWithPath(path string) *logrus.Logger {
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{TimestampFormat: time.StampMicro}

	hook := lumberjack.Logger{
		Filename:   path,
		MaxSize:    50,
		MaxBackups: 5,
		LocalTime:  true,
	}
	l.SetOutput(&hook)

	return l
}

var lgMap = make(map[string]*logrus.Logger)
var lgMapLock sync.RWMutex

func GetLgWithPathUni(path string) *logrus.Logger {
	lgMapLock.RLock()
	if g, ok := lgMap[path]; ok {
		lgMapLock.RUnlock()
		return g
	}
	lgMapLock.RUnlock()

	lgMapLock.Lock()
	defer lgMapLock.Unlock()
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{TimestampFormat: time.StampMicro}

	hook := lumberjack.Logger{
		Filename:   path,
		MaxSize:    50,
		MaxBackups: 5,
		LocalTime:  true,
	}
	l.SetOutput(&hook)

	lgMap[path] = l
	return l
}
