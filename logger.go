package logs

import (
	"runtime"
	"sync"
	"time"

	"github.com/hereyou-go/logs/common"
	"github.com/hereyou-go/logs/level"
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

// Log 将根据给定参数记录日志。
// calldepth 是将要(或)记录的堆栈深度。
// level 日志的级别。
// message 日志信息,这是一个复合的参数(主要是因为GO没有重载)：
// 1.如果第1个参数是 error 类型则只使用这一个，其它忽略或报异常。
// 2.如果第1个参数是 string 类型则可格式化输出，其它为其参数。
// 3.如果最后1个参数是 error 类型则将其附加到已经格式化的消息后面，该处理主要与第2种方式结合使用。
func (l *Logger) Log(calldepth int, level level.Level, message ...interface{}) {
	if !l.isEnabled(level) {
		return
	}

	entry := &common.Entry{
		Level:   level,
		Time:    time.Now(),
		Message: message,
	}

	if pc, file, line, ok := runtime.Caller(calldepth); ok {
		entry.Frame = &common.TraceFrame{
			Frame: runtime.Frame{
				PC:   pc,
				File: file,
				Line: line,
			},
		}
	} else {
		entry.Frame = &common.TraceFrame{}
	}
	// entry.Caller = runtime.FuncForPC(pc).Name()
	//
	// if l.logStack || level >= level.ERROR {
	// 	entry.StackTrace = getFrames(l.calldepth)
	// }

	l.provider.Log(entry)
}

// func (l *Logger) Debug(message ...interface{}) {
// 	l.log(level.DEBUG, message...)
// }

// func (l *Logger) Info(message ...interface{}) {
// 	l.log(level.INFO, message...)
// }

// func (l *Logger) Warn(message ...interface{}) {
// 	l.log(level.WARN, message...)
// }

// func (l *Logger) Error(message ...interface{}) {
// 	l.log(level.ERROR, message...)
// }

// func (l *Logger) Fatal(message ...interface{}) {
// 	l.log(level.FATAL, message...)
// }

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
