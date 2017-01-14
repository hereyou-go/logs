package common

import (
	"runtime"
	"time"

	"github.com/one-go/logs/level"
)

type Entry struct {
	Time       time.Time
	Level      level.Level
	Message    []interface{}
	Caller     string
	Line       int
	StackTrace []*runtime.Frame
}
