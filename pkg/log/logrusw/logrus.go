package logrusw

import (
	"github.com/sirupsen/logrus"
	"github.com/wisnuragaprawida/project/pkg/log"
)

// Logger is a wrapper for logrus.Logger
type Logger struct {
	*logrus.Logger
}

func (l *Logger) Log(level log.Level, arg ...interface{}) {
	l.Logger.Log(logrus.Level(level), arg...)
}
