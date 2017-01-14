package logs

import (
	"os"
	"strings"

	"github.com/one-go/logs/common"
	"github.com/one-go/logs/console"
	"github.com/one-go/logs/level"
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
	rootLogger.Debug(message...)
}

func Info(message ...interface{}) {
	rootLogger.Info(message...)
}

func Warn(message ...interface{}) {
	rootLogger.Warn(message...)
}

func Error(message ...interface{}) {
	rootLogger.Error(message...)
}

func Fatal(message ...interface{}) {
	rootLogger.Fatal(message...)
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
		panic(NewError(err))
	}
}

/*****************************************/

// NewError returns a Exception with the specified detail message and cause.
// cause is a optional
func NewError(cause error, message ...interface{}) *common.Exception {
	return common.NewException(2, "", cause, message...)
}

func NewException(calldepth int, code string, cause error, message ...interface{}) *common.Exception {
	return common.NewException(calldepth, code, cause, message...)
}

func Exception(err interface{}, calldepth ...int) *common.Exception {
	return common.ToException(err, calldepth...)
}
