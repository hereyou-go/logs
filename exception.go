package logs

import "runtime"

type Exception struct {
	code    string
	message []interface{}
	stack   []byte
	cause   error
	frames  []*runtime.Frame
}

func (ex *Exception) Code() string {
	return ex.code
}

func (ex *Exception) Message() []interface{} {
	return ex.message
}

func (ex *Exception) Stack() []byte {
	return ex.stack
}

func (ex *Exception) Trace() []*runtime.Frame {
	return ex.frames
}

func (ex *Exception) Cause() error {
	return ex.cause
}

func (ex *Exception) Error() string {
	if len(ex.message) == 0 {
		return ex.cause.Error()
	}
	// panic(formatMessage(false, ex.message))
	return formatMessage(false, ex.message...)
}

func getFrames(calldepth int) []*runtime.Frame {
	var frames []*runtime.Frame
	for {
		if pc, file, line, ok := runtime.Caller(calldepth); ok {
			frames = append(frames, &runtime.Frame{
				File:     file,
				Line:     line,
				Function: runtime.FuncForPC(pc).Name(),
			})
		} else {
			break
		}
		calldepth++
	}
	return frames
}

func NewException(calldepth int, code string, cause error, message ...interface{}) *Exception {
	if cause == nil && len(message) == 0 {
		panic("[mano.logs.NewError] Invalid argument")
	}
	frames := getFrames(calldepth)
	return &Exception{
		message: message,
		cause:   cause,
		frames:  frames,
		code:    code,
		//stack:   debug.Stack(),
	}
}

// NewError returns a Exception with the specified detail message and cause.
// cause is a optional
func NewError(cause error, message ...interface{}) *Exception {
	return NewException(2, "", cause, message...)
}

func ToException(err interface{}, calldepth ...int) *Exception {
	if err == nil {
		return nil
	}
	depth := 2
	if len(calldepth) > 0 {
		depth = calldepth[0]
	}
	if ex, ok := err.(*Exception); ok {
		return ex
	} else if ex, ok := err.(error); ok {
		return NewException(depth, "", ex)
	} else if ex, ok := err.(string); ok {
		return NewException(depth, "", nil, ex)
	}
	panic("[mano.logs.ToException] Unsupport argument")
}
