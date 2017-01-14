package console

import (
	"fmt"
	"strings"

	"github.com/one-go/logs/common"
)

func getFilename(file string) string {
	if file == "" {
		return "???"
	}
	last := strings.LastIndex(file, "/")
	if last < 0 {
		last = strings.LastIndex(file, "\\")
	}
	if last < 0 {
		return file
	}
	return file[last+1:]
}

func FormatLog(provider common.Provider, entry *common.Entry) string {
	if entry == nil {
		return ""
	}
	//fmt.Fprintf(provider.Output(),)
	// msg := entry.Time.Format("2006-01-02 15:04:05")
	// msg += " "
	msg := provider.Color("cyan", fmt.Sprintf("%s %s %s:%d\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level.Lable(), entry.Caller, entry.Line))
	// msg += "\n"
	// msg += fmt.Sprintf("%s%s%s:", colorForLevel(entry.Level, isTerminal), entry.Level.Lable(), ResetSetter(isTerminal))
	// msg += provider.Color("cyan", fmt.Sprintf("%s:%d\n", entry.Caller, entry.Line))
	if ex, ok := entry.Message[0].(*common.Exception); ok {
		msg += ex.Error()
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		i := 0
		max := len(ex.Trace())
		for i < max {
			frame := ex.Trace()[i]
			if strings.EqualFold(frame.Function, "runtime.goexit") {
				break
			}

			msg += fmt.Sprintf("\tat %s(%s:%d)\n", frame.Function, getFilename(frame.File), frame.Line)

			i++
		}

		//msg += string(ex.Stack())
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

	} else {
		msg += common.FormatMessage(provider.Color, entry.Message...)
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

		if entry.StackTrace != nil && len(entry.StackTrace) > 0 {
			i := 0
			max := len(entry.StackTrace)
			for i < max {
				frame := entry.StackTrace[i]
				if strings.EqualFold(frame.Function, "runtime.goexit") {
					break
				}

				msg += fmt.Sprintf("\tat %s(%s:%d)\n", frame.Function, getFilename(frame.File), frame.Line)

				i++
			}
			if !strings.HasSuffix(msg, "\n") {
				msg += "\n"
			}
		}
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
