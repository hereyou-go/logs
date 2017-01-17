package console

import (
	"fmt"
	"strings"

	"github.com/one-go/logs/common"
)

func FormatLog(provider common.Provider, entry *common.Entry) string {
	if entry == nil {
		return ""
	}
	msg := provider.Color("cyan", fmt.Sprintf("%s %s %s:%d\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level.Lable(), entry.Frame.Caller(), entry.Frame.Line))
	var ex *common.Exception
	args := entry.Message
	if e, ok := args[0].(*common.Exception); ok {
		ex = e
		args = args[1:]
	}

	if len(args) > 0 {
		msg += common.FormatMessage(provider.Color, args...)
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
	}

	if ex != nil {
		msg += common.ExceptionString(ex, provider.Color)
	}

	return msg
}

// var (
// 	green   = "" // "\033[34;1m"
// 	white   = "" // string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
// 	yellow  = "" // "\033[33;1m"
// 	red     = "" // "\033[31;1m"
// 	blue    = "\033[32;1m"
// 	magenta = "" // string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
// 	cyan    = "" // string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
// 	reset   = "\033[0m"
// )

// func colorForLevel(level level.Level, isTerminal bool) string {
// 	var setter ColorSetter
// 	switch level {
// 	case LINFO:
// 		setter = GetColorSetter("white")
// 	case LDEBUG:
// 		setter = GetColorSetter("blue")
// 	case LWARN:
// 		setter = GetColorSetter("yello")
// 	case LERROR:
// 		setter = GetColorSetter("red")
// 	}

// 	//TODO
// 	setter = GetColorSetter("green")

// 	if setter != nil {
// 		return setter(isTerminal)
// 	} else if isTerminal {
// 		return ResetSetter(isTerminal)
// 	}

// 	return ""
// }
