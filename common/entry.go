package common

import (
	"runtime"
	"time"

	"github.com/one-go/logs/level"
)

type TraceFrame struct {
	runtime.Frame
}

func (f *TraceFrame) Caller() string {
	return runtime.FuncForPC(f.PC).Name()
}

type Entry struct {
	Time    time.Time
	Level   level.Level
	Message []interface{}
	// Caller  string
	// Line    int
	Frame *TraceFrame
	// TracePoint uintptr
	// StackTrace []*runtime.Frame
}
