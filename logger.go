package logs

import (
	"runtime"
	"sync"
	"time"

	"github.com/one-go/logs/common"
	"github.com/one-go/logs/level"
)

type Logger struct {
	mu        sync.Mutex
	name      string
	level     level.Level
	provider  common.Provider
	logStack  bool
	calldepth int
}

func NewLogger(name string, provider ...common.Provider) *Logger {
	var p common.Provider
	if len(provider) > 0 && provider[0] != nil {
		p = provider[0]
	} else {
		p = GetProvider()
	}

	return &Logger{
		name:      name,
		level:     level.ERROR,
		provider:  p,
		calldepth: 2,
	}
}

func (l *Logger) Name() string {
	return l.name
}

func (l *Logger) Provider() common.Provider {
	return l.provider
}

func (l *Logger) Level() level.Level {
	return l.level
}

func (l *Logger) IsLogStack() bool {
	return l.logStack
}

func (l *Logger) isEnabled(level level.Level) bool {
	return l.level <= level
}

func (l *Logger) IsDebugEnabled() bool {
	return l.isEnabled(level.DEBUG)
}

func (l *Logger) IsInfoEnabled() bool {
	return l.isEnabled(level.INFO)
}

func (l *Logger) IsWarnEnabled() bool {
	return l.isEnabled(level.WARN)
}

func (l *Logger) IsErrorEnabled() bool {
	return l.isEnabled(level.ERROR)
}

func (l *Logger) IsFatalEnabled() bool {
	return l.isEnabled(level.FATAL)
}

func (l *Logger) log(level level.Level, message ...interface{}) {
	if !l.isEnabled(level) {
		return
	}

	entry := &common.Entry{
		Level:   level,
		Time:    time.Now(),
		Message: message,
	}

	if pc, _, line, ok := runtime.Caller(l.calldepth); ok {
		entry.Caller = runtime.FuncForPC(pc).Name()
		entry.Line = line
	} else {
		entry.Caller = "???"
		entry.Line = 0
	}

	// if l.logStack || level >= level.ERROR {
	// 	entry.StackTrace = getFrames(l.calldepth)
	// }

	l.provider.Log(entry)
}

func (l *Logger) Debug(message ...interface{}) {
	l.log(level.DEBUG, message...)
}

func (l *Logger) Info(message ...interface{}) {
	l.log(level.INFO, message...)
}

func (l *Logger) Warn(message ...interface{}) {
	l.log(level.WARN, message...)
}

func (l *Logger) Error(message ...interface{}) {
	l.log(level.ERROR, message...)
}

func (l *Logger) Fatal(message ...interface{}) {
	l.log(level.FATAL, message...)
}

func (l *Logger) SetProvider(provider common.Provider) {
	if provider == nil {
		panic("[mano.logs.SetProvider] Invalid argument: provider")
	}

	l.mu.Lock()

	defer l.mu.Unlock()

	l.provider = provider

}

func (l *Logger) SetLevel(level level.Level) {

	l.mu.Lock()

	defer l.mu.Unlock()

	l.level = level

}

func (l *Logger) SetLogStack(logStack bool) {

	l.mu.Lock()

	defer l.mu.Unlock()

	l.logStack = logStack

}
