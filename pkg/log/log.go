package log

import "fmt"

type Logger interface {
	Log(Level, ...interface{})
}
type Level uint32

const (
	//PanicsLevel level, highest level of severity. Logs and then calls panic with the
	//message passed to Debug, Info, ...
	PanicsLevel Level = iota

	//FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	//logging level is set to Panic.
	FatalLevel
	//ErorrLevel level. Logs. Used for errors that should definitely be noted.
	//Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	//WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	//InfoLevel level. General operational entries about what's going on inside the
	//application.
	InfoLevel
	//DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	//TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

type nilLogger struct {
}

func (nilLogger) Log(Level, ...interface{}) {
	//do nothing
}

var defaultlogger Logger = nilLogger{}

func SetLogger(l Logger) {
	if l == nil {
		defaultlogger = nilLogger{}

	}
	defaultlogger = l
}

func Log(l Level, args ...interface{}) {
	defaultlogger.Log(l, args...)
}

func Debug(args ...interface{}) {
	defaultlogger.Log(DebugLevel, args...)
}

func Debugf(format string, args ...interface{}) {
	defaultlogger.Log(DebugLevel, fmt.Sprintf(format, args...))
}

func Info(args ...interface{}) {
	defaultlogger.Log(InfoLevel, args...)
}

func Infof(format string, args ...interface{}) {
	defaultlogger.Log(InfoLevel, fmt.Sprintf(format, args...))
}

func Warn(args ...interface{}) {
	defaultlogger.Log(WarnLevel, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultlogger.Log(WarnLevel, fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	defaultlogger.Log(ErrorLevel, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultlogger.Log(ErrorLevel, fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	defaultlogger.Log(FatalLevel, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultlogger.Log(FatalLevel, fmt.Sprintf(format, args...))
}

func Panic(args ...interface{}) {
	defaultlogger.Log(PanicsLevel, args...)
}

func Panicf(format string, args ...interface{}) {
	defaultlogger.Log(PanicsLevel, fmt.Sprintf(format, args...))
}

func Trace(args ...interface{}) {
	defaultlogger.Log(TraceLevel, args...)
}

func Tracef(format string, args ...interface{}) {
	defaultlogger.Log(TraceLevel, fmt.Sprintf(format, args...))
}
