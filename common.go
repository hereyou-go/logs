package logs

import (
	"os"
	"strings"

	"github.com/one-go/logs/common"
	"github.com/one-go/logs/console"
	"github.com/one-go/logs/level"
)

const (
	defCalldepth = 2
)

var rootLogger *Logger

func init() {
	rootLogger = NewLogger("root", console.NewConsole())
	rootLogger.calldepth = 3
}

/*****************************************/

func GetProvider() common.Provider {
	return rootLogger.Provider()
}

func GetLevel() level.Level {
	return rootLogger.Level()
}

func IsLogStack() bool {
	return rootLogger.IsLogStack()
}

func SetProvider(provider common.Provider) {
	rootLogger.SetProvider(provider)
}

func SetLevel(level level.Level) {
	rootLogger.SetLevel(level)
}

func SetLogStack(logStack bool) {
	rootLogger.SetLogStack(logStack)
}

/*****************************************/

func Debug(message ...interface{}) {
	rootLogger.Log(defCalldepth, level.DEBUG, message...)
}

func Info(message ...interface{}) {
	rootLogger.Log(defCalldepth, level.INFO, message...)
}

func Warn(message ...interface{}) {
	rootLogger.Log(defCalldepth, level.WARN, message...)
}

func Error(message ...interface{}) {
	rootLogger.Log(defCalldepth, level.ERROR, message...)
}

func Fatal(message ...interface{}) {
	rootLogger.Log(defCalldepth, level.FATAL, message...)
}

/*****************************************/

func IsDebugEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

func IsInfoEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

func IsWarnEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

func IsErrorEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

func IsFatalEnabled() bool {
	return rootLogger.IsDebugEnabled()
}

/*****************************************/

func Print(v ...interface{}) {
	os.Stdout.WriteString(common.FormatMessage(common.NoColor, v...))
}

func Println(v ...interface{}) {
	msg := common.FormatMessage(common.NoColor, v...)
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}
	os.Stdout.WriteString(msg)
}

func PanicIf(err error) {
	if err != nil {
		if _, ok := err.(*common.Exception); !ok {
			err = common.NewException(3, "", err)
		}
		panic(err)
	}
}

/*****************************************/

// NewError returns a Exception with the specified detail message and cause.
// cause is a optional
func NewError(code string, message ...interface{}) *common.Exception {
	return common.NewException(3, code, nil, message...)
}

func NewCauseError(code string, cause error, message ...interface{}) *common.Exception {
	return common.NewException(3, code, cause, message...)
}

func WrapError(cause error, code ...string) *common.Exception {
	if cause == nil {
		return nil
	} else if ex, ok := cause.(*common.Exception); ok {
		return ex
	}
	c := ""
	if len(code) > 0 {
		c = code[0]
	}
	return common.NewException(3, c, cause)
}
