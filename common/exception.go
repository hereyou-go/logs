package common

import (
	"fmt"
	"runtime"
	"strings"
)

type Exception struct {
	code    string
	message []interface{}
	cause   error
	frames  []*TraceFrame
}

func (ex *Exception) Code() string {
	return ex.code
}

func (ex *Exception) Arguments() []interface{} {
	return ex.message
}

func (ex *Exception) Message() string {
	if len(ex.message) != 0 {
		return FormatMessage(NoColor, ex.message...)
	}
	if err, ok := ex.Cause().(*Exception); ok {
		return err.Message()
	}
	return ex.Cause().Error()
}
func (ex *Exception) Trace() []*TraceFrame {
	return ex.frames
}

func (ex *Exception) Cause() error {
	return ex.cause
}

func (ex *Exception) Error() string {
	return ExceptionString(ex, NoColor)
}

func getFrames(calldepth int) []*TraceFrame {
	var frames []*TraceFrame
	for {
		if pc, file, line, ok := runtime.Caller(calldepth); ok {
			frame := &TraceFrame{
				Frame: runtime.Frame{
					PC:   pc,
					File: file,
					Line: line,
				},
			}
			frames = append(frames, frame)
			calldepth++
			continue
		}
		break
	}
	return frames
}

func NewException(calldepth int, code string, cause error, message ...interface{}) *Exception {
	if cause == nil && len(message) == 0 {
		panic("[mano.logs.NewException] Invalid arguments")
	}
	return &Exception{
		message: message,
		cause:   cause,
		code:    code,
		frames:  getFrames(calldepth),
	}
}
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
func strTraces(traces []*TraceFrame) string {
	msg := ""
	if traces != nil {
		i := 0
		max := len(traces)
		for i < max {
			frame := traces[i]
			caller := frame.Caller()
			if strings.EqualFold(caller, "runtime.main") || strings.EqualFold(caller, "runtime.goexit") { //skip runtime.?
				break
			}
			msg += fmt.Sprintf("\tat %s(%s:%d)\n", caller, getFilename(frame.File), frame.Line)
			i++
		}
	}
	return msg
}

func exceptionString(ex *Exception, c Color, root bool) string {
	traces := ex.Trace()
	msg := ""
	if len(ex.message) != 0 {
		msg = FormatMessage(c, ex.message...)
		if ex.Cause() == nil {
			ex = nil
		} else {
			if err, ok := ex.Cause().(*Exception); ok {
				ex = err
			} else {
				msg += "[Caused]" + ex.Cause().Error()
				ex = nil
			}
		}
	} else if ex.Cause() != nil {
		if err, ok := ex.Cause().(*Exception); !ok {
			msg = ex.Cause().Error()
			ex = nil
		} else {
			ex = err
			if len(ex.message) != 0 {
				msg = FormatMessage(c, ex.message...)
			}
		}
	}
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}
	msg += strTraces(traces)
	if ex != nil {
		if !root {
			msg += "Caused by\n"
		}
		msg += exceptionString(ex, c, false)
	}
	return msg
}

func ExceptionString(ex *Exception, c Color) string {
	return exceptionString(ex, c, true)
}

// func ToException(err interface{}, calldepth ...int) *Exception {
// 	if err == nil {
// 		return nil
// 	}
// 	depth := 2
// 	if len(calldepth) > 0 {
// 		depth = calldepth[0]
// 	}
// 	if ex, ok := err.(*Exception); ok {
// 		return ex
// 	} else if ex, ok := err.(error); ok {
// 		return NewException(depth, "", ex)
// 	} else if ex, ok := err.(string); ok {
// 		return NewException(depth, "", nil, ex)
// 	}
// 	panic("[mano.logs.ToException] Unsupport argument")
// }
